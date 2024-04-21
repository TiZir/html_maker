package db

type BookBody struct {
	ID     int
	Author string
	Title  string
	Genre  string
	Copies int
	Rating float64
}
