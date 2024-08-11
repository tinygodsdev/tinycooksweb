package recipe

import (
	"strings"

	"github.com/google/uuid"
)

const (
	TagSeparator = ":"
)

type Tag struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Group string    `json:"group"`
	Slug  string    `json:"slug"`
}

func (t *Tag) Title() string {
	return strings.Join([]string{t.Group, t.Name}, TagSeparator)
}

func TagFromString(tag string) *Tag {
	parts := strings.Split(tag, TagSeparator)
	if len(parts) != 2 {
		return nil
	}

	return &Tag{
		Group: parts[0],
		Name:  parts[1],
		Slug:  Slugify(tag),
	}
}
