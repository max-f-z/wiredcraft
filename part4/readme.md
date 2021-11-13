## Part 4

### prerequisite

```
docker pull redis
docker run -p 6379:6379 --name redis -d redis redis-server --requirepass "123456"
```

### run
```
cd wiredcraft
go mod download
go run main.go
```

### lint & test
```
cd wiredcraft

go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.43.0
golangci-lint run -c .golangci.yml

go test ./...
```

### Swagger is up & running 
```
http://localhost:8080/swagger/index.html
```
