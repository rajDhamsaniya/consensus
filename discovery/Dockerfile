FROM golang:1.11 AS build-env

RUN apt-get update && \
    apt-get -y install git unzip build-essential autoconf libtool

RUN git clone https://github.com/google/protobuf.git && \
    cd protobuf && \
    ./autogen.sh && \
    ./configure && \
    make && \
    make install && \
    ldconfig && \
    make clean && \
    cd .. && \
    rm -r protobuf

RUN go get google.golang.org/grpc

RUN go get github.com/golang/protobuf/protoc-gen-go



WORKDIR /discovery

COPY ./ /discovery

RUN pwd

RUN ls -la

RUN protoc -I protoc/discovery protoc/discovery/discovery.proto --go_out=plugins=grpc:discovery

RUN go build -o registry discovery/discovery_server.go



# Making image
FROM debian:latest AS host
#RUN apt-get install ca-certificates

WORKDIR /root/

COPY --from=build-env /discovery/registry .
RUN chmod +x registry
EXPOSE 50052
ENTRYPOINT [ "./registry" ]