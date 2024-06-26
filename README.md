# DiosProcGo

build protocbuf cmd:
  buf build
  buf generate

gen liberary from define and dll for gcc
  dlltool -D libDiosProc.dll -d libDiosProc.def -l libDiosProc.a -k

build 
   go env -w CGO_LDFLAGS="-O2 -g -Lg:/diosproc/src/cdevice -llibDiosProc"
   go build -ldflags "-X google.golang.org/protobuf/reflect/protoregistry.conflictPolicy=warn" -o DiosProcGo.exe .\main.go
