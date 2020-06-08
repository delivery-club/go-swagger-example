#### Генерация http-сервера с использованием кастомного шаблона
Для примера будет использована та же swagger-спека, что и в 1м примере ([swagger-api/swagger.yml](https://github.com/delivery-club/go-swagger-example/blob/master/example2/swagger-api/swagger.yml)).

Для генерации кода http-сервиса нем необходимо выполнить следующую команду:
```shell script
swagger generate server \
		-f ./swagger-api/swagger.yml \
		-t ./internal/generated -C ./swagger-templates/default-server.yml \
		--template-dir ./swagger-templates/templates \
		--name example2
```

Структура проекта будет следующая:
```
.
├── Makefile
├── README.md
├── cmd
│   └── main.go
├── go.mod
├── go.sum
├── internal
│   ├── app
│   │   ├── app.go
│   │   ├── hello_world.go
│   │   └── hello_world_full.go
│   ├── config
│   │   └── config.go
│   └── generated
│       ├── models
│       │   ├── error.go
│       │   └── hello_world.go
│       └── restapi
│           ├── configure_example2.go
│           ├── embedded_spec.go
│           ├── operations
│           │   ├── example2_api.go
│           │   └── hello
│           │       ├── hello_world.go
│           │       ├── hello_world_full.go
│           │       ├── hello_world_full_parameters.go
│           │       ├── hello_world_full_responses.go
│           │       ├── hello_world_parameters.go
│           │       └── hello_world_responses.go
│           └── server.go
├── swagger-api
│   └── swagger.yml
└── swagger-templates
    ├── default-server.yml
    └── templates
        └── server
            ├── app.gotmpl
            ├── config.gotmpl
            ├── handler.gotmpl
            └── main.gotmpl

```

Для запуска сгенерированного http-сервиса:
```shell script
go run cmd/main.go
2020/06/08 21:39:12 Serving example2 at http://[::]:8001
```

Проверяем (порт может меняться):
```shell script
curl http://localhost:8001/hello-world
"operation hello HelloWorldFull has not yet been implemented"
```