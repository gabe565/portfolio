#syntax=docker/dockerfile:1.6

FROM --platform=$BUILDPLATFORM golang:1.22.0-alpine AS go-builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./
COPY migrations/ migrations/
COPY internal/ internal/

# Set Golang build envs based on Docker platform string
ARG TARGETPLATFORM
RUN <<EOT
  set -eux
  case "$TARGETPLATFORM" in
    'linux/amd64') export GOARCH=amd64 ;;
    'linux/arm/v6') export GOARCH=arm GOARM=6 ;;
    'linux/arm/v7') export GOARCH=arm GOARM=7 ;;
    'linux/arm64') export GOARCH=arm64 ;;
    *) echo "Unsupported target: $TARGETPLATFORM" && exit 1 ;;
  esac
  go build -ldflags='-w -s' -trimpath -tags disable_automigrate
EOT


FROM --platform=$BUILDPLATFORM node:20-alpine AS node-builder
WORKDIR /app

COPY frontend/package.json frontend/package-lock.json frontend/.npmrc ./
ARG FONTAWESOME_NPM_AUTH_TOKEN
RUN npm ci

COPY frontend/ ./
RUN npm run build


FROM alpine:3.19 as backend
WORKDIR /app

RUN apk add --no-cache tzdata

ARG USERNAME=portfolio
ARG UID=1000
ARG GID=$UID
RUN addgroup -g "$GID" "$USERNAME" \
    && adduser -S -u "$UID" -G "$USERNAME" "$USERNAME"

RUN mkdir pb_data && chown 1000:1000 pb_data

COPY --from=go-builder /app/portfolio ./

USER $UID
CMD ["./portfolio", "serve", "--http=0.0.0.0:80", "--dir=/data", "--public=public"]


FROM backend as all-in-one
COPY --from=node-builder /app/dist ./public
