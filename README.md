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

You can also `wmic logon | wmic2json`.

## Recent Builds

1. Version 0.0
    1. Windows 64: [wmic2json-0.0.zip](https://tanium.box.com/s/irk164nxqizl88i3fmq3rzxenoyqrruq)

## Hacking

1. `export GOPATH=$(pwd)`
1. `go get golang.org/x/text`
1. `GOOS=windows GOARCH=amd64 go install wmic2json`
