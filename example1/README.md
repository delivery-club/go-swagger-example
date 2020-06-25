#### Простая генерация http-сервера из swagger-спецификации
Для демонстрации будем использовать следующую простую схему ([swagger-api/swagger.yml](https://github.com/delivery-club/go-swagger-example/blob/master/example1/swagger-api/swagger.yml)):
```yaml
swagger: "2.0"
info:
  title: Example service
  description: Example service
  version: 0.0.1
host: localhost
schemes:
  - http
basePath: /
consumes:
  - application/json
produces:
  - application/json

paths:
  /hello:
    get:
      tags:
        - Hello
        - World
      summary: Example Route
      operationId: HelloWorld
      description: Some description
      responses:
        200:
          description: successful response
          schema:
            $ref: '#/definitions/HelloWorld'
        400:
          description: Bad request
          schema:
            $ref: '#/definitions/Error'
        500:
          description: Internal server error
          schema:
            $ref: '#/definitions/Error'
  /hello-world:
    get:
      tags:
        - Hello
        - World
      summary: Example Route
      operationId: HelloWorldFull
      description: Some description
      responses:
        200:
          description: successful response
          schema:
            $ref: '#/definitions/HelloWorld'
        400:
          description: Bad request
          schema:
            $ref: '#/definitions/Error'
        500:
          description: Internal server error
          schema:
            $ref: '#/definitions/Error'

definitions:
  Error:
    type: object
    required:
      - message
      - code
    description: Error message
    properties:
      message:
        type: string
      code:
        type: integer

  HelloWorld:
    type: object
    required:
      - message
    properties:
      message:
        type: string
```

Для генерации кода http-сервиса нем необходимо выполнить следующую команду:
```shell script
goswagger generate server \
    --with-context -f ./swagger-api/swagger.yml \
    --name example1
```

Для запуска сгенерированного http-сервиса:
```shell script
go run cmd/example1-server/main.go
2020/06/08 21:30:34 Serving example1 at http://127.0.0.1:59598
```

Проверяем (порт может меняться):
```shell script
curl http://localhost:59598/hello-world
"operation HelloHelloWorldFull has not yet been implemented"
```