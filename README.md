# RBAC api with Go

## Development

### 1. Run docker compose and Attach the container
<div align="center">
    <img src="ref/open-devcontainer.png" style="width:60%;" />
</div>

### 2. Run main.go on debugger

<div align="center">
    <img src="ref/run-debug.png" style="width:60%;"></img>
</div>

### 3. Create docs

```
$ cd /root/src/
$ swag init
```

### 4. Open Browser
<div align="center">
    <img src="ref/swagger.png" style="width:60%;"></img>
</div>

## Build deployment image

```
$ docker build -t rbac-go:latest .

$ docker run -p 8000:8000 --env-file .env rbac-go:latest
```

## Reference
  - Hands-On Full-Stack Development with Go
  - https://github.com/gin-gonic/gin
  - https://github.com/golang-jwt/jwt