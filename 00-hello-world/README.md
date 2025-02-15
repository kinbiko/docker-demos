```sh
docker run hello-world
```

- Unable to find 'hello-world:latest' locally.
- Pulls from `library/hello-world`
- Prints `Hello from Docker!` message.

## Memo

- The _image_ is downloaded.
- A container is then run, based on this image.
- `Image:Container` is like `Class:Object` from object-oriented programming.
- The image is `library/hello-world`

## Questions

- We don't need to write a Dockerfile?
- Where is the image downloaded from?
- Where is the Dockerfile for the code that just executed?

## Refs

- [code w/ Dockerfile and C program](https://github.com/docker-library/hello-world)
- [Docker Hub for hello-world](https://hub.docker.com/_/hello-world)
