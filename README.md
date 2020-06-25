### Примеры использования go-swagger

Официальная документация - [goswagger.io](https://goswagger.io/).

Git-репозиторий - [github.com](https://github.com/go-swagger/go-swagger)

#### Установка go-swagger
Взято из документации, с поправкой на версию:
```shell script
download_url=$(curl -s https://api.github.com/repos/go-swagger/go-swagger/releases/tags/v0.24.0 | \
  jq -r '.assets[] | select(.name | contains("'"$(uname | tr '[:upper:]' '[:lower:]')"'_amd64")) | .browser_download_url')
curl -o /usr/local/bin/goswagger -L'#' "$download_url"
chmod +x /usr/local/bin/goswagger
```

#### Содержание
- [Пример 1](https://github.com/delivery-club/go-swagger-example/tree/master/example1) - простая генерация http-сервера из swagger-спецификации.
- [Пример 2](https://github.com/delivery-club/go-swagger-example/tree/master/example2) - генерация http-сервера с использованием кастомного шаблона.
- [Пример 3](https://github.com/delivery-club/go-swagger-example/tree/master/example3) - генерация слиента для http-сервиса.
