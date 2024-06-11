generate:
	rm -rf gen
	mkdir gen
	protoc -I ./proto \
	--go_out ./gen --go_opt paths=source_relative \
	--go-grpc_out ./gen --go-grpc_opt paths=source_relative \
	--grpc-gateway_out ./gen --grpc-gateway_opt paths=source_relative \
	./proto/**/*/*.proto

	rm -rf gen/openapiv2
	mkdir gen/openapiv2
	protoc -I ./proto \
	--openapiv2_out ./gen/openapiv2 \
    proto/vlr/api/api_service.proto