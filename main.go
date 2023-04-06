package main

import (
	"RunLengthEncoding/utils"
	"fmt"
	"log"
	"strconv"
	"sync"

	"github.com/gofiber/fiber"
	jsoniter "github.com/json-iterator/go"
)

const LEN_CHUNK = 3

type Storage struct {
	m  map[int][]string
	mu sync.Mutex
	wg sync.WaitGroup
}

type Msg struct {
	Data []string `json:"data"`
}

func RunLengthEncode(msg []string) []string {
	var ofset int
	var part int
	storage := Storage{m: make(map[int][]string)}

	for i := 0; i < len(msg); i += LEN_CHUNK {
		limit := ofset + LEN_CHUNK
		if limit > len(msg) {
			limit = len(msg)
		}
		storage.wg.Add(1)
		go encode(msg[ofset:limit], part, &storage)
		part++
		ofset += LEN_CHUNK
		limit += LEN_CHUNK
	}
	storage.wg.Wait()
	res := make([]string, 0, len(msg))
	for i := 0; i < part; i++ {
		res = append(res, storage.m[i]...)
	}
	return res
}

func encode(msg []string, part int, storage *Storage) {
	var firsrElement rune
	count := 1
	res := make([]string, len(msg))
	for i := range msg {
		word := msg[i]
		for _, elem := range word {
			if firsrElement == 0 {
				firsrElement = elem
			} else if elem == firsrElement {
				count++
				continue
			} else if elem != firsrElement && count > 1 {
				res[i] += strconv.Itoa(count) + string(firsrElement)
				count = 1
			} else if elem != firsrElement && count == 1 {
				res[i] += string(firsrElement)
			}
			firsrElement = elem
		}
		if count > 1 {
			res[i] += strconv.Itoa(count) + string(firsrElement)
			count = 1
		} else {
			res[i] += string(firsrElement)
		}
		firsrElement = 0
	}
	storage.mu.Lock()
	defer storage.mu.Unlock()
	defer storage.wg.Done()
	storage.m[part] = res
}

func RunLengthDecode(msg []string) []string {
	var ofset int
	var part int
	storage := Storage{m: make(map[int][]string)}

	for i := 0; i < len(msg); i += LEN_CHUNK {
		limit := ofset + LEN_CHUNK
		if limit > len(msg) {
			limit = len(msg)
		}

		storage.wg.Add(1)
		go decode(msg[ofset:limit], part, &storage)
		part++
		ofset += LEN_CHUNK
		limit += LEN_CHUNK
	}
	storage.wg.Wait()
	res := make([]string, 0, len(msg))
	for i := 0; i < part; i++ {
		res = append(res, storage.m[i]...)
	}
	return res
}

func decode(msg []string, part int, storage *Storage) {
	var numberStr string
	var number int
	var secondElement string
	res := make([]string, len(msg))
	for i := range msg {
		for _, elem := range msg[i] {
			if elem >= 48 && elem <= 57 && secondElement == "" {
				numberStr += string(elem)
			} else if secondElement == "" && numberStr != "" {
				secondElement = string(elem)
			} else if elem == 32 || utils.NotNumber(secondElement) && numberStr != "" {
				number, _ = strconv.Atoi(numberStr) // обработать ошибку
				for j := 0; j < number; j++ {
					res[i] += secondElement
				}
				if utils.NotNumber(string(elem)) {
					res[i] += string(elem)
					numberStr = ""
				} else {
					numberStr = string(elem)
				}
				secondElement = ""
			} else {
				res[i] += string(elem)
			}
		}
		number, _ = strconv.Atoi(numberStr) // обработать ошибку
		for j := 0; j < number; j++ {
			res[i] += secondElement
		}
		secondElement = ""
		numberStr = ""
	}
	storage.mu.Lock()
	defer storage.mu.Unlock()
	defer storage.wg.Done()
	storage.m[part] = res
}

func EncodeHandler(c *fiber.Ctx) {
	bodyByte := c.Fasthttp.Request.Body()
	msg := Msg{}
	jsoniter.Unmarshal(bodyByte, &msg)
	res := RunLengthEncode(msg.Data)
	c.Write(res)
}

func DecodeHandler(c *fiber.Ctx) {
	bodyByte := c.Fasthttp.Request.Body()
	msg := Msg{}
	jsoniter.Unmarshal(bodyByte, &msg)
	res := RunLengthDecode(msg.Data)
	c.Write(res)
}

// решил оставить все в одном файле, не большое приложение, так удoбнее будет смотреть
// тесты сделал по быстрому, по хорошему добавить еще несколько кейсов для проверки и заюзать testify
// возможно по коду если пройтись свежим взглядом можно пооптимизировать
// по web-серверу по хорошему нужно добавить валидацию
// не обрабатывал ошибки
func main() {
	// проверка на множестве Мандельброта
	msg := utils.CreateMandelbrot()
	fmt.Println(msg)
	code := RunLengthEncode(msg)
	fmt.Println(code)
	fmt.Println(RunLengthDecode(code))

	// web сервер

	app := fiber.New()
	app.Post("/encode", EncodeHandler)
	app.Post("/decode", DecodeHandler)
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
