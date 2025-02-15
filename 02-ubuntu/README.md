Let's try something more useful:

```
docker run -it ubuntu bash
```

- `-it`: Think of this as `InTeractive`, although this is a lie.
- The _image_ is `ubuntu`.
- The _process to execute inside_ is `bash`.

Prove that it's actually Ubuntu:

```
cat /etc/os-release
```

It's a very bare-bones installation of Ubuntu, but you're root, so you can do anything you want.
Including installing software like `vim`.

---

For now we've seen a very simple application as well as a whole operating system.
These are actually both very basic.
You can get Docker to do some really fancy stuff if you set your mind to it:

- Jupyter notebooks
- DOOM (obviously).
- macOS (a little naughty!).
- Old Windows versions.
- Spotify.

Good luck!
