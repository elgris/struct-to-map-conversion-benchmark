package benchmark

import (
	"testing"

	"github.com/fatih/structs"
)

func init() {
	structs.DefaultTagName = "db"
}

type ItemStructs struct {
	ID          int     `db:"id" custom_tag:"id"`
	Name        string  `db:"name"`
	somePrivate string  `db:"some_private" custom_tag:"some_private"`
	Number      int     `db:"number" custom_tag:"num"`
	Created     string  `custom_tag:"created_time" db:"created"`
	Updated     string  `db:"updated" custom_tag:"updated_time"`
	Price       float64 `db:"price"`
	Discount    float64 `db:"discount"`
	IsReserved  bool    `db:"reserved" custom_tag:"is_reserved"`
	Points      int64   `db:"points"`
	Rating      float64 `db:"rating"`
	IsVisible   bool    `db:"visible" custom_tag:"visible"`
}

func TestStructs(t *testing.T) {
	expected := getExpectedResultStructs()

	item := getItemStructs()
	actual := structs.Map(item)

	AssertMapsEqual(t, expected, actual)
}

func TestStructsPtr(t *testing.T) {
	expected := getExpectedResultStructs()

	item := getItemStructs()
	actual := structs.Map(&item)

	AssertMapsEqual(t, expected, actual)
}

func BenchmarkStructs(b *testing.B) {
	item := getItemStructs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		structs.Map(item)
	}
}

func BenchmarkStructsPtr(b *testing.B) {
	item := getItemStructs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		structs.Map(&item)
	}
}

func getItemStructs() ItemStructs {
	return ItemStructs{
		ID:         1,
		Name:       "item_1",
		Number:     11,
		Created:    "1975-12-12 12:13:14",
		Updated:    "1975-12-12 12:13:14",
		Price:      1111.0,
		Discount:   111.0,
		IsReserved: true,
		Points:     11,
		Rating:     1.0,
		IsVisible:  true,
	}
}

func getExpectedResultStructs() map[string]interface{} {
	return map[string]interface{}{
		"id":       1,
		"name":     "item_1",
		"number":   11,
		"created":  "1975-12-12 12:13:14",
		"updated":  "1975-12-12 12:13:14",
		"discount": 111.0,
		"price":    1111.0,
		"reserved": true,
		"points":   int64(11),
		"rating":   1.0,
		"visible":  true,
	}
}
