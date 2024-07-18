build: 
	go build -o bin/exchange.exe
run: build
	./bin/exchange
test:
	go test -v ./...