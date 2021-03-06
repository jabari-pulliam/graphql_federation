// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type PageInfo struct {
	TotalCount int `json:"totalCount"`
}

type Widget struct {
	ID    int         `json:"id"`
	Size  int         `json:"size"`
	Color WidgetColor `json:"color"`
}

func (Widget) IsEntity() {}

type WidgetFilter struct {
	Colors  []WidgetColor `json:"colors"`
	MinSize *int          `json:"minSize"`
	MaxSize *int          `json:"maxSize"`
}

type WidgetPage struct {
	Items    []*Widget `json:"items"`
	PageInfo *PageInfo `json:"pageInfo"`
}

type WidgetSource struct {
	Ids     []int       `json:"ids"`
	Widgets *WidgetPage `json:"widgets"`
}

func (WidgetSource) IsEntity() {}

type WidgetColor string

const (
	WidgetColorRed    WidgetColor = "RED"
	WidgetColorGreen  WidgetColor = "GREEN"
	WidgetColorBlue   WidgetColor = "BLUE"
	WidgetColorOrange WidgetColor = "ORANGE"
	WidgetColorBlack  WidgetColor = "BLACK"
)

var AllWidgetColor = []WidgetColor{
	WidgetColorRed,
	WidgetColorGreen,
	WidgetColorBlue,
	WidgetColorOrange,
	WidgetColorBlack,
}

func (e WidgetColor) IsValid() bool {
	switch e {
	case WidgetColorRed, WidgetColorGreen, WidgetColorBlue, WidgetColorOrange, WidgetColorBlack:
		return true
	}
	return false
}

func (e WidgetColor) String() string {
	return string(e)
}

func (e *WidgetColor) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = WidgetColor(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid WidgetColor", str)
	}
	return nil
}

func (e WidgetColor) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
