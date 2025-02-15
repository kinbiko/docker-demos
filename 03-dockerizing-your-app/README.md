This directory contains the source code of a Hello World web server written in Go (like Docker is).

When running locally it looks like this:

```
PORT=8080 go run .
open http://localhost:8080/hello
```

---

Let's run our application in a Docker container by creating a `Dockerfile`.

Build the Dockerfile (create an image) with `docker build -t demo-03 .`.

Use `docker images` to see the new `demo-03` image.

Then run a container based on this image:

```
docker run demo-03
```

Then try going to the listed URL, and...

It doesn't work...

Because the container is entirely isolated, port 8080 is only exposed in the container, not all the way through to the host system.
So we need to add an `EXPOSE 8080` line to the Dockerfile so the container exposes the port.

Aaaand, it still doesn't work...

Just because we've exposed it on the container doesn't mean our host has anything on this port.
When we try to open the URL in our browser we can only access our hosts open port.
We solve this by port-forwarding the port in the `docker run` command.

```
docker system prune -a
docker build -t demo-03 .
docker run -p 8080:8080 demo-03
open http://localhost:8080/hello
```
