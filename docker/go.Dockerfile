# Builds the auth, api or migrate image. The build context is the repository
# root, and the target names the binary under cmd/:
#
#   docker build -f docker/go.Dockerfile --target auth .
#
# Each binary keeps its name (/auth, /api, /migrate): conf and secretload derive
# the environment variable prefix from os.Args[0] (AUTH_DB, API_DB, ...).
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
RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -ldflags "-w -s" -o /out/ ./cmd/auth ./cmd/api ./cmd/migrate

FROM scratch AS runtime

COPY --from=build /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

FROM runtime AS auth
COPY --from=build /out/auth /auth
ENTRYPOINT ["/auth"]

FROM runtime AS api
COPY --from=build /out/api /api
ENTRYPOINT ["/api"]

FROM runtime AS migrate
COPY --from=build /out/migrate /migrate
ENTRYPOINT ["/migrate"]
