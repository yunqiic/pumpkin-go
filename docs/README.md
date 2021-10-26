GO111MODULE="auto" | "off"
下载的包会安装在GOPATH/src下，import导入非标准包的时候也会从这个目录搜索。

GO111MODULE="on"
下载的包会安装在GOPATH/pkg/mod/下，import导入非标准包的时候是从这个目录搜索，不会去GOPATH/src目录下找包。

go mod vendor

go mod protoc File not found

find ~/dev/go -name "go-proto-validators"
~/dev/go/pkg/mod/cache/download/github.com/mwitkow/go-proto-validators

cp -r ~/dev/go/pkg/mod/cache/download/github.com/mwitkow ~/dev/go/src/github.com

cd ${GOPATH}/src/github.com/mwitkow
git clone git@github.com:mwitkow/go-proto-validators.git

go get github.com/gorilla/mux
go get github.com/go-kratos/kratos
go get github.com/go-kratos/kratos/v2

go get github.com/elazarl/go-bindata-assetfs