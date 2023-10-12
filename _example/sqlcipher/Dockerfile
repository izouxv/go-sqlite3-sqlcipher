# =============================================================================
#  Multi-stage Dockerfile Example
# =============================================================================
#  This is a simple Dockerfile that will build an image of scratch-base image.
#  Usage:
#    docker build -t sqlcipher:local . && docker run --rm sqlcipher:local
# =============================================================================

# -----------------------------------------------------------------------------
#  Build Stage
# -----------------------------------------------------------------------------
FROM golang:alpine  AS build

RUN echo "http://mirrors.aliyun.com/alpine/latest-stable/main/" > /etc/apk/repositories && \
    echo "http://mirrors.aliyun.com/alpine/latest-stable/community/" >> /etc/apk/repositories && \
    apk update

# Important:
#   Because this is a CGO enabled package, you are required to set it as 1.
ENV CGO_ENABLED=1

RUN apk add --no-cache \
    # Important: required for go-sqlite3-sqlcipher
    gcc \
    # Required for Alpine
    musl-dev \
    # encryption
    libressl-dev

WORKDIR /workspace

COPY . /workspace/

RUN \
    go mod init github.com/maxbad/sqlcipher && \
    go mod tidy && \
    go install -tags='sqlcipher,sqlite_omit_load_extension' -ldflags='-s -w -extldflags "-static"' ./sqlcipher.go

RUN \
    # Smoke test
    set -o pipefail; \
    /go/bin/sqlcipher | grep 99\ こんにちは世界099

# -----------------------------------------------------------------------------
#  Main Stage
# -----------------------------------------------------------------------------
FROM scratch

COPY --from=build /go/bin/sqlcipher /usr/local/bin/sqlcipher

ENTRYPOINT [ "/usr/local/bin/sqlcipher" ]
