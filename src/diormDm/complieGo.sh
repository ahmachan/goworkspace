gofmt -w ./main.go
go tool vet ./main.go
go build -v ./main.go
