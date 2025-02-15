See downloaded images:

```
docker images
```

---

```
clear
```

Oops, I forgot to read the output from the container, can I get the logs again?

Yes! With:

```
docker logs
```

Try

```
docker logs hello-world
```

Didn't work!!

Image != container.

See all _containers_ with:

```
docker ps
```

Oh, the container is gone, so I guess I can't get the logs.

Not quite!
The container completed its mission of running `./hello` (c.f. command in demo `00`), so it is currently stopped.

```
docker ps -a
```

This shows all containers, including stopped ones.
Here you find output that contains a `CONTAINER ID` as well as `NAMES`.

```
CONTAINER ID   IMAGE         COMMAND    CREATED          STATUS                     PORTS     NAMES
aab467530219   hello-world   "/hello"   19 minutes ago   Exited (0) 9 minutes ago             xenodochial_rhodes
```

It's the `CONTAINER ID` we need:

```
docker logs aab467530219
```

You don't need to write the whole ID -- just enough to disambiguate it from other containers.

```
docker logs aab
```

---

If a docker container is stopped... Can it be started again?

Yes! You can guess the command.

```
docker start aab
```

## Question

- What do we expect to see if we run `docker logs aab`?
- What happens if we run `docker run hello-world` again?
  - How many containers do we see in `docker ps`?
  - How many images do we see in `docker images`?

---

This can get messy, and images in particular can be many gigabytes in size (hello-world is only 5 kilobytes though!).
How can I clean it up?

```
docker rmi hello-world
```

Ah, it doesn't let me remove it if there are still containers, even stopped, that reference this image.

```
docker rm aab
docker rmi hello-world
```

For a more extreme "just delete everything docker knows about":

```
docker system prune -a
```

Running containers should be unaffected, and must be stopped before they can be pruned.
