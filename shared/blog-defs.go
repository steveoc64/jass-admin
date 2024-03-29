package shared

import (
	"fmt"
	"time"
)

type Blog struct {
	ID              int       `db:"id"`
	PostOrder       int       `db:"post_order"`
	Image           string    `db:"image"`
	Date            time.Time `db:"date"`
	Name            string    `db:"name"`
	Title           string    `db:"title"`
	Content         string    `db:"content"`
	ShareTwitter    int       `db:"share_twitter"`
	ShareFacebook   int       `db:"share_facebook"`
	ShareInstagram  int       `db:"share_instagram"`
	ShareGooglePlus int       `db:"share_google_plus"`
	Archived        bool      `db:"archived"`
	ImageURL        string
}

func (t *Blog) GetImageURL() string {
	return "/img/models/" + t.Image
}

func (t Blog) ModelId() string {
	return fmt.Sprintf("%d", t.ID)
}

func (t Blog) RootURL() string {
	return "/api/blog"
}
