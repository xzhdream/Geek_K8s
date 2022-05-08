### 1、Deploy httpserver
```shell
dream@k8s0:~/yaml/istio/homework$ kubectl create ns xzh
namespace/xzh created

dream@k8s0:~/yaml/istio/homework$ kubectl label ns xzh istio-injection=enabled
namespace/xzh labeled

dream@k8s0:~/yaml/istio/homework$ kubectl create -f httpserver.yaml -n xzh
deployment.apps/httpserver created
service/httpserver created
```

### 2、Generate cert
```shell
dream@k8s0:~/yaml/istio/homework$ openssl req -x509 -sha256 -nodes -days 365 -newkey rsa:2048 -subj '/O=cncamp Inc./CN=*.xzh.io' -keyout xzh.io.key -out xzh.io.crt
Generating a RSA private key
...................+++++
....................+++++
writing new private key to 'xzh.io.key'
-----
```

### 3、Create secret
```shell
dream@k8s0:~/yaml/istio/homework$ kubectl create -n istio-system secret tls xzh-credential --key=xzh.io.key --cert=xzh.io.crt
secret/xzh-credential created
```

### 4、Apply gateway
```shell
dream@k8s0:~/yaml/istio/homework$ kubectl apply -f istio-specs.yaml -n xzh
virtualservice.networking.istio.io/httpsserver created
gateway.networking.istio.io/httpsserver created
```

### 5、Check
```shell
dream@k8s0:~/yaml/istio/homework$ curl --resolve httpsserver.xzh.io:443:10.20.62.12 https://httpsserver.xzh.io/healthz -v -k
* Added httpsserver.xzh.io:443:10.20.62.12 to DNS cache
* Hostname httpsserver.xzh.io was found in DNS cache
*   Trying 10.20.62.12:443...
* TCP_NODELAY set
* Connected to httpsserver.xzh.io (10.20.62.12) port 443 (#0)
* ALPN, offering h2
* ALPN, offering http/1.1
* successfully set certificate verify locations:
*   CAfile: /etc/ssl/certs/ca-certificates.crt
    CApath: /etc/ssl/certs
* TLSv1.3 (OUT), TLS handshake, Client hello (1):
* TLSv1.3 (IN), TLS handshake, Server hello (2):
* TLSv1.3 (IN), TLS handshake, Encrypted Extensions (8):
* TLSv1.3 (IN), TLS handshake, Certificate (11):
* TLSv1.3 (IN), TLS handshake, CERT verify (15):
* TLSv1.3 (IN), TLS handshake, Finished (20):
* TLSv1.3 (OUT), TLS change cipher, Change cipher spec (1):
* TLSv1.3 (OUT), TLS handshake, Finished (20):
* SSL connection using TLSv1.3 / TLS_AES_256_GCM_SHA384
* ALPN, server accepted to use h2
* Server certificate:
*  subject: O=cncamp Inc.; CN=*.xzh.io
*  start date: May  8 11:37:45 2022 GMT
*  expire date: May  8 11:37:45 2023 GMT
*  issuer: O=cncamp Inc.; CN=*.xzh.io
*  SSL certificate verify result: self signed certificate (18), continuing anyway.
* Using HTTP2, server supports multi-use
* Connection state changed (HTTP/2 confirmed)
* Copying HTTP/2 data in stream buffer to connection buffer after upgrade: len=0
* Using Stream ID: 1 (easy handle 0x5589592a52f0)
> GET /healthz HTTP/2
> Host: httpsserver.xzh.io
> user-agent: curl/7.68.0
> accept: */*
>
* TLSv1.3 (IN), TLS handshake, Newsession Ticket (4):
* TLSv1.3 (IN), TLS handshake, Newsession Ticket (4):
* old SSL session ID is stale, removing
* Connection state changed (MAX_CONCURRENT_STREAMS == 2147483647)!
  < HTTP/2 503
  < content-length: 201
  < content-type: text/plain
  < date: Sun, 08 May 2022 11:39:44 GMT
  < server: istio-envoy
  <
* Connection #0 to host httpsserver.xzh.io left intact
  upstream connect error or disconnect/reset before headers. reset reason: connection failure, transport failure reason: TLS error: 268436501:SSL routines:OPENSSL_internal:SSLV3_ALERT_CERTIFICATE_EXPIREDdream@k8s0:~/yaml/istio/homework$ 

```






