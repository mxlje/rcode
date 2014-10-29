# Response Code Chain Checker

The `rcode` CLI lets you see the raw response code and, if applicable, the
redirect chain of a URL. It ends if the response code is `200 OK`.

## Installation

Assuming you have a working Go environment and `GOPATH/bin` is in your
`PATH`, install `rcode` with

```shell
$ go get github.com/mxlje/rcode
```

Alternatively you can grab the latest compiled version from the
[releases section on GitHub][releases] and put it in your `PATH` by hand.

  [releases]: https://github.com/mxlje/rcode/releases

## Usage

```shell
$ rcode http://miele.de
[301] http://miele.de
[200] http://www.miele.de/haushalt/index.htm
```

[httpbin.org](httpbin.org) is a great service to test redirect chains:

```shell
$ rcode httpbin.org/redirect/2
[302] http://httpbin.org/redirect/2
[302] http://httpbin.org/redirect/1
[200] http://httpbin.org/get
```

### CSV Output

Useful for further analysis.

```shell
$ rcode httpbin.org/redirect/1 --csv
302,http://httpbin.org/redirect/1
200,http://httpbin.org/get
```

## License

This project is released under the [WTFPL](http://www.wtfpl.net/). Enjoy.