Before we explore how docker-compose works, let's understand the app:

- It's a backend that's returning HTML.
- It connects to a MySQL database.

To see it running, we start everything defined in the `docker-compose.yml` file with:

```
docker compose up
open http://localhost:8080/
```

---

Looking inside the docker-compose.yml file.

- Volumes are persisted on the host, so you can stop and remove all containers, but the DB data will still be there.
- `restart: always` reboots the application until it can connect to MySQL.

---

These files can get quite large, and are useful for, for example small to medium companies for running their services locally.

A large example is [open-telemetry's demo repo](https://github.com/open-telemetry/opentelemetry-demo).
