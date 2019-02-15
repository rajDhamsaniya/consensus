#!/bin/bash
sudo apt-get update >> /dev/null

sudo apt-get -y upgrade >> /dev/null

echo "update upgrade done"

wget https://dl.google.com/go/go1.11.4.linux-amd64.tar.gz >> /dev/null

sudo mv go1.11.4.linux-amd64.tar.gz /usr/local
cd /usr/local
sudo tar -xvf go1.11.4.linux-amd64.tar.gz >> /dev/null
#sudo mv go /usr/local

echo "tar extracted"

tmp="PATH="$HOME/bin:$HOME/.local/bin:$PATH""
echo $tmp >> ~/.profile

tmp="export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin"
echo $tmp >> ~/.profile

tmp="export PATH=$PATH:/home/user/go/bin"
echo $tmp >> ~/.profile

source ~/.profile

go version

go get -u google.golang.org/grpc
echo "go get grpc done"



# Make sure you grab the latest version
curl -OL https://github.com/google/protobuf/releases/download/v3.7.0/protoc-3.7.0-linux-x86_64.zip


mv protoc-3.7.0-linux-x86_64.zip /usr/local
cd /usr/local
# Unzip
sudo unzip protoc-3.7.0-linux-x86_64.zip -d protoc3 >> /dev/null
echo "unzip protoc done"

go get -u github.com/golang/protobuf/protoc-gen-go
echo "go get protobuf done"


# Optional: change owner
sudo chown -R $USER go

sudo chown $USER /usr/local/bin/protoc 
sudo chown -R $USER /usr/local/include/google 

sudo apt-get remove docker docker-engine docker.io containerd runc >> /dev/null

sudo apt-get -y install \
    apt-transport-https \
    ca-certificates \
    curl \
    gnupg-agent \
    software-properties-common >> /dev/null

curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add - >> /dev/null

sudo add-apt-repository \
   "deb [arch=amd64] https://download.docker.com/linux/ubuntu \
   $(lsb_release -cs) \
   stable" >> /dev/null

sudo apt-get update >> /dev/null

sudo apt-get -y install docker-ce docker-ce-cli containerd.io >> /dev/null

sudo docker run hello-world 

