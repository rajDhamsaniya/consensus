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

RUN go get github.com/Shopify/sarama





WORKDIR /orderer

COPY ./ /orderer

# RUN go install ./kafka/
# RUN protoc -I protoc/helloworld protoc/helloworld/helloworld.proto --go_out=plugins=grpc:protoc/helloworld
# RUN protoc -I protoc/discovery protoc/discovery/discovery.proto --go_out=plugins=grpc:protoc/discovery
# RUN protoc -I protoc/contractcode protoc/contractcode/contractcode.proto --go_out=plugins:protoc/contractcode
RUN protoc -I . protoc/orderer/orderer.proto --go_out=plugins=grpc:$GOPATH/src

RUN go build -o serviceNode orderer/serviceNode.go



# Making image
FROM debian:latest AS host
#RUN apt-get install ca-certificates

WORKDIR /root/

COPY --from=build-env /orderer .
RUN chmod +x serviceNode
EXPOSE 50054
ENTRYPOINT [ "./serviceNode" ]
