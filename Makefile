# Set the root project path and release folder path
ROOT_DIR := $(shell pwd)

# Build All Proto
proto: clean
	@echo "TASK: Compile All Proto"
	@find api/ -type f -name "*.proto" | xargs protoc --go_out=./model --go-grpc_out=./model
	@find model/ -type f -name "*.pb.go" | xargs -I {} protoc-go-inject-tag --input={}


install:
	@clear
	@echo "TASK: Installing Go Mod"
	@go mod tidy

# Clean built binaries
clean: install
	@echo "TASK: Cleaning up..."
	@rm -rf model/*


