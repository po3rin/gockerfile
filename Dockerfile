FROM golang:1.12-alpine AS builder

ENV GO111MODULE on
WORKDIR /go/src/github.com/po3rin/gockerfile
COPY . ./
RUN apk update && apk add git && \
    CGO_ENABLED=0 go build -o /gocker -tags "v0.0.1" --ldflags "-s -w" ./cmd/gocker

FROM scratch
COPY --from=builder /gocker /bin/gocker
ENTRYPOINT ["/bin/gocker"]
