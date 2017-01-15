package shared

import (
	"fmt"
	"time"
)

type Newsletter struct {
	ID      int       `db:"id"`
	Date    time.Time `db:"date"`
	Name    string    `db:"name"`
	Content string    `db:"content"`
}

func (t Newsletter) ModelId() string {
	return fmt.Sprintf("%d", t.ID)
}

func (t Newsletter) RootURL() string {
	return "/api/news"
}
