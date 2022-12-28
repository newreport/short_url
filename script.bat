 
  go env -w GOPROXY=https://goproxy.cn,direct
 go install github.com/beego/bee/v2@latest
 
 go env -w CGO_CFLAGS="-g -O2 -Wno-return-local-addr"