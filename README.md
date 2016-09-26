# eagameslib-api
A demo API for eagameslib written in Golang

<b>One-time setup</b>

Install go: - https://golang.org/doc/install - Make sure you have version 1.6 or later for vendor folder support

Set the GOPATH environment variable, and add GOPATH/bin to PATH.

    mkdir -p /scratch/go
    export GOPATH='/scratch/go'
    export PATH=$GOPATH/bin:$PATH
    
Clone the eagameslib-api repo into the place that Go expects to find it.

    mkdir -p $GOPATH/src/github.com
    cd $GOPATH/src/github.com
    git clone https://github.com/JacyGao/eagameslib-api.git
