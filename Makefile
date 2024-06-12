build:
	go build -o bin/main cmd/tnderlike/main.go

runapp:
	go run cmd/tnderlike/main.go

clean:
	rm bin/*

run: test runapp

check: test

test:
	go test ./...

rundocker:
	docker-compose up
