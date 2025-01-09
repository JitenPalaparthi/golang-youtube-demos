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