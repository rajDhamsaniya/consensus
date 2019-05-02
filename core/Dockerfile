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

RUN go get github.com/gogo/protobuf/proto

RUN go get github.com/leesper/couchdb-golang



WORKDIR /core

COPY ./ /core

RUN pwd

RUN ls -la

RUN protoc -I . protoc/helloworld/helloworld.proto --go_out=plugins=grpc:$GOPATH/src
RUN protoc -I . protoc/discovery/discovery.proto --go_out=plugins=grpc:$GOPATH/src
RUN protoc -I . protoc/contractcode/contractcode.proto --go_out=plugins=grpc:$GOPATH/src
RUN protoc -I . protoc/orderer/orderer.proto --go_out=plugins=grpc:$GOPATH/src
RUN protoc -I . protoc/gossip/gossip.proto --go_out=plugins=grpc:$GOPATH/src
RUN protoc protoc/core/core.proto -I . --go_out=plugins=grpc:$GOPATH/src





RUN go build -o contractHandler core/peer/contractHandler.go



# Making image
FROM debian:latest AS host
#RUN apt-get install ca-certificates

WORKDIR /root/

COPY --from=build-env /core/contractHandler .
RUN chmod +x contractHandler
EXPOSE 50051
ENTRYPOINT [ "./contractHandler" ]
