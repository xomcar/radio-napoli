build: mac win linux
run:
	go run main.go
mac:
	GOARCH=arm64 GOOS=darwin go build -ldflags "-s -w" -o bin/mac/napoli main.go
win:
	GOARCH=amd64 GOOS=windows go build -ldflags "-s -w" -o bin/win/napoli.exe
linux:
	GOARCH=amd64 GOOS=linux go build -ldflags "-s -w" -o bin/linux/napoli
clean:
	rm -rf bin