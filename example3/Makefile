gen:
	@rm -rf ./cmd
	@rm -rf ./internal
	@mkdir -p cmd
	@mkdir -p internal/generated
	@mkdir -p internal/app
	@mkdir -p internal/config
	@swagger generate server \
		--with-context -f ./swagger-api/swagger.yml \
		-t ./internal/generated -C ./swagger-templates/default-server.yml \
		--template-dir ./swagger-templates/templates \
		--name example3
	@rm -rf ./pkg/example3/client
	@rm -rf ./pkg/example3/models
	@mkdir -p ./pkg/example3
	@swagger generate client -f ./swagger-api/swagger.yml -t ./pkg/example3

run:
	@go run cmd/main.go