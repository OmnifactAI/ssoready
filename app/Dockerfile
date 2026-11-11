# The build stage holds the full toolchain (tailwind, esbuild, ...). The
# runtime stage only gets the built SPA and the two packages bin/serve-public
# needs (serve/package.json), so the image ships no build tooling.
#
# Base images are pinned by digest. Dependabot opens a PR when a new digest is
# published, and merging it rebuilds and republishes the image.
FROM node:26-trixie-slim@sha256:5adad64f9a1fca2d51972caa65983634c8a774eb300d72b20f7f2a0f60368072 AS build

WORKDIR /work

COPY package.json package-lock.json ./
RUN npm ci --ignore-scripts --no-audit --no-fund

COPY .env tailwind.config.js build.mjs tsconfig.json ./
COPY bin ./bin
COPY public ./public
COPY src ./src

RUN npm run build-docker

FROM node:26-trixie-slim@sha256:5adad64f9a1fca2d51972caa65983634c8a774eb300d72b20f7f2a0f60368072

# Apply Debian security updates that are newer than the base image, so the
# image does not wait for an upstream rebuild of node:26-trixie-slim.
RUN apt-get update \
    && apt-get upgrade -y --no-install-recommends \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /work

COPY serve/package.json serve/package-lock.json ./
RUN npm ci --omit=dev --ignore-scripts --no-audit --no-fund \
    && npm cache clean --force \
    # The runtime never calls npm or yarn; dropping them removes their bundled
    # dependencies from the image and from its scan findings.
    && rm -rf /usr/local/lib/node_modules /usr/local/bin/npm /usr/local/bin/npx \
        /usr/local/bin/corepack /usr/local/bin/yarn /usr/local/bin/yarnpkg /opt/yarn-*

COPY bin ./bin
# bin/prepare-config-json reads the Config type from src/config.tsx at start.
COPY src/config.tsx ./src/config.tsx
COPY --from=build --chown=node:node /work/public ./public

USER node

# Same as `npm run serve-docker`: write public/config.json from the
# environment, then serve the SPA.
ENTRYPOINT ["/bin/sh", "-c", "./bin/prepare-config-json > ./public/config.json && exec ./bin/serve-public"]
