In general, smaller images are better.
Why? Security, space (⇾ cost).
Therefore, prodution images are usually different to the image that builds the app.

```
docker build -t demo-05 .
docker run -p 8080:8080 demo-05
```

Notice how `docker images -a` has a single docker image listed -- the builder is not present.

Very often when doing multi-stage builds you'll get something wrong, and the error message might not be helpful.

E.g. remove a `WORKDIR` from the above images.

Here's how to build just the builder image:

```
docker build --target builder -t debug .
```
