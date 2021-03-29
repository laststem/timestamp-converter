# timestamp-converter
Simple CLI timestamp converter

### Install
```bash
go get -u github.com/laststem/timestamp-converter/cmd/tc
```

### Usage
```bash
$ tc 1593766111
2020-07-03T17:48:31+09:00

$ TZ=America/New_York tc 1593766111
2020-07-03T04:48:31-04:00

$ tc gen
1617027039

$ tc gen -m
1617027100055

$ tc parse "2020-01-01T00:00:00Z"
1577836800

$ tc parse "2020-01-01 00:00:00" -f "2006-01-02 15:04:05"
1577836800
```