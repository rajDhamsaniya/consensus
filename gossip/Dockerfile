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



WORKDIR /gossip

COPY ./ /gossip

RUN pwd

RUN protoc -I . protoc/gossip/gossip.proto --go_out=plugins=grpc:$GOPATH/src
RUN protoc -I . protoc/discovery/discovery.proto --go_out=plugins=grpc:$GOPATH/src
RUN protoc protoc/core/core.proto -I . --go_out=plugins=grpc:$GOPATH/src

RUN go build -o gossip_server gossip/gossip.go



# Making image
FROM debian:latest AS host
#RUN apt-get install ca-certificates

WORKDIR /root/

COPY --from=build-env /gossip/gossip_server .
RUN chmod +x gossip_server
EXPOSE 50052
ENTRYPOINT [ "./gossip_server" ]