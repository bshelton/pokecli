MAIN=./cmd/main/main.go
OUTPUT_NAME=pokecli

build:
	go build -o bin/pokecli ${MAIN}

run:
	go run ${MAIN}

compile:
	echo "Compiling for multiple OS and Platform"
	GOOS=darwin GOARCH=amd64 go build -o bin/${OUTPUT_NAME}-darwin-amd64 ${MAIN}
	GOOS=darwin GOARCH=arm64 go build -o bin/${OUTPUT_NAME}-darwin-arm64 ${MAIN}
	GOOS=linux GOARCH=386 go build -o bin/${OUTPUT_NAME}-linux-386 ${MAIN}
	GOOS=linux GOARCH=amd64 go build -o bin/${OUTPUT_NAME}-linux-amd64 ${MAIN}
	GOOS=windows GOARCH=386 go build -o bin/${OUTPUT_NAME}-windows-386 ${MAIN}