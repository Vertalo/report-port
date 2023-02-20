# report-port
Expose the state of one or more tcp ports over http. Return HTTP Status 200 if all ports are open or HTTP Status 500 if one or more ports are not.

```
 env CHECKHOST=the-hostname-to-check.tld PORTS="443,22" go run main.go
```

If all the ports on the-hostname-to-check.tld are open we get `OK`

```
curl -v http://localhost:8080
*   Trying 127.0.0.1:8080...
* Connected to localhost (127.0.0.1) port 8080 (#0)
> GET / HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.85.0
> Accept: */*
>
* Mark bundle as not supporting multiuse
< HTTP/1.1 200 OK
< Date: Mon, 20 Feb 2023 14:25:45 GMT
< Content-Length: 3
< Content-Type: text/plain; charset=utf-8
<
OK
* Connection #0 to host localhost left intact

```

Or `NOT OK` if one or more of the ports are not open.

```
curl -v http://localhost:8080
*   Trying 127.0.0.1:8080...
* Connected to localhost (127.0.0.1) port 8080 (#0)
> GET / HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.85.0
> Accept: */*
>
* Mark bundle as not supporting multiuse
< HTTP/1.1 500 Internal Server Error
< Date: Mon, 20 Feb 2023 14:25:45 GMT
< Content-Length: 3
< Content-Type: text/plain; charset=utf-8
<
NOT OK
* Connection #0 to host localhost left intact
```

# Building
Run `gmake build` to build

# Cleanup
Run `gmake clean` to clean up
