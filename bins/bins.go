package bins

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
