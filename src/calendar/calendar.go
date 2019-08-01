package calendar 

import (
	"errors"
	"unicode/utf8"
)

type Event struct {
	title string
	Date
}

type Date struct {
	year int
	month int
	day int
}

func (d *Date) SetYear(year int) error{
	if year < 1 {
		return errors.New("invalid year")
	}
	d.year = year
	return nil
}

func (d *Date) SetMonth(month int) error{
	if month < 1 || month > 12{
		return errors.New("invalid month")
	}
	d.month = month
	return nil
}

func (d *Date) SetDay(day int) error{
	if day < 1 || day > 31 {
		return errors.New("invalid day")
	}
	d.day = day
	return nil
}

func (d *Date) Year() int {
	return d.year
}

func (d *Date) Month() int {
	return d.month
}

func (d *Date) Day() int {
	return d.day
}

func (e *Event) SetTitle(title string) error {
	if utf8.RuneCountInString(title) > 30{
		return errors.New("invalid title")
	}
	e.title = title
	return nil
}

func (e *Event) Titile() string{
	return e.title
}
