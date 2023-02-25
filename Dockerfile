#syntax=docker/dockerfile:1.4

FROM --platform=$BUILDPLATFORM golang:1.20-alpine as go-builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./
COPY migrations/ migrations/

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
  go build -ldflags='-w -s'
EOT


FROM --platform=$BUILDPLATFORM node:18-alpine AS node-builder
WORKDIR /app

COPY frontend/package.json frontend/package-lock.json frontend/.npmrc ./
ARG FONTAWESOME_NPM_AUTH_TOKEN
RUN npm ci

COPY frontend/ ./
RUN npm run build


FROM alpine
WORKDIR /app

RUN apk add --no-cache tzdata

COPY --from=go-builder /app/portfolio ./
COPY --from=node-builder /app/dist ./public

ARG USERNAME=portfolio
ARG UID=1000
ARG GID=$UID
RUN addgroup -g "$GID" "$USERNAME" \
    && adduser -S -u "$UID" -G "$USERNAME" "$USERNAME"
USER $UID

ENV PUBLIC_DIR "public"
CMD ["./portfolio", "serve", "--http=0.0.0.0:80", "--dir=/data"]
