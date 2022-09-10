## GrpcMe makes any executable a GRPC Service

### But why? 
Let's say your containerized application depends on some good-old executables to work, like cUrl for example. Today the 
two options you have are either adding that executable to your main application image (potentially slowing down the 
build time significantly) or you can build a bespoke service to wrap that executable so it can be package as a 
standalone container - well, GrpcMe does the latter for you, hopefully.  

For example, turning cUrl into a GRPC service: 

```shell
% cat etc/configs/curl.toml 
listen = "127.0.0.1:8089"

[executable.curl]
path = "/usr/bin/curl"
healthcheck_argument = "--version"

 % ./dist/grpcmed etc/configs/curl.toml
INFO[2022-09-10T12:53:43-04:00] parsing etc/configs/curl.toml                
INFO[2022-09-10T12:53:43-04:00] Parsing exposed executables...               
INFO[2022-09-10T12:53:43-04:00] [curl] -> /usr/bin/curl                      
INFO[2022-09-10T12:53:43-04:00] Server version 0                             
INFO[2022-09-10T12:53:43-04:00] Listening on 127.0.0.1:8089  
              
```

For convenience, GrpcMe also provides a client you can use to test your freshly minted GRPC service:

```shell
 % ./dist/grpcme dns:127.0.0.1:8089 curl -I https://www.google.com
HTTP/2 200 
content-type: text/html; charset=ISO-8859-1
p3p: CP="This is not a P3P policy! See g.co/p3phelp for more info."
date: Sat, 10 Sep 2022 16:54:40 GMT
server: gws
x-xss-protection: 0
x-frame-options: SAMEORIGIN
expires: Sat, 10 Sep 2022 16:54:40 GMT
cache-control: private
set-cookie: 1P_JAR=2022-09-10-16; expires=Mon, 10-Oct-2022 16:54:40 GMT; path=/; domain=.google.com; Secure
set-cookie: AEC=AakniGOYVpBG1mXusM3UkOV_rGFkUrOoNkaqevlr8ELzB9ok0Q4P39u-8Q; expires=Thu, 09-Mar-2023 16:54:40 GMT; path=/; domain=.google.com; Secure; HttpOnly; SameSite=lax
set-cookie: NID=511=L7-XJhgIs68pr7-ZFbRFxDDKfgrPlRJbY08hlGaVZBcF1hNBvOqgy7RJhG1hU2fZozOlI0NJP3_3G6Bnrp3FDdxYDgo9o8zMuW9a3KKKleDWBrvOFWOX6vRCLNOHkxwpE2nr5co03RLNETKpzNdD_g8jsOVVmLME31-hYTvzP5A; expires=Sun, 12-Mar-2023 16:54:40 GMT; path=/; domain=.google.com; HttpOnly
alt-svc: h3=":443"; ma=2592000,h3-29=":443"; ma=2592000,h3-Q050=":443"; ma=2592000,h3-Q046=":443"; ma=2592000,h3-Q043=":443"; ma=2592000,quic=":443"; ma=2592000; v="46,43"

```

