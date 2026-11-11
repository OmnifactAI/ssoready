# Builds the auth, api or migrate image. The build context is the repository
# root, and CMD names the binary under cmd/:
#
#   docker build -f docker/go.Dockerfile --build-arg CMD=auth .
#
# The runtime image holds only that binary, CA certificates and time zone data.
#
# Base images are pinned by digest. Dependabot opens a PR when a new digest is
# published, and merging it rebuilds and republishes the images.
FROM --platform=$BUILDPLATFORM golang:1.27-alpine@sha256:fa3d596898d9f126af7c01d14ee7597bab3f39cccf392ef36b69fe77cd681f18 AS source

RUN apk update && apk add --no-cache git ca-certificates tzdata && update-ca-certificates

WORKDIR /work

COPY go.mod go.sum ./
RUN go mod download

COPY cmd ./cmd
COPY internal ./internal

# The images workflow runs the Go tests in this stage (--target source).
FROM source AS build

ARG TARGETOS
ARG TARGETARCH
ARG CMD
RUN test -n "$CMD" || { echo "Set --build-arg CMD=auth, api or migrate." >&2; exit 1; }
RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -ldflags "-w -s" -o ssoready "./cmd/$CMD"

FROM scratch

COPY --from=build /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
# One fixed path for all three binaries: an exec-form ENTRYPOINT cannot use
# the CMD build arg, and scratch has no shell to expand it.
COPY --from=build /work/ssoready /ssoready

ENTRYPOINT ["/ssoready"]
