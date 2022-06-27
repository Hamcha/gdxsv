# Multi-step dockerfile so we don't end up with GB-sized environments
ARG GO_VERSION=1.18

# Build step
FROM golang:${GO_VERSION}-alpine AS build
RUN apk add --no-cache protoc git build-base ca-certificates

WORKDIR /src
COPY ./go.mod ./go.sum ./
RUN go mod download

COPY ./ ./

RUN CGO_ENABLED=1 GOOS=linux make install-tools build

# Run step
FROM alpine:edge AS final

# copy compiled app
COPY --from=build /src/bin/gdxsv /gdxsv

ENTRYPOINT ["/gdxsv"]