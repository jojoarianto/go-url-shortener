package model

import "github.com/jinzhu/gorm"

type Shorten struct {
	gorm.Model
	Url       string `json:url`
	Shortcode string `json:shortcode`
}
