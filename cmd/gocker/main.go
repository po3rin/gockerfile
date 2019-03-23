package main

import (
	"fmt"
	"os"

	"github.com/moby/buildkit/client/llb"
	"github.com/po3rin/gockerfile"
)

func main() {
	dt, err := gockerfile.Gocker2LLB()
	if err != nil {
		fmt.Println(err)
		return
	}
	llb.WriteTo(dt, os.Stdout)
}
