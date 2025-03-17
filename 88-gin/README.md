### Create a docker network 

```
docker network create demo-network
```

### Create a postgres container

1.  Pull postgres

```
docker pull postgres
```

2. run postgres (create a postgres container)


```
docker run -d --name pg -p 5432:5432 --network demo-network -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=admin123 -e POSTGRES_DB=usersdb postgres
```

3. Admin UI

```
docker run -d --name pgui -p 48080:8080 --network demo-network adminer
```

### Unit Testing

- Mock --> Simulates the behavior of a real object 
- Stub --> Simple implementation of a dependency that provides predefined responses
- Fake --> Lite weight implementation of a dependency that mimics the real implementation but it is very simple and fast (fake payment , in memory database) 

### go test

```
go test -v models/*
```

- run specific test

```
go test -timeout 30s -run ^TestValidate$ demo/models
```

- test CoverProfile of a package

```
go test -timeout 30s -coverprofile=output.out demo/models
```
