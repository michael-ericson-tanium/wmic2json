# wmic2json

Convert WMIC output to JSON.

```
C:\>wmic logon
AuthenticationPackage  Caption  Description  InstallDate  LogonId  LogonType  Na
me  StartTime                  Status
NTLM                                                      294547   2
    20170213214032.539705-420  Error
NTLM                                                      294511   2
    20170213214032.538707-420

C:\>wmic2json logon
[
  {
    "AuthenticationPackage": "NTLM",
    "LogonId": 293547,
    "LogonType": 2,
    "StartTime": "20170213214032.539705-420",
    "Status": "Error"
  },
  {
    "AuthenticationPackage": "NTLM",
    "LogonId": 293521,
    "LogonType": 2,
    "StartTime": "20170213214032.538707-420"
  }
]
```


## Hacking

1. `export GOPATH=$(pwd)`
1. `go get golang.org/x/text`
1. `GOOS=windows GOARCH=amd64 go install wmic2json`
