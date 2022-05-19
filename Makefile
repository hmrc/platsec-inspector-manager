compile: ./cmd/InspectorManager/main.go
	go build -o ./cmd/InspectorManager/inspector ./cmd/InspectorManager/main.go

build: compile
	mkdir -p bin
	mv ./cmd/InspectorManager/inspector ./bin

format: ./cmd/InspectorManager/main.go
	gofmt .

test: ./inspector/inspector2_test.go
	go test ./inspector

test-coverage: ./inspector
	go test -coverprofile cover.out ./...
	go tool cover -html=cover.out
