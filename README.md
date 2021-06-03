# uhttpd

A small Unix domain socket HTTP server for use as a tor hidden service.

## torrc config

```bash
HiddenServiceDir /var/lib/tor/name-of-service/
HiddenServicePort 80 unix:/var/lib/tor-onion-sockets/httpd.sock
```

## Test the Unix domain socket

```bash
$ socat - UNIX-CONNECT:/var/lib/tor-onion-sockets//httpd.sock
GET / HTTP/1.0
HOST: HOST
```

## Test connectivity to the hidden service

```bash
$ curl --socks5-hostname 127.0.0.1:9050 http://xxxxxxxxxxxxxxx.onion
```

## Notes

### Debian Unix Domain Socket Issue

  * From: https://gitweb.torproject.org/debian/tor.git/tree/debian/README.Debian
  * Solution below works on Debian 10

#### Onion services using UNIX domain sockets

Tor's onion services can access their backends not only via TCP but also
via UNIX domain sockets (see also the tor(1) manpage).

UNIX domain sockets are not currently covered by filesystem restrictions
from systemd and apparmor. As such, they can be put anywhere as long as
the filesystem namespace is visible. In particular, this means /home
does not work by default, but anything under /var should.

It is suggested that sockets be placed in a special directory under
/var/lib [3]. Note that while systemd and apparmor do not currently
limit access to sockets, you still need to ensure that the Tor
process may access them according to the default UNIX file and directory
permissions.

Example: The backend is run by user webfu. The default Tor instance
should be able to access the socket.

The admin creates the directory /var/lib/tor-onion-sockets/default/webfu/
and makes it mode 02750 owned by webfu:debian-tor. Then she configures
the backend service to create the socket in this directory and
configures Tor to use that socket. The socket should be read and
writable by the tor process as well, so it should be either g+rw when
its group can be debian-tor, or a+rw when not.

  1. /lib/systemd/system/tor@default.service and /lib/systemd/system/tor@.service
  2. /var/lib/tor and /var/lib/tor-instances/<instancename>
  3. /var/lib/tor-onion-sockets/default and /var/lib/tor-onion-sockets/<instancename>
  4. https://bugs.debian.org/846275

## Misc Unrelated Tor Hidden Service Notes

To enable an SSH server as a Tor Hidden Service, add these two lines to the server's /etc/tor/torrc:

```
HiddenServiceDir /var/lib/tor/ssh-server/
HiddenServicePort 22
```

Then, get the onion name of the service from /var/lib/tor/ssh-server/hostname:

```
$ cat /var/lib/tor/ssh/hostname
xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx.onion
```

Finally, ssh to the server from a client:

```
$ torsocks ssh user@xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx.onion
```

If you only use keys to authenticate, you can add normal host definitions to /home/user/ssh/config:

```
host xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx.onion
user bob
identityfile /home/bob/.ssh/x16
```
