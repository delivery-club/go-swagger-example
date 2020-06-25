gen:
	@rm -rf ./cmd
	@rm -rf ./internal
	@mkdir -p cmd
	@mkdir -p internal/generated
	@mkdir -p internal/app
	@mkdir -p internal/config
	@goswagger generate server \
		-f ./swagger-api/swagger.yml \
		-t ./internal/generated -C ./swagger-templates/default-server.yml \
		--template-dir ./swagger-templates/templates \
		--name example2

run:
	@go run cmd/main.go