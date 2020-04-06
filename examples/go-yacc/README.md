```sh
## setup
go get golang.org/x/tools/cmd/goyacc
## run example
goyacc -o parser.go parser.go.y ; go run parser.go '5 - 2 * 3'
```
