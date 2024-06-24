FROM debian:bookworm-slim AS build

RUN --mount=type=cache,target=/var/cache/apt \
    apt-get -y update && apt-get -y install \
    git cmake pkg-config \
    libcrypto++-dev \
    libcrypto++ \
    build-essential curl zip unzip autoconf autoconf-archive nasm \
    libasio-dev zlib1g-dev


WORKDIR /stealthflix

COPY . .

RUN git submodule update --init

RUN mkdir build/

WORKDIR /stealthflix/build
RUN cmake ..
RUN make .

# run the app
RUN ./stealthflix
