
all: fmt lint vet build

fmt:
	@echo "Formatting the source code"
	go fmt ./...

lint:
	@echo "Linting the source code"
	# go get -u golang.org/x/lint/golint
	golint ./...

vet:
	@echo "Checking for code issues"
	go vet ./...

build:
	@echo "Removing the Robo binary"
	rm -f bin/robo
	@echo "Building the Robo binary"
	go build -o bin/robo main.go

play:
	@echo "Playing Robo - e.g. command to start PLACE 0,0,NORTH"
	go run main.go

tests:
	@echo "Running Tests"
	go test -timeout 30s  -cover github.com/xplor-robot/pkg/robot/direction
	go test -timeout 30s  -cover github.com/xplor-robot/pkg/robot/roboplay
