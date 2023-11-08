
## Инструкции по запуску

`make grpc-redis` - Запуск grpc сервера с хранилищем Redis

`make grpc-postgres` - Запуск grpc сервера с хранилищем Postgres

`make http-redis` - Запуск grpc сервера с хранилищем Redis

`make http-postgres` - Запуск grpc сервера с хранилищем Postgres

`make tests` - Запуск тестов

## API 

### GRPC
```protobuf
syntax = "proto3";

package url;

option go_package = "url.v1";

service UrlShortenService {
  rpc CreateShortUrl(OriginalUrlRequest) returns (ShortUrlResponse);
  rpc GetOriginalUrl(ShortUrlRequest) returns (OriginalUrlResponse);
}

message OriginalUrlRequest {
  string URL = 1;
}

message ShortUrlResponse {
  string URL = 1;
}

message OriginalUrlResponse {
  string URL = 1;
}

message ShortUrlRequest  {
  string URL = 1;
}
```

### HTTP

`POST /api/create` - создание короткой ссылки

Запрос:
```json
{
	"original_url":"https://golangbyexample.com/net-http-pckage-get-query-params-golang/"
}
```

Ответ:
```json
{
	"short_url": "https://golangbyexample.com/M2X3EOktia"
}
```

`POST /api/get/{shortURL}` - получение оригинальной ссылки

Ответ:
```json
{
	"original_url": "https://golangbyexample.com/net-http-pckage-get-query-params-golang/"
}
```