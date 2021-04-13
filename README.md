# uhttpd

A small Unix domain socket HTTP server for use as a tor hidden service.

## torrc config

```bash
HiddenServiceDir /var/tor/name-of-service/
HiddenServicePort 80 unix:/path/to/httpd.sock
```

## Test the Unix domain socket

```bash
$ socat - UNIX-CONNECT:/path/to/httpd.sock
GET / HTTP/1.0
HOST: HOST
```

## Test connectivity to the hidden service

```bash
$ curl --socks5-hostname 127.0.0.1:9050 http://xxxxxxxxxxxxxxx.onion
```

## Notes

