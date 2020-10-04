# Port Domain Service

Docker, GRPC, 2 microservices and MongoDB party

## Overview
Here we have two microservices which in pair gives a possibility to upload json file with port list and save it to database.
 
The first service plays client role - it is a simple http server that receives http requests and calls the second service via GRPC

The second service is a GRPC server which interacts with database.

Every component is wrapped in docker and all this company spin up by docker-compose.

## Getting started

### Run by docker-compose
Firstly, you should generate proto files by running ```make proto```

Next, run ```docker-compose up``` command in the root of repo.

### Run locally
You also can run it locally by running several commands in separate terminal tabs

```make proto && make mongo && make server``` 

```make client```

Notice - MongoDB runs in daemon mode - you should stop and kill docker container with it to use docker-compose up command

### Test it
To test code run ```make test``` in the root of repo

### Call API
We have prepared command for you - just run ```make json``` and it will send post request to client with ports.json file to save content to database.

Or just go to the root of repo end send curl request:

```curl localhost:8080/json -X POST -F 'uploadFile=@/Users/ashch/go/src/github.com/silverspase/portService/ports.json'```

You will receive result of update ```{"api":"v1","created":1632,"records":1632}```

Where records - all records in file, created - new instances inserted in database, updated - existed instances with new values and skipped - instances without changes.

Another endpoints:

- Get all ports:

```curl localhost:8080```

- Get one port:

```curl localhost:8080/IDBNT``` , where IDBNT - is port unique ID

- Delete port:

```curl localhost:8080/IDBNT -X DELETE``` , where IDBNT - is port unique ID

Enjoy!



