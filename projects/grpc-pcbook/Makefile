gen:
	protoc --proto_path=proto --go_out=pb proto/*.proto

clean:
	rm -r pb
	mkdir pb

test:
	go test ./... -cover -v

run:
	go run main.go

.PHONY: gen clean run test