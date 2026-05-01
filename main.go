package main

import (
	"app3/api"
	"app3/bins"
	"app3/file"
	"app3/storage"
	"fmt"
)

func main() {
	// binList := []byte(`{"id": "12345abcde", "private": true, "createdAt": "2024-04-27T15:34:56+03:00", "name": "Пример объекта"}`)

	// Создаем новый Bin
	bin1 := bins.NewBin("12345abcde", true, "2024-04-27T15:34:56+03:00", "Пример объекта 1")
	bin2 := bins.NewBin("67890fghij", false, "2024-04-27T15:34:56+03:00", "Пример объекта 2")

	// Создаем список Bin
	binList := bins.NewBinList()

	// Добавляем Bin в список
	binList.AddBin(bin1)
	binList.AddBin(bin2)

	fmt.Println(binList)

	// Trial use of libraries
	api.Api("Checking the api library")
	file.ReadFile("Checking the file library. ReadFile")
	file.WriteFile("Checking the file library. WriteFile")
	storage.StorageTest()
}
