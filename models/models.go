package models

import (
	"time"

	"github.com/lib/pq"
)

type Post struct {
	NEWSID      int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	// Likes       int    `json:"likes"`
	// DisLikes    int    `json:"dislikes"`
	// Comments []comment
	// private
	CreatedAt time.Time   `json:"createdat"`
	UpdatedAt pq.NullTime `json:"updatedat"`
}
