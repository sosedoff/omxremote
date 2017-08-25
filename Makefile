BINDATA=

dev: dev-assets
	go build

assets:
	go-bindata $(BINDATA) -ignore=\\.gitignore -ignore=\\.DS_Store -ignore=\\.gitkeep static/...

dev-assets:
	@$(MAKE) --no-print-directory assets BINDATA="-debug"

setup:
	go get github.com/jteeuwen/go-bindata/...
	go get github.com/stretchr/testify/assert
	go get

test:
	go test ./...

release: clean assets
	GOOS=linux GOARCH=arm go build

clean:
	rm -f ./omxremote
