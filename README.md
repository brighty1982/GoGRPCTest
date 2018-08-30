# GoGRPCTest
learning GRPC and Protobuf with Go. Further information from which this README is derived can be found [GRPC](https://grpc.io/docs/quickstart/go.html)

## Prerequsites

- Your GO environment is setup and ```GOPATH``` is added to your ```PATH```. You can download the binary from [Go install](https://golang.org/dl/) or use Homebrew as below.

```
$ brew install go
```

The next 2 lines you'll want to add to your ```~/.bash_profile``` else they'll be forgotton when you restart terminal!

```
$ export GOPATH=$HOME/workspace/go

$ export PATH=$PATH:$GOPATH/bin
```   

Create a ```'bin', 'pkg' and 'src'```  directory under you ```GOPATH```. Executables will be compiled to ```'bin'```. Packages will be installed in ```'pkg'``` and all your source code and dependencies will be in ```'src'```

- You have installed protobuf
```
$ brew install protobuf
```

Protobuf will be used to define the message format that we use as the GRPC request payload and the GRPC service contract

- Install GRPC
```
$ go get -u google.golang.org/grpc
```

- Install the Go protobuf generator
```
$ go get -u github.com/golang/protobuf/protoc-gen-go
```

## Usage

The basic idea is for a rpc service ```MeterReadService``` with a single exported end point ```SubmitMeterRead```, with a specified message format for both the input and return values to be running on a server and receiving meter read requests from a client. 

There is no logic inside the server validating the read so it will always return the same success message.

### Defining the service and message content

```
proto/meterReadProto.proto
```
This file defines the service and message content of the rpc service.

```
proto/meterReadProto.pb.go
```
This file is the generated go code created by running the ```protoc``` command and is imported by the server and client code so they know about the service and message format.

You can delete this file and recreate it by running the following. Navigate to the ```GoGRPCTest`` directory and run:

```
$ protoc -I proto/ proto/meterReadProto.proto --go_out=plugins=grpc:proto
```

This tells the protobuf compiler to parse the .proto file and generate the Go source code using the grpc plugin and put the file in the ```proto``` directory

### Server and Client

```
grpcserver/grpcserver.go
```

This file implements the ```SubmitMeterRead``` function of our ```MeterReadService``` and creates a new grpc server listening on the specified port for ```SubmitMeterRead``` requests

```
grpclient/grpcclient.go
```

This file constructs a new ```MeterRead``` and then makes a grpc connection on the specified port and calls ```SubmitMeterRead``` with the ```MeterRead``` as the payload.

### Running

Start the grpc server

```
$ go run GOPATH/src/github.com/user/GoGRPCTest/grpcserver/grpcserver.go
```
you will not see any out put at this point. It is listening for requests.

Run the grpc client

```
$ go run GOPATH/src/github.com/user/GoGRPCTest/grpclient/grpcclient.go
```

Output:

On the server you should now see (dates will be different)

```
2018/08/30 16:00:36 Serial Number: SG628162H
2018/08/30 16:00:36 Timestamp: 2018-08-30 15:00:36.528198058 +0000 UTC
2018/08/30 16:00:36 Reg1: 12745
2018/08/30 16:00:36 Reg2: 2516
```
This is the server parsing the ```MeterRead``` in the payload.

On the client you should now see

```
2018/08/30 16:00:36 Response: valid read for SG628162H
```

This is the grpc response message returned by the server.


