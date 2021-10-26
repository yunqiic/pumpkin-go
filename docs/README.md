GO111MODULE="auto" | "off"
下载的包会安装在GOPATH/src下，import导入非标准包的时候也会从这个目录搜索。

GO111MODULE="on"
下载的包会安装在GOPATH/pkg/mod/下，import导入非标准包的时候是从这个目录搜索，不会去GOPATH/src目录下找包。

go mod vendor

go mod protoc File not found