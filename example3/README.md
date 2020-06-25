#### Генерация слиента для http-сервиса
Для примера будет использована та же swagger-спека, что и в 1м и 2м примерах ([swagger-api/swagger.yml](https://github.com/delivery-club/go-swagger-example/blob/master/example3/swagger-api/swagger.yml)).

Генерация клиента для http-сервиса:
```shell script
goswagger generate client -f ./swagger-api/swagger.yml -t ./pkg/example3
```

Пример использования:
```go
package main

import (
	"context"

	"github.com/delivery-club/go-swagger-example/example3/pkg/example3/client"
	"github.com/delivery-club/go-swagger-example/example3/pkg/example3/client/hello"
	httptransport "github.com/go-openapi/runtime/client"
)

func main() {
	t := httptransport.New("service-host", "/", []string{"http"})
	example3client := client.New(t, nil)

	reqParams := hello.NewHelloWorldParams().
		WithContext(context.Background())

	response, err := example3client.Hello.HelloWorld(reqParams)
	if err != nil {
		// handle error
	}

	// handle response
}
```

При указании правильного тэга, можно удобно работать с версиями клиентов.
В нашем случае корректным тэгом будет `example3/pkg/example3/v0.0.1`.
go.mod для пользователя этого клиента будет следующий:
```
module github.com/delivery-club/foo

go 1.13

require (
	github.com/delivery-club/go-swagger-example/example3/pkg/example3 v0.0.1
	github.com/go-openapi/runtime v0.19.11
)
```