package models

import (
	"github.com/ruspatrick/stan-svc/domain/models/news"
)

type News struct {
	Title string `json:"title,omitempty"`
	Date  string `json:"date,omitempty"`
}

func (n News) ToEntity() news.News {
	return news.News{
		Title: n.Title,
		Date:  n.Date,
	}
}

type NullableString struct {
	value    string
	hasValue bool
}

func NewNullableString(value string, hasValue bool) NullableString {
	return NullableString{value: value, hasValue: hasValue}
}

func (ns NullableString) HasValue() bool {
	return ns.hasValue
}

func (ns NullableString) GasValue() string {
	return ns.value
}
