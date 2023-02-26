.PHONY: build
build:
	mkdir -p bin && export CGO_ENABLED=0 GOOS=linux && go build -o bin/app -ldflags '-s -w' ./core/main.go



