package gockerfile

import (
	"github.com/moby/buildkit/client/llb"
	"github.com/moby/buildkit/util/system"
)

// Gocker2LLB converts gockerfile yml to LLB.
func Gocker2LLB() (*llb.Definition, error) {
	state := buildkit()
	dt, err := state.Marshal(llb.LinuxAmd64)
	if err != nil {
		return nil, err
	}
	return dt, nil
}

func goBuildBase() llb.State {
	goAlpine := llb.Image("docker.io/library/golang:1.12-alpine")
	return goAlpine.
		AddEnv("PATH", "/usr/local/go/bin:"+system.DefaultPathEnv).
		AddEnv("GO111MODULE", "on").
		Run(llb.Shlex("apk add --no-cache g++ linux-headers")).
		Run(llb.Shlex("apk add --no-cache git libseccomp-dev make")).Root()
}

func alpineBase() llb.State {
	return llb.Image("docker.io/library/alpine:latest")
}

func copy(src llb.State, srcPath string, dest llb.State, destPath string) llb.State {
	cpImage := llb.Image("docker.io/library/alpine:latest")
	cp := cpImage.Run(llb.Shlexf("cp -a /src%s /dest%s", srcPath, destPath))
	cp.AddMount("/src", src)
	return cp.AddMount("/dest", dest)
}

func buildkit() llb.State {
	src := goBuildBase().
		Run(llb.Shlex("git clone https://github.com/po3rin/gockerfile.git /go/src/github.com/po3rin/gockerfile")).
		Dir("/go/src/github.com/po3rin/gockerfile").
		Run(llb.Shlex("go build -o /bin/go_sample_server ./example/server"))

	r := alpineBase()
	r = copy(src.Root(), "/bin/server", r, "/bin/")
	return r
}
