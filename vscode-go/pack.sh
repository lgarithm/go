# https://github.com/Microsoft/vscode-go/blob/master/src/goInstallTools.ts#L20

export GOPATH=$HOME/vscode-go
mkdir -p $GOPATH

go get -v -u -d github.com/acroca/go-symbols
go get -v -u -d github.com/alecthomas/gometalinter
go get -v -u -d github.com/cweill/gotests/...
go get -v -u -d github.com/davidrjenni/reftools/cmd/fillstruct
go get -v -u -d github.com/derekparker/delve/cmd/dlv
go get -v -u -d github.com/fatih/gomodifytags
go get -v -u -d github.com/golang/lint/golint
go get -v -u -d github.com/haya14busa/goplay/cmd/goplay
go get -v -u -d github.com/josharian/impl
go get -v -u -d github.com/nsf/gocode
go get -v -u -d github.com/ramya-rao-a/go-outline
go get -v -u -d github.com/rogpeppe/godef
go get -v -u -d github.com/sourcegraph/go-langserver
go get -v -u -d github.com/sqs/goreturns
go get -v -u -d github.com/tylerb/gotype-live
go get -v -u -d github.com/uudashr/gopkgs/cmd/gopkgs
go get -v -u -d github.com/zmb3/gogetdoc
go get -v -u -d golang.org/x/tools/cmd/godoc
go get -v -u -d golang.org/x/tools/cmd/goimports
go get -v -u -d golang.org/x/tools/cmd/gorename
go get -v -u -d golang.org/x/tools/cmd/guru
go get -v -u -d honnef.co/go/tools/...

cd $GOPATH
tar -cf src.tar src
