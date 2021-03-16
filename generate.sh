go generate ./...
cp ./internal/generator/gchalkgen/gchalkgen.go.txt ./internal/generator/gchalkgen/gchalkgen.go
go run ./internal/generator/gchalkgen/gchalkgen.go > generated.go
gofmt -s -w generated.go