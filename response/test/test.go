package test

import (
	"time"

	"github.com/gin-gonic/gin"
)

type Post struct {
	ID        uint       `json:"id" db:"id" gorm:"primary_key;index"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" db:"deleted_at"` // `sql:"index"`

	Title string `json:"title" db:"title"`
}

func (p Post) Foo() string { return "foo" }
func (p Post) Zoo() string { return "zoo" }

var Action func(*gin.Context)
var Handler = func(c *gin.Context) { Action(c) }
