build:
	go build -o bin/main cmd/tnderlike/main.go

runapp:
	go run cmd/tnderlike/main.go

clean:
	rm bin/*

run: test sec runapp

check: test sec

test:
	go test ./...

sec:
	gosec ./...
