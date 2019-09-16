#/bin/sh
export GOPATH=$PWD
go get github.com/jung-kurt/gofpdf
go install image2pdf
