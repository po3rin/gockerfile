package main

import (
	"flag"
	"log"
	"os"

	"github.com/moby/buildkit/client/llb"
	"github.com/moby/buildkit/frontend/gateway/grpcclient"
	"github.com/moby/buildkit/util/appcontext"
	"github.com/po3rin/gockerfile"
	"github.com/po3rin/gockerfile/config"
)

var (
	filename string
	graph    bool
)

func main() {
	flag.BoolVar(&graph, "graph", false, "output a graph and exit")
	flag.StringVar(&filename, "f", "Gockerfile.yaml", "the file to read from")
	flag.Parse()

	if graph {
		c, err := config.NewConfigFromFilename(filename)
		if err != nil {
			log.Fatal(err)
		}

		st, _, err := gockerfile.Gocker2LLB(c)
		if err != nil {
			log.Fatal(err)
		}
		dt, err := st.Marshal()
		if err != nil {
			log.Fatal(err)
		}
		llb.WriteTo(dt, os.Stdout)
		os.Exit(0)
	}

	if err := grpcclient.RunFromEnvironment(appcontext.Context(), gockerfile.Build); err != nil {
		log.Println(err)
	}
}
