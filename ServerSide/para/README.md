```base
go env
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct
```

```
go mod init
```

```go

name := c.Query("name"")  // name=elliot
age := c.Param("") // path/elliot  
```
