package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Bin struct {
	ID        string    `json:"id"`
	Private   bool      `json:"private"`
	CreatedAt time.Time `json:"createdAt"`
	Name      string    `json:"name"`
}

func CheckErrors(err error) {
	if err != nil {
		fmt.Println("Ошибка:", err)
	}
}

func main() {
	binList := []byte(`{"id": "12345abcde", "private": true, "createdAt": "2024-04-27T15:34:56+03:00", "name": "Пример объекта"}`)

	var bin Bin

	CheckErrors(json.Unmarshal(binList, &bin))

	fmt.Println("Структура:", bin)

}
