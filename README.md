# Martini WebMF

This is brief tutorial on using Martini web microframework for the Go Language.

## Install and Run

With Go install, setup [GOPATH](https://golang.org/doc/code.html#GOPATH), and run these in bash:

```bash
go get github.com/go-martini/martini
go run server.go &
curl -i ${WEBSERVER}:3000/
curl -i ${WEBSERVER}:3000/hello/Simon
```

## Resources

* https://github.com/go-martini/martini