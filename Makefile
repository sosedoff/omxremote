dev:
	go build

setup:
	go get

release: clean
	GOOS=linux GOARCH=arm go build

clean:
	rm -f ./omxremote
