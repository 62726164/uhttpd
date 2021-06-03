## SSH as a Tor Location Hidden Service

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
