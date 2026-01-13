# Microservice reklame

Aggregates upcoming movies and serves them to advertisement surfaces around the movie theater.

## Env vars

Check out .env.example for example values

| ENV         | Description                          |
| ----------- | ------------------------------------ |
| LOG_LEVEL   | Log level (DEBUG, INFO, WARN, ERROR) |
| TZ          | Timezone                             |
| SPORED_HOST | Address of spored microservice       |

## Running

Run the application via

```shell
godotenv go run main.go
```

Regenerate swagger docs via

```shell
make docs
```

Regenerate swagger clients via

```shell
make swagger-clients
```

Run all application tests via

```shell
make test
```
