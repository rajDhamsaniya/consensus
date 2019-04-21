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



WORKDIR /example_cc

COPY ./ /example_cc


# RUN protoc -I protoc/helloworld protoc/helloworld/helloworld.proto --go_out=plugins=grpc:helloworld

RUN protoc -I . protoc/contractcode/contractcode.proto --go_out=plugins=grpc:$GOPATH/src
# RUN protoc -I protoc/contractcode protoc/contractcode/contractcode.proto --go_out=plugins=grpc:contractcode

RUN go build -o example example_cc/example_new.go



# Making image
FROM debian:latest AS host
#RUN apt-get install ca-certificates

WORKDIR /root/

COPY --from=build-env /example_cc/example .
RUN chmod +x example
EXPOSE 50053
ENTRYPOINT [ "./example" ]
