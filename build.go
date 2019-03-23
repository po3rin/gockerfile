package gockerfile

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/containerd/containerd/platforms"
	"github.com/moby/buildkit/client/llb"
	"github.com/moby/buildkit/exporter/containerimage/exptypes"
	"github.com/moby/buildkit/frontend/dockerfile/dockerfile2llb"
	"github.com/moby/buildkit/frontend/gateway/client"
	"github.com/pkg/errors"
	"github.com/po3rin/gockerfile/config"
)

const (
	LocalNameContext      = "context"
	LocalNameDockerfile   = "dockerfile"
	keyTarget             = "target"
	keyFilename           = "filename"
	keyCacheFrom          = "cache-from"
	defaultDockerfileName = "Gockerfile.yaml"
	dockerignoreFilename  = ".dockerignore"
	buildArgPrefix        = "build-arg:"
	labelPrefix           = "label:"
	keyNoCache            = "no-cache"
	keyTargetPlatform     = "platform"
	keyMultiPlatform      = "multi-platform"
	keyImageResolveMode   = "image-resolve-mode"
	keyGlobalAddHosts     = "add-hosts"
	keyForceNetwork       = "force-network-mode"
	keyOverrideCopyImage  = "override-copy-image" // remove after CopyOp implemented
)

func Build(ctx context.Context, c client.Client) (*client.Result, error) {
	cfg, err := GetGockerfileConfig(ctx, c)
	if err != nil {
		return nil, errors.Wrap(err, "got error getting gockerfile")
	}
	st, img, err := Gocker2LLB(cfg)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to convert gocker to llb")
	}

	def, err := st.Marshal()
	if err != nil {
		return nil, errors.Wrapf(err, "failed to marshal local source")
	}
	res, err := c.Solve(ctx, client.SolveRequest{
		Definition: def.ToPB(),
	})
	if err != nil {
		return nil, errors.Wrapf(err, "failed to resolve dockerfile")
	}
	ref, err := res.SingleRef()
	if err != nil {
		return nil, err
	}

	config, err := json.Marshal(img)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to marshal image config")
	}
	k := platforms.Format(platforms.DefaultSpec())

	res.AddMeta(fmt.Sprintf("%s/%s", exptypes.ExporterImageConfigKey, k), config)
	res.SetRef(ref)

	return res, nil
}

func GetGockerfileConfig(ctx context.Context, c client.Client) (*config.Config, error) {
	opts := c.BuildOpts().Opts
	filename := opts[keyFilename]
	if filename == "" {
		filename = defaultDockerfileName
	}

	name := "load Gockerfile"
	if filename != "Gockerfile" {
		name += " from " + filename
	}

	src := llb.Local(LocalNameDockerfile,
		llb.IncludePatterns([]string{filename}),
		llb.SessionID(c.BuildOpts().SessionID),
		llb.SharedKeyHint(defaultDockerfileName),
		dockerfile2llb.WithInternalName(name),
	)

	def, err := src.Marshal()
	if err != nil {
		return nil, errors.Wrapf(err, "failed to marshal local source")
	}

	var dtDockerfile []byte
	res, err := c.Solve(ctx, client.SolveRequest{
		Definition: def.ToPB(),
	})
	if err != nil {
		return nil, errors.Wrapf(err, "failed to resolve dockerfile")
	}

	ref, err := res.SingleRef()
	if err != nil {
		return nil, err
	}

	dtDockerfile, err = ref.ReadFile(ctx, client.ReadRequest{
		Filename: filename,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "failed to read dockerfile")
	}

	cfg, err := config.NewConfigFromBytes(dtDockerfile)
	if err != nil {
		return nil, errors.Wrap(err, "getting config")
	}

	return cfg, nil
}
