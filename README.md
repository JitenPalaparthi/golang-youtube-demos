go.dev
VSCode

two types of projects

1. binary/executable
2. packages/libraries 

### create a go mod

```bash
go mod init demo
```

### What would happen when you run go application

```bash
go run main.go
```

1. It is compiled
2. Assembled (converted to assembly code in .o file)
3. Linked ( all the necessary dlls or .o files or .so files or .a files are linked)
4. Binary is created
5. That binary is run

### Build the application

```bash
go build main.go
go build -o demo main.go
```

1. It is compiled
2. Assembled
3. Linked
4. Binary is created

## Keywords 

const, case, default, if ,else, switch, var,import, package (9 out of 25)

## builtin

print, println, complex (3 out of  18)


## go env 

1. GOROOT  -> The actual installation of Go
2. GOPATH  -> All third party or user defined libraries reside
3. GOBIN   -> All binaries installed using go install stored in this directory 
4. GOOS    -> Go OS 
5. GOARCH  -> Go Architecture

- all cross compilation os/archs

```bash
go tool dist list
```

- To cross compilation and build 

```bash
GOOS=windows GOARCH=amd64 go build -o buil
d/switch-case-demo-win-amd64.exe main.go

GOOS=linux GOARCH=amd64 
go build -o build/switch-case-demo-linux-amd64 main.go
```

- release build 

```bash
GOOS=windows GOARCH=amd64 go build -ldflags="-w -s"  -o build/switch-case-demo-release-win-amd64.exe main.go
```

- Type of applications

1. Binary/executable --> main package and main func , it is a binary
2. Library/package   --> no main but other packages , it is library
 
### To get a package 

```
go get github.com/JitenPalaparthi/shapes1
```

- tidy

```
go mod tidy
```

- download

```
go mod download
```

- vendor 

```
go mod vendor
```