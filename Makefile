build:
	go build -o bin/main cmd/tnderlike/main.go

runapp:
	go run cmd/tnderlike/main.go

clean:
	rm bin/*

envgen:
	sh scripts/envgen.sh

run: test runapp

test:
	go test ./...

testapi:
	sh scripts/testapis.sh
