# eagameslib-api
A demo API for eagameslib written in Golang

You can find a front-end demo using this API here: https://github.com/JacyGao/eagameslib

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

    cd $GOPATH/src/github.com/eagameslib-api
    
Build all binaries in developmenet environment and kick off the service by running:

    tools/all.sh
    
<b>Security</b>

<b>Authentication</b> has not been implemented due to time constraints. To ensure security, implement an access token tier such as JWT (https://jwt.io/)

<b>CORS</b> can also be strictly defined to grant access to specific domains. CORS can be defined in /config location

More security? Implement refresh token used together with access token (https://auth0.com/docs/tokens/refresh-token)

<b>Test</b>

<b>Unit test</b> has not been implemented due to time constraints. However it must be implemented to prevent regression in further development.

<b>Integration test</b> has not been implemented due to time constraints. However it's recommended to be implemented to ensure the APIs are serving properly.

The offcial golang testing package: (https://golang.org/pkg/testing/)

<h1>API Documentation</h1>

<h2>GET (/title)</h2>

List all existing titles.

<b>Sample Successful Response:</b>

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

<h2>POST (/title)</h2>

Add a new title.

<b>Parameters:</b>
    
    title   |   string  |   required
    
<b>Sample Json body</b>

    {
	    "title":"Need for speed"
    }
    
<b>Sample Successful Response:</b>

    {
        "Success": true,
        "Body": {
            "id": "avkakpsa9j2t4ihg1fhg",
            "title": "Need for speed",
            "created": "2016-09-26T14:56:07.040516966+10:00"
        },
        "Error": ""
    }
    
<b>Sample Unsuccessful Response:</b>

    {
        "Success": false,
        "Body": {
	    "id": "avkal4sa9j2t4ihg1fi0",
	    "title": "Need for speed",
	    "created": "2016-09-26T14:56:51.384081986+10:00"
        },
        "Error": "Title name already exists"
    }
        
