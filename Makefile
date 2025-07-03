proto_name = subscription
file_name = Subscription-Service 
getway:
	cd ./Get-way && air

castmer:
	cd ./Castmer-Service && air

generate:
	cd ./proto && protoc --go_out=. --go-grpc_out=. ./${proto_name}.proto


subscription :
	cd ./Subscription-Service/ && air

notification :
	cd ./Notification-Service/ && npx nodemon server.js

work:
	cd ./${file_name} && go mod init github.com/microservic/${proto_name}
	cd .. && go work init ./${file_name}


