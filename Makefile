dev:
	go build

setup:
	go get

release: clean
	GOOS=linux GOARCH=arm go build omxremote.go

clean:
	rm -f ./omxremote
