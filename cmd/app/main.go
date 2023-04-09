package main

import (
	"RunLengthEncoding/internal/transport/rest"
	"log"

	"github.com/gofiber/fiber"
)

// решил оставить все в одном файле, не большое приложение, так удoбнее будет смотреть
// тесты сделал по быстрому, по хорошему добавить еще несколько кейсов для проверки и заюзать testify
// возможно по коду если пройтись свежим взглядом можно пооптимизировать
// по web-серверу по хорошему нужно добавить валидацию
// не обрабатывал ошибки
func main() {
	app := fiber.New()
	app.Post("/encode", rest.EncodeHandler)
	app.Post("/decode", rest.DecodeHandler)
	log.Fatal(app.Listen(":3000"))
}

/*
ПРИМЕРЫ ЗАПРОСОВ

curl --location '127.0.0.1:3000/encode' \
--header 'Content-Type: application/json' \
--data '{
    "data": [
        "AAAAA",
        "AAA BBB",
        "ABC DDD",
        "     ",
        "A B C",
        "ABC"
    ]
}'

curl --location '127.0.0.1:3000/decode' \
--header 'Content-Type: application/json' \
--data '{
    "data": [
        "5A",
        "3A 3B",
        "ABC 3D",
        "5 ",
        "A B C",
        "ABC"
    ]
}'

*/
