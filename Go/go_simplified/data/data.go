package data

import (
	"fmt"
	"sync"
)

type Book struct {
	ID       int
	Title    string
	Finished bool
}

var books = []*Book{
	{1, "Dune", false},
	{2, "El perfume", false},
	{3, "El principito", false},
	{4, "Cobro de sangre", false},
	{5, "La divina comedia", false},
	{6, "El alquimista", false},
	{7, "Pedro picapiedras", false},
	{8, "El hoyo", false},
	{9, "Baldor", false},
}

func FindBook(id int, m *sync.RWMutex) (int, *Book) {
	index := -1
	var book *Book
	m.RLock()
	for _, b := range books {
		if b.ID == id {
			index = 1
			book = b
		}
	}
	m.RUnlock()
	return index, book
}

func FinishBook(id int, m *sync.RWMutex) {
	i, book := FindBook(id, m)
	if i < 0 {
		return
	}
	m.Lock()
	book.Finished = true
	books[i] = book
	m.Unlock()
	fmt.Printf("Finished book: %s\n", book.Title)
}
