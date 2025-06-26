proto_name = castmer

getway:
	cd ./Get-way && air

castmer:
	cd ./Castmer-Service && air

generate:
	cd ./proto && protoc --go_out=. --go-grpc_out=. ./${proto_name}.proto