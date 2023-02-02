ARG GO_VERSION=1.19

FROM golang:${GO_VERSION} AS base
ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /src
COPY go.mod ./
COPY go.sum ./

RUN go mod download

FROM base AS build

COPY . ./

RUN go build -ldflags="-w -s" -o /app/main cmd/cinder-csi-plugin/main.go

FROM k8s.gcr.io/build-image/debian-base-amd64:v2.1.3

LABEL name="cinder-csi-plugin" \
      license="Apache Version 2.0" \
      maintainers="Kubernetes Authors" \
      description="Cinder CSI Plugin" \
      architecture=amd64 \
      distribution-scope="public" \
      summary="Cinder CSI Plugin" \
      help="none"

RUN clean-install ca-certificates e2fsprogs mount xfsprogs udev

COPY --from=build /app/main /bin/controller
ENTRYPOINT ["/bin/controller"]
