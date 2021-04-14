[![build](https://github.com/pgollangi/go-portfinder/actions/workflows/build.yml/badge.svg)](https://github.com/pgollangi/go-portfinder/actions/workflows/build.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/pgollangi/fastget)](https://goreportcard.com/report/github.com/pgollangi/fastget)

# go-portfinder
Go implementation of [node-portfinder](https://www.npmjs.com/package/portfinder). A simple tool to find an open port on the current machine.

Installation
--------------

```bash
$ go get github.com/pgollang/go-portfinder
```
# Usage

```go
package main

import (
	"github.com/pgollang/go-portfinder"
)

func main(){
     // scan localhost with a 2 second timeout per port in 5 concurrent threads
     openPort, err := portfinder.GetPort(PortFinderOptions {
         StartPort: 9090,
         StopPort: 9099
     })
}


```
## Author
Built with ❤ by [Prasanna Kumar](https://pgollangi.com/tabs/about/)