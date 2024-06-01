FROM alpine:latest AS build

RUN apk update && \
    apk add --no-cache \
    build-base \
    cmake=3.28.0 \
    libcrypto
