proto:
	protoc internal/**/pb/*.proto --go_out=. --go_opt=paths=import --go-grpc_out=. --go-grpc_opt=paths=import

server:
	go run cmd/main.go