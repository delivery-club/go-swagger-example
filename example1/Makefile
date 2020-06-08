gen:
	@rm -rf ./cmd
	@rm -rf ./models
	@rm -rf ./restapi
	@swagger generate server \
		--with-context -f ./swagger-api/swagger.yml \
		--name example1

run:
	@go run cmd/example1-server/main.go