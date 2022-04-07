# syntax=docker/dockerfile:1.3-labs
ARG GO_VERSION=1.16

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
    go build -ldflags="-w -s" -o /app/main cmd/openstack-cloud-controller-manager/main.go

FROM amd64/alpine:3.11

LABEL name="openstack-cloud-controller-manager" \
      license="Apache Version 2.0" \
      maintainers="Kubernetes Authors" \
      description="OpenStack cloud controller manager" \
      architecture=$ARCH \
      distribution-scope="public" \
      summary="OpenStack cloud controller manager" \
      help="none"

RUN apk add --no-cache ca-certificates

COPY --from=build /app/main /bin/controller
USER 65532:65532
ENTRYPOINT ["/bin/controller"]
