FROM ubuntu:18.04

RUN sed -i -e 's,archive.ubuntu.com,ja.archive.ubuntu.com,g' /etc/apt/sources.list
RUN apt-get update && \
    apt-get install -y make git gcc-aarch64-linux-gnu golang && \
    apt-get clean all

RUN git clone https://github.com/friendlyarm/WiringNP.git
WORKDIR WiringNP/wiringPi
RUN sed -i 's/gcc/aarch64-linux-gnu-gcc/g' Makefile && \
    make && make install && \
    mv /usr/local/lib/libwiringPi.so.2.0 /usr/local/lib/libwiringPi.so

WORKDIR /src
ENV CC=aarch64-linux-gnu-gcc
ENV GOOS=linux
ENV GOARCH=arm64
ENV CGO_ENABLED=1
CMD go build -v -o bin/ledmatrix .
