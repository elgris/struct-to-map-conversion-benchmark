package benchmark

import (
	"database/sql"
	"testing"
	"time"

	"github.com/elgris/stom"
	"github.com/go-sql-driver/mysql"
)

type ItemStomGeneral struct {
	ID              int             `db:"id" custom_tag:"id"`
	Name            string          `db:"name"`
	somePrivate     string          `db:"some_private" custom_tag:"some_private"`
	Number          int             `db:"number" custom_tag:"num"`
	Checksum        int32           `custom_tag:"sum"`
	Created         time.Time       `custom_tag:"created_time" db:"created"`
	Updated         mysql.NullTime  `db:"updated" custom_tag:"updated_time"`
	Price           float64         `db:"price"`
	Discount        *float64        `db:"discount"`
	IsReserved      sql.NullBool    `db:"reserved" custom_tag:"is_reserved"`
	Points          sql.NullInt64   `db:"points"`
	Rating          sql.NullFloat64 `db:"rating"`
	IsVisible       bool            `db:"visible" custom_tag:"visible"`
	SomeIgnoreField int             `db:"-" custom_tag:"i_ignore_nothing"`
	Notes           string
}

func TestStomGeneral(t *testing.T) {
	expected := GetExpectedResultStom()

	item := getItemStomGeneral()
	actual, _ := stom.ConvertToMap(item)

	AssertMapsEqual(t, expected, actual)
}

func TestStomGeneralPtr(t *testing.T) {
	expected := GetExpectedResultStom()

	item := getItemStomGeneral()
	actual, _ := stom.ConvertToMap(&item)

	AssertMapsEqual(t, expected, actual)
}

func BenchmarkStomGeneral(b *testing.B) {
	item := getItemStomGeneral()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		stom.ConvertToMap(item)
	}
}

func BenchmarkStomGeneralPtr(b *testing.B) {
	item := getItemStomGeneral()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		stom.ConvertToMap(&item)
	}
}

func getItemStomGeneral() ItemStomGeneral {
	discount := 111.0
	return ItemStomGeneral{
		ID:              1,
		Name:            "item_1",
		Number:          11,
		Checksum:        111,
		Created:         time.Unix(10000, 0),
		Updated:         mysql.NullTime{time.Unix(11000, 0), true},
		Price:           1111.0,
		Discount:        &discount,
		IsReserved:      sql.NullBool{true, true},
		Points:          sql.NullInt64{int64(11), false},
		Rating:          sql.NullFloat64{1.0, true},
		IsVisible:       true,
		Notes:           "foo",
		SomeIgnoreField: 10,
	}
}
