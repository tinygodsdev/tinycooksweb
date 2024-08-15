package recipe

import (
	"strings"

	"github.com/google/uuid"
	"github.com/samber/lo"
)

const (
	TagSeparator = ":"
)

type Tag struct {
	ID    uuid.UUID `json:"-" yaml:"-"`
	Name  string    `json:"name"`
	Group string    `json:"group"`
	Slug  string    `json:"-" yaml:"-"`
}

func (t *Tag) Title() string {
	return strings.Join([]string{t.Group, t.Name}, TagSeparator)
}

func TagFromTitle(tag string) *Tag {
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

func GetTagGroups(tags []*Tag) []string {
	return lo.Uniq(lo.Map(tags, func(t *Tag, _ int) string {
		return t.Group
	}))
}

func FilterTagsByGroup(tags []*Tag, group string) []*Tag {
	return lo.Filter(tags, func(t *Tag, _ int) bool {
		return t.Group == group
	})
}

func (t *Tag) Slugify() {
	t.Slug = Slugify(t.Title())
}
