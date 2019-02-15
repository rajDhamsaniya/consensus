
sudo apt-get update >> /dev/null

sudo apt-get -y upgrade >> /dev/null

echo "update upgrade done"

wget https://dl.google.com/go/go1.11.4.linux-amd64.tar.gz >> /dev/null

sudo mv go1.11.4.linux-amd64.tar.gz /usr/local >> /dev/null
sudo cd /usr/local >> /dev/null
sudo tar -xvf go1.11.4.linux-amd64.tar.gz >> /dev/null
#sudo mv go /usr/local

echo "tar extracted"

tmp="export GOROOT=/usr/local/go"
echo $tmp >> ~/.profile

tmp="export GOPATH=$HOME/go"
echo $tmp >> ~/.profile

tmp="export PATH=$GOPATH/bin:$GOROOT/bin:$PATH"
echo $tmp >> ~/.profile

source ~/.profile

go version

go get -u google.golang.org/grpc >> /dev/null
echo "go get grpc done"


go get -u github.com/golang/protobuf/protoc-gen-go >> /dev/null
echo "go get protobuf done"

#! /bin/bash
# Make sure you grab the latest version
curl -OL https://github.com/google/protobuf/releases/download/v3.6.1/protoc-3.6.1-linux-x86_64.zip
https://github.com/google/protobuf/releases/download/v3.6.1/protoc-3.6.1-linux-x86_64.zip >> /dev/null
# Unzip
unzip protoc-3.6.1-linux-x86_64.zip -d protoc3 >> /dev/null
echo "unzip protoc done"

# Move protoc to /usr/local/bin/
sudo mv protoc3/bin/* /usr/local/bin/ >> /dev/null

# Move protoc3/include to /usr/local/include/
sudo mv protoc3/include/* /usr/local/include/ >> /dev/null

# Optional: change owner
sudo chown $USER /usr/local/bin/protoc >> /dev/null
sudo chown -R $USER /usr/local/include/google >> /dev/null

sudo apt-get remove docker docker-engine docker.io containerd runc >> /dev/null

sudo apt-get install \
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

sudo apt-get install docker-ce docker-ce-cli containerd.io >> /dev/null

#sudo docker run hello-world >> /dev/null


echo "docker complete vs code start"
curl https://packages.microsoft.com/keys/microsoft.asc | gpg --dearmor > microsoft.gpg >> /dev/null

sudo mv microsoft.gpg /etc/apt/trusted.gpg.d/microsoft.gpg >> /dev/null

sudo sh -c 'echo "deb [arch=amd64] https://packages.microsoft.com/repos/vscode stable main" > /etc/apt/sources.list.d/vscode.list' >> /dev/null

sudo apt update >> /dev/null

sudo apt install code >> /dev/null
echo "vs code installed"