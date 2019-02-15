#!/bin/bash
sudo apt-get update >> /dev/null

sudo apt-get -y upgrade >> /dev/null

echo "update upgrade done"

go version
if [ $? -eq 0 ]; then
    echo "PASS"

else
    wget https://dl.google.com/go/go1.11.4.linux-amd64.tar.gz >> /dev/null

    sudo mv go1.11.4.linux-amd64.tar.gz /usr/local
    cd /usr/local
    sudo tar -xvf go1.11.4.linux-amd64.tar.gz >> /dev/null
    #sudo mv go /usr/local

    echo "tar extracted"

    echo "PATH=\"\$HOME/bin:\$HOME/.local/bin:\$PATH\"" >> ~/.profile

    echo "export PATH=\$PATH:/usr/local/go/bin:\$GOPATH/bin" >> ~/.profile

    echo "export PATH=\$PATH:/home/user/go/bin" >> ~/.profile

    source ~/.profile

    go version
fi

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

cd ~
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

