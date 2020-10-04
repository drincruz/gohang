# gohang
A simple HTTP server for testing

## Build from source

Clone this repo and run `go build .` then you should have a working `gohang` binary.

## Examples

Run `gohang`

### /

```
curl -I localhost:80/
HTTP/1.1 200 OK
Date: Sun, 04 Oct 2020 01:38:03 GMT
Content-Length: 19
Content-Type: text/plain; charset=utf-8
```

### /404

```
$ curl -I localhost:80/404
HTTP/1.1 404 Not Found
Date: Sun, 04 Oct 2020 01:37:17 GMT
Content-Length: 26
Content-Type: text/plain; charset=utf-8
```

### /500

```
$ curl -I localhost:80/500
HTTP/1.1 500 Internal Server Error
Date: Sun, 04 Oct 2020 01:34:28 GMT
Content-Length: 38
Content-Type: text/plain; charset=utf-8
```

### /slow

```
$ time curl localhost:80/slow
{ "data": "slow response"}
real    0m5.023s
user    0m0.006s
sys     0m0.006s
```
