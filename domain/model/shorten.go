package model

import "github.com/jinzhu/gorm"

// Shorten struct of data
type Shorten struct {
	gorm.Model
	Url       string `validate:"required" json:"url"`
	Shortcode string `gorm:"index:idx_shortcode" json:"shortcode"`
}
