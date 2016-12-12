#### Dash Server

Small, command-line HTTP/2 file server for MPEG-DASH content.


Example usage with self-signed certificate:
```
curl https://github.com/theolampert/dash-server/releases/download/0.0.1/dash-server | echo $GOPATH/bin
openssl req -x509 -newkey rsa:4096 -keyout key.pem -out cert.pem -days 365
dash-server path/to/your/cert.pem path/to/your/cert.key .
```

Usage with Letsencrypt:
```
curl https://github.com/theolampert/dash-server/releases/download/0.0.1/dash-server | echo $GOPATH/bin
certbot certonly
dash-server path/to/your/cert.pem path/to/your/cert.key .
```

Develop:
```
git clone git@github.com:theolampert/dash-server.git
cd dash-server && glide install
go run dash-server.go .
```
