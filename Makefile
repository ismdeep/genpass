.PHONY: help
help:
	@cat Makefile | grep '# `' | grep -v '@cat Makefile'

# `make build`
.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux   GOARCH=amd64 go build -o build/genpass_linux_amd64       -mod vendor -trimpath -ldflags '-s -w' .
	CGO_ENABLED=0 GOOS=linux   GOARCH=arm64 go build -o build/genpass_linux_arm64       -mod vendor -trimpath -ldflags '-s -w' .
	CGO_ENABLED=0 GOOS=darwin  GOARCH=amd64 go build -o build/genpass_darwin_amd64      -mod vendor -trimpath -ldflags '-s -w' .
	CGO_ENABLED=0 GOOS=darwin  GOARCH=arm64 go build -o build/genpass_darwin_arm64      -mod vendor -trimpath -ldflags '-s -w' .
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o build/genpass_windows_amd64.exe -mod vendor -trimpath -ldflags '-s -w' .
	CGO_ENABLED=0 GOOS=windows GOARCH=arm64 go build -o build/genpass_windows_arm64.exe -mod vendor -trimpath -ldflags '-s -w' .

# `make clean`
.PHONY: clean
clean:
	rm -rf build/
