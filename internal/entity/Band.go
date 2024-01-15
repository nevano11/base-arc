package entity

import "fmt"

type Band struct {
	Id       int8
	BandName string
	Year     uint16
}

func (b Band) String() string {
	return fmt.Sprintf("Band: {id=%d, name=%s, year=%d}", b.Id, b.BandName, b.Year)
}
