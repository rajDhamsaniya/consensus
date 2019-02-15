FROM golang:1.11

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

RUN ls -la

WORKDIR /helloworld

COPY . /helloworld

RUN protoc -I helloworld/ helloworld/helloworld.proto --go_out=plugins=grpc:helloworld

CMD ["go", "run", "helloworld/greeter_server/main.go"]