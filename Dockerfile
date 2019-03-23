FROM golang:1.11-alpine AS builder

ENV GO111MODULE on
COPY . /go/src/github.com/po3rin/gockerfile
RUN CGO_ENABLED=0 go build -o /mocker -tags "$BUILDTAGS" --ldflags '-extldflags "-static"' github.com/po3rin/gockerfile/cmd/gocker

FROM scratch
COPY --from=builder /gocker /bin/gocker
ENTRYPOINT ["/bin/gocker"]
