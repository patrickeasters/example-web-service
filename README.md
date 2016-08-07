# example-web-service

Nothing to see here, just a pretty useless web service to use in demos and such.

## Build

```
CGO_ENABLED=0 GOOS=linux go build
docker build -t patrickeasters/example-web-service .
```

## run

```
docker run -d -p 8080:8080 patrickeasters/example-web-service
```
