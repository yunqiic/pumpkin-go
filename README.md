# pumpkin-go
pumpkin-go

```shell
protoc --go_out=plugins=grpc:./ --go_opt=Msimple.proto=./ ./simple.proto
protoc --go-grpc_out=./ --go-grpc_opt=Msimple.proto=./ ./simple.proto

protoc --go_out=./ --go-grpc_out=./ --go_opt=Msimple.proto=./ --go-grpc_opt=Msimple.proto=./ ./simple.proto

protoc --go_out=./ --go-grpc_out=./ ./proto/*.proto

protoc --proto_path=. --proto_path=${GOPATH}/src --govalidators_out=. --go_out=./ --go-grpc_out=./ ./proto/*.proto
protoc --proto_path=. --proto_path=${GOPATH}/src --go_out=./ --go-grpc_out=./ ./proto/*.proto

protoc --proto_path=. --proto_path=${GOPATH}/src --govalidators_out=. --go_out=./ --go-grpc_out=./ --go-http_out=./ --openapiv2_out=./ ./proto/*.proto

go get github.com/mwitkow/go-proto-validators
go install github.com/mwitkow/go-proto-validators/protoc-gen-govalidators
go get github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway

go install \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
    google.golang.org/protobuf/cmd/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc
    
go run server.go
go run client.go

openssl genrsa -out server.key 2048
openssl req -new -x509 -sha256 -key server.key -out server.pem

Country Name (2 letter code) []:CN
State or Province Name (full name) []:Beijing
Locality Name (eg, city) []:Beijing
Organization Name (eg, company) []:yunqiic
Organizational Unit Name (eg, section) []:dev
Common Name (eg, fully qualified host name) []:yunqiic.com
Email Address []:admin@yunqiic.com

openssl ecparam -genkey -name secp384r1 -out server.key
openssl req -new -x509 -sha256 -key server.key -out server.pem -days 3650

https://www.cnblogs.com/jackluo/p/13841286.html

openssl req -new -nodes -keyout server.key -out server.csr -config openssl.cnf

Country Name (2 letter code) [AU]:CN
State or Province Name (full name) [Some-State]:Beijing
Locality Name (eg, city) []:Beijing
Organization Name (eg, company) [Internet Widgits Pty Ltd]:yunqiic
Organizational Unit Name (eg, section) []:dev
Common Name (e.g. server FQDN or YOUR name) []:*.yunqiic.com
Email Address []:admin@yunqiic.com

Please enter the following 'extra' attributes
to be sent with your certificate request
A challenge password []:grpc
An optional company name []:yunqiic

openssl genrsa -out ca.key 2048

openssl req -x509 -new -nodes -key ca.key -subj "/CN=*.yunqiic.com" -days 3650 -out ca.crt

openssl req -new -sha256 \
    -key ca.key \
    -subj "/C=CN/ST=Beijing/L=Beijing/O=yunqiic/OU=dev/CN=*.yunqiic.com" \
    -reqexts SAN \
    -config <(cat /System/Library/OpenSSL/openssl.cnf \
        <(printf "[SAN]\nsubjectAltName=DNS:*.yunqiic.com,DNS:yunqiic.com")) \
    -out test.csr
    
openssl x509 -req -days 365000 \
    -in test.csr -CA ca.crt -CAkey ca.key -CAcreateserial \
    -extfile <(printf "subjectAltName=DNS:*.yunqiic.com,DNS:yunqiic.com") \
    -out test.crt
    
openssl x509 -text -noout -in test.crt

go get github.com/grpc-ecosystem/go-grpc-middleware

展示两个版本之间改动细节

git diff sdafsfsfgf4  sdfsfdhkllg
 

 

展示两个版本之间改动文件列表名,不展示细节

git diff --name-status sdafsfsfgf4  sdfsfdhkllg

Git 对比两个版本间某一个文件的变化
先列出两个版本间发生更改的文件列表

git diff commit1 commit2 --stat --name-only
查看指定文件在两个版本间发生的变更

git diff commit1 commit2 -- somefile.js
如果感觉这种显示不够直观，可以使用 vimdiff 查看

git difftool commit1 commit2 -- somefile.js
```

```
k8s api go
https://www.cnblogs.com/you-men/p/14551630.html
```

