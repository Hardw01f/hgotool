
build:
	go build -o bin/hgotool
	GOOS=linux GOARCH=amd64 go build -o bin/linux/hgotool
	GOOS=windows GOARCH=amd64 go build -o bin/windows/hgotool_amd64
