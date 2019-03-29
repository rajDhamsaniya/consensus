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



WORKDIR /helloworld

COPY ./ /helloworld

RUN pwd

RUN ls -la

RUN protoc -I protoc/helloworld protoc/helloworld/helloworld.proto --go_out=plugins=grpc:helloworld
RUN protoc -I protoc/discovery protoc/discovery/discovery.proto --go_out=plugins=grpc:discovery

RUN go build -o server helloworld/greeter_server/main.go



# Making image
FROM debian:latest AS host
#RUN apt-get install ca-certificates

WORKDIR /root/

COPY --from=build-env /helloworld/server .
RUN chmod +x server
EXPOSE 50051
ENTRYPOINT [ "./server" ]