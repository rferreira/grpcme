# GrpcMe makes any executable a GRPC Service

### But why? 
Let's say your containerized application depends on some good-old executables to work, like cUrl for example. Today the two options you have are either adding that executable to your main application image (potentially slowing down the build time significantly) or you can build a bespoke service to wrap that executable so it can be package as a standalone container. Well, GrpcMe does the latter for you. 

For example, turning cUrl into a GRPC service: 

```shell
% ./dist/grpcmed /usr/bin/curl  
INFO[2022-08-28T17:48:27-04:00] wrapping executable /usr/bin/curl            
INFO[2022-08-28T17:48:27-04:00] pre-flight check passed, starting GRPC server... 
INFO[2022-08-28T17:48:27-04:00] Listening on 127.0.0.1:8089                 
```

For convenience, GrpcMe also provides a client you can use to test your freshly minted GRPC service: 

