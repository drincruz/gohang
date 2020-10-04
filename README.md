# gohang
A simple HTTP server for testing

## Build from source

Clone this repo and run `go build .` then you should have a working `gohang` binary.

## Examples

Run `gohang`

### /

```
curl -I localhost:5000/
HTTP/1.1 200 OK
Date: Sun, 04 Oct 2020 01:38:03 GMT
Content-Length: 19
Content-Type: text/plain; charset=utf-8
```

### /404

```
$ curl -I localhost:5000/404
HTTP/1.1 404 Not Found
Date: Sun, 04 Oct 2020 01:37:17 GMT
Content-Length: 26
Content-Type: text/plain; charset=utf-8
```

### /500

```
$ curl -I localhost:5000/500
HTTP/1.1 500 Internal Server Error
Date: Sun, 04 Oct 2020 01:34:28 GMT
Content-Length: 38
Content-Type: text/plain; charset=utf-8
```

### /slow

```
$ time curl localhost:5000/slow
{ "data": "slow response"}
real    0m5.023s
user    0m0.006s
sys     0m0.006s
```

## Build a Docker container

Did you want to run this in a Docker container?
Simply run `docker-compose build` and then `docker-compose up`.

### docker-compose build

```
$ docker-compose build
Building web
Step 1/7 : FROM golang:1.14
 ---> d6747a138341
Step 2/7 : WORKDIR /go/src/app
 ---> Using cache
 ---> e0c45301390b
Step 3/7 : COPY . /go/src/app
 ---> 8baa0d29af0d
Step 4/7 : RUN env GOOS=linux go build gohang.go
 ---> Running in bde0f5e072ab
Removing intermediate container bde0f5e072ab
 ---> 0a454230e0c5
Step 5/7 : RUN go install .
 ---> Running in 9701cb98bf9a
Removing intermediate container 9701cb98bf9a
 ---> c2b7cf38e979
Step 6/7 : ENTRYPOINT /go/src/app/gohang
 ---> Running in 60bbd1397191
Removing intermediate container 60bbd1397191
 ---> ad4f7bc0dbe0
Step 7/7 : EXPOSE 5000
 ---> Running in e56fedc7f4d6
Removing intermediate container e56fedc7f4d6
 ---> a15ced7b2772

Successfully built a15ced7b2772
Successfully tagged gohang_web:latest
```

### docker-compose up

```
$ docker-compose up
Recreating gohang_web_1 ... done
Attaching to gohang_web_1
web_1  | 2020/10/04 21:16:09 [INFO] Now listening on :5000
```
