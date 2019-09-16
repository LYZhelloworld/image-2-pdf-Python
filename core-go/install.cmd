@echo off
set GOPATH=%cd%
go get github.com/jung-kurt/gofpdf
go install image2pdf
