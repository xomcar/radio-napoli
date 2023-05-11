build: clean mac win
run:
	go run main.go
mac:
	CGO_ENABLED=1 GOARCH=arm64 GOOS=darwin go build -ldflags "-s -w" -o bin/mac/napoli-arm64 main.go
	CGO_ENABLED=1 GOARCH=amd64 GOOS=darwin go build -ldflags "-s -w" -o bin/mac/napoli-amd64 main.go
win:
	CGO_ENABLED=1 GOARCH=amd64 GOOS=windows go build -ldflags "-s -w" -o bin/win/napoli.exe
clean:
	rm -rf bin