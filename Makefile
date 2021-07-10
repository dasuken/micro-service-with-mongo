generate-grpc:
	protoc --go_out=. --go-grpc_out=. ./messages/*.proto