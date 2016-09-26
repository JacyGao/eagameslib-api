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

<b>Kick off the service</b>

Make sure you are in the service base location:

    cd $GOPATH/src/github.com/eagames
    
Build all binaries in developmenet environment and kick off the service by running:

    tools/all.sh
    
    
<h2>API Documentation</h2>

<b>GET (/title)</b>

List all existing titles.

<b>Sample successful response:</b>

    {
        "Success": true,
        "Body": [
            {
                "id": "avjmdick5bn6bia5v4o0",
                "title": "Real Racing",
                "created": "2016-09-25T15:55:21.08610998+10:00"
            },
            {
                "id": "avjmehck5bn6ccvlrdp0",
                "title": "Real Racing 2",
                "created": "2016-09-25T15:57:25.637563598+10:00"
            }
        ],
        "Error": ""
    }
