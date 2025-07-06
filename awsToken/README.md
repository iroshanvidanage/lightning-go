# Go module for AWS sts authentication

This test module is for AWS STS token generation.

- Need to run the following commands

```go
go mod init <proj_name>

go get github.com/aws/aws-sdk-go-v2/aws
go get github.com/aws/aws-sdk-go-v2/config
go get github.com/aws/aws-sdk-go-v2/service/sts

// Add following to the import block
"github.com/aws/aws-sdk-go-v2/aws"
"github.com/aws/aws-sdk-go-v2/config"
"github.com/aws/aws-sdk-go-v2/service/sts"
```

## Create a Go module

- `go mod init example.com/go-module-name`
- `example.com/go-module-name` can be replaces with the desired module path.
- multiple packages can be within a go module
- `go mod init` initalizes the dir as a module.