```
protoc --help
Usage: protoc [OPTION] PROTO_FILES
Parse PROTO_FILES and generate output based on the options given:
  -IPATH, --proto_path=PATH   Specify the directory in which to search for
                              imports.  May be specified multiple times;
                              directories will be searched in order.  If not
                              given, the current working directory is used.
                              If not found in any of the these directories,
                              the --descriptor_set_in descriptors will be
                              checked for required proto file.
  --version                   Show version info and exit.
  -h, --help                  Show this text and exit.
  --encode=MESSAGE_TYPE       Read a text-format message of the given type
                              from standard input and write it in binary
                              to standard output.  The message type must
                              be defined in PROTO_FILES or their imports.
  --deterministic_output      When using --encode, ensure map fields are
                              deterministically ordered. Note that this order
                              is not canonical, and changes across builds or
                              releases of protoc.
  --decode=MESSAGE_TYPE       Read a binary message of the given type from
                              standard input and write it in text format
                              to standard output.  The message type must
                              be defined in PROTO_FILES or their imports.
  --decode_raw                Read an arbitrary protocol message from
                              standard input and write the raw tag/value
                              pairs in text format to standard output.  No
                              PROTO_FILES should be given when using this
                              flag.
  --descriptor_set_in=FILES   Specifies a delimited list of FILES
                              each containing a FileDescriptorSet (a
                              protocol buffer defined in descriptor.proto).
                              The FileDescriptor for each of the PROTO_FILES
                              provided will be loaded from these
                              FileDescriptorSets. If a FileDescriptor
                              appears multiple times, the first occurrence
                              will be used.
  -oFILE,                     Writes a FileDescriptorSet (a protocol buffer,
    --descriptor_set_out=FILE defined in descriptor.proto) containing all of
                              the input files to FILE.
  --include_imports           When using --descriptor_set_out, also include
                              all dependencies of the input files in the
                              set, so that the set is self-contained.
  --include_source_info       When using --descriptor_set_out, do not strip
                              SourceCodeInfo from the FileDescriptorProto.
                              This results in vastly larger descriptors that
                              include information about the original
                              location of each decl in the source file as
                              well as surrounding comments.
  --dependency_out=FILE       Write a dependency output file in the format
                              expected by make. This writes the transitive
                              set of input file paths to FILE
  --error_format=FORMAT       Set the format in which to print errors.
                              FORMAT may be 'gcc' (the default) or 'msvs'
                              (Microsoft Visual Studio format).
  --fatal_warnings            Make warnings be fatal (similar to -Werr in
                              gcc). This flag will make protoc return
                              with a non-zero exit code if any warnings
                              are generated.
  --print_free_field_numbers  Print the free field numbers of the messages
                              defined in the given proto files. Groups share
                              the same field number space with the parent
                              message. Extension ranges are counted as
                              occupied fields numbers.
  --plugin=EXECUTABLE         Specifies a plugin executable to use.
                              Normally, protoc searches the PATH for
                              plugins, but you may specify additional
                              executables not in the path using this flag.
                              Additionally, EXECUTABLE may be of the form
                              NAME=PATH, in which case the given plugin name
                              is mapped to the given executable even if
                              the executable's own name differs.
  --cpp_out=OUT_DIR           Generate C++ header and source.
  --csharp_out=OUT_DIR        Generate C# source file.
  --java_out=OUT_DIR          Generate Java source file.
  --js_out=OUT_DIR            Generate JavaScript source.
  --kotlin_out=OUT_DIR        Generate Kotlin file.
  --objc_out=OUT_DIR          Generate Objective-C header and source.
  --php_out=OUT_DIR           Generate PHP source file.
  --python_out=OUT_DIR        Generate Python source file.
  --ruby_out=OUT_DIR          Generate Ruby source file.
  @<filename>                 Read options and filenames from file. If a
                              relative file path is specified, the file
                              will be searched in the working directory.
                              The --proto_path option will not affect how
                              this argument file is searched. Content of
                              the file will be expanded in the position of
                              @<filename> as in the argument list. Note
                              that shell expansion is not applied to the
                              content of the file (i.e., you cannot use
                              quotes, wildcards, escapes, commands, etc.).
                              Each line corresponds to a single argument,
                              even if it contains spaces.
```