# syntax=docker/dockerfile:1.3-labs
ARG GO_VERSION=1.16

# get modules, if they don't change the cache can be used for faster builds
FROM golang:${GO_VERSION} AS base
ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /src
COPY go.* .

RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download

FROM base AS build

RUN --mount=target=. \
    --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    go build -ldflags="-w -s" -o /app/main cmd/cinder-csi-plugin/main.go

# Import the binary from build stage
FROM k8s.gcr.io/build-image/debian-base-amd64:v2.1.3

LABEL name="cinder-csi-plugin" \
      license="Apache Version 2.0" \
      maintainers="Kubernetes Authors" \
      description="Cinder CSI Plugin" \
      architecture=amd64 \
      distribution-scope="public" \
      summary="Cinder CSI Plugin" \
      help="none"

# Install e4fsprogs for format
RUN clean-install ca-certificates e2fsprogs mount xfsprogs udev

COPY --from=build /app/main /bin/controller

#USER 65532:65532
ENTRYPOINT ["/bin/controller"]
