package handle

import "github.com/wongkaew4/godb_book/model"

type BookData struct {
	data []model.Book
}

func NewBook() *BookData {
	return &BookData{}
}

func (b *BookData) AllData() []model.Book {
	return b.data
}

func (b *BookData) AddData(d model.Book) model.Book {
	b.data = append(b.data, d)
	return d
}
