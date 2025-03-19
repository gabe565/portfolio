#syntax=docker/dockerfile:1

FROM --platform=$BUILDPLATFORM tonistiigi/xx:1.6.1 AS xx

FROM --platform=$BUILDPLATFORM golang:1.24.1-alpine AS go-builder
WORKDIR /app

COPY --from=xx / /

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./
COPY migrations/ migrations/
COPY internal/ internal/

ARG TARGETPLATFORM
RUN --mount=type=cache,target=/root/.cache \
  CGO_ENABLED=0 xx-go build -ldflags='-w -s' -trimpath -tags disable_automigrate,grpcnotrace


FROM --platform=$BUILDPLATFORM node:22-alpine AS node-builder
WORKDIR /app

COPY frontend/package.json frontend/package-lock.json frontend/.npmrc ./
RUN npm ci

COPY frontend/ ./
RUN npm run build


FROM alpine:3.21 AS backend
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


FROM backend AS all-in-one
COPY --from=node-builder /app/dist ./public
