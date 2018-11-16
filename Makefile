
build:
	go build -o bin/hgotool
	GOOS=linux GOARCH=amd64 go build -o bin/linux/hgotool_amd64
	GOOS=linux GOARCH=arm64 go build -o bin/linux/hgotool_arm64
	GOOS=windows GOARCH=amd64 go build -o bin/windows/hgotool_amd64

