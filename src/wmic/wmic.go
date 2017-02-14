package wmic

import (
	"bytes"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"os/exec"
	"regexp"
	"runtime"
	"strconv"
	"strings"
)

type Output map[string]interface{}

func Exec(args ...string) ([]Output, error) {
	cmd := exec.Command("wmic", args...)
	cmdOut := new(bytes.Buffer)
	cmd.Stdin = new(bytes.Buffer)
	cmd.Stdout = cmdOut

	err := cmd.Run()
	if err != nil {
		return nil, err
	}
	return Parse(cmdOut.String())
}

func Translate(reader io.Reader) ([]Output, error) {
	var bs []byte
	var err error
	if runtime.GOOS == "windows" {
		bs, err = ioutil.ReadAll(reader)
	} else {
		bs, err = ReadUTF16LE(reader)
	}
	if err != nil {
		return nil, err
	}
	return Parse(string(bs))
}

func ReadUTF16LE(reader io.Reader) ([]byte, error) {
	utf16 := unicode.UTF16(unicode.LittleEndian, unicode.UseBOM)
	utf16reader := transform.NewReader(reader, utf16.NewDecoder())
	bytes, err := ioutil.ReadAll(utf16reader)
	return bytes, err
}

var newlineRegex = regexp.MustCompile(`\r\n`)
var wmicHeaderLabelRegex = regexp.MustCompile(`\w+\s*`)

func Parse(s string) ([]Output, error) {
	items := make([]Output, 0)
	lines := newlineRegex.Split(s, -1)
	if len(lines) <= 1 {
		return items, nil
	}

	// Parse the header labels and bounds.
	headerLine := []byte(lines[0])
	matchIndices := wmicHeaderLabelRegex.FindAllSubmatchIndex(headerLine, -1)
	headerLabels := make([]string, len(matchIndices))
	for i, indices := range matchIndices {
		a, b := indices[0], indices[1]
		headerLabels[i] = strings.TrimSpace(string(headerLine[a:b]))
	}

	// Parse values from the data lines.
	for _, line := range lines[1:] {
		if line != "" {
			lineLen := len(line)
			values := map[string]interface{}{}
			for i, matchIndex := range matchIndices {
				a, b := matchIndex[0], matchIndex[1]
				if (lineLen >= a) && (lineLen >= b) {
					value := valueFor(strings.TrimSpace(line[a:b]))
					if value != nil {
						values[headerLabels[i]] = value
					}
				}
			}
			if len(values) > 0 {
				items = append(items, values)
			}
		}
	}

	return items, nil
}

var digitsRegex = regexp.MustCompile(`^-?\d+$`)
var trueRegex = regexp.MustCompile(`^[Tt][Rr][Uu][Ee]$`)
var falseRegex = regexp.MustCompile(`^[Ff][Aa][Ll][Ss][Ee]$`)

func valueFor(s string) interface{} {
	switch true {
	case s == "":
		return nil
	case trueRegex.MatchString(s):
		return true
	case falseRegex.MatchString(s):
		return false
	case digitsRegex.MatchString(s):
		n, err := strconv.ParseInt(s, 10, 64)
		if err == nil {
			return n
		}
		return s
	default:
		return s
	}
}
