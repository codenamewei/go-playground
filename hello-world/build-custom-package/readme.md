### structure of a simple go project

```
myproject/
├── main.go
└── mypackage/
    └── add.go
```

this creates go.mod in `/build-custom-package`

```
go mod init github.com/codenamewei/go-playground/hello-world/build-custom-package
```

```
go build -o build-custom-package
./build-custom-package
```
