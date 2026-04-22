package main

import (
	"fmt"
	"time"
)

type Bin struct {
	ID        string
	Private   bool
	CreatedAt time.Time
	Name      string
}

// Структура BinList — список или коллекция Bin
type BinList struct {
	Bins []Bin
}

// Функция для создания нового Bin
func NewBin(id string, private bool, createdAtStr string, name string) Bin {
	// парсим строку даты
	t, err := time.Parse("2006-01-02T15:04:05-07:00", createdAtStr)
	if err != nil {
		fmt.Println("Ошибка:", err)
	}

	return Bin{
		ID:        id,
		Private:   private,
		CreatedAt: t,
		Name:      name,
	}
}

// Функция для создания нового BinList
func NewBinList() *BinList {
	return &BinList{
		Bins: []Bin{},
	}
}

// Метод для добавления Bin в BinList
func (bl *BinList) AddBin(bin Bin) {
	bl.Bins = append(bl.Bins, bin)
}

func main() {
	// binList := []byte(`{"id": "12345abcde", "private": true, "createdAt": "2024-04-27T15:34:56+03:00", "name": "Пример объекта"}`)

	// Создаем новый Bin
	bin1 := NewBin("12345abcde", true, "2024-04-27T15:34:56+03:00", "Пример объекта 1")
	bin2 := NewBin("67890fghij", false, "2024-04-27T15:34:56+03:00", "Пример объекта 2")

	// Создаем список Bin
	binList := NewBinList()

	// Добавляем Bin в список
	binList.AddBin(bin1)
	binList.AddBin(bin2)

	fmt.Println(binList)
}
