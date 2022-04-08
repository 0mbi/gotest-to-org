# gotest-to-org
Simple go test report to org-mode transformer

## Building

``` sh
go build ./...
```

## Usage

In your go PWD run
``` sh
$ go test -v ./... > /tmp/gotest.txt
```

For transformation run:

``` sh
$ gottorg /tmp/gotest.txt > gotest.org
```

## Author / Contact
Markus Hihn

## License
MIT
