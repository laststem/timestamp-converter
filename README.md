# timestamp-converter
Simple CLI timestamp converter

### Install
```bash
go get -u github.com/laststem/timestamp-converter/cmd/tc
```

### Usage
- tc {timestamp}

### Example
```bash
$ tc 1593766111
2020-07-03T17:48:31+09:00

$ TZ=America/New_York tc 1593766111
2020-07-03T04:48:31-04:00
```