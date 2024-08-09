package handler

import (
	"fmt"
	"html/template"
	"time"

	"github.com/bradfitz/iter"
	"github.com/tinygodsdev/tinycooksweb/pkg/locale"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var funcMap = template.FuncMap{
	"N":     iter.N,
	"Plus1": func(i int) int { return i + 1 },
	"Mean": func(data ...float64) float64 {
		if len(data) == 0 {
			return 0
		}
		var sum float64
		for _, n := range data {
			sum += n
		}
		return sum / float64(len(data))
	},
	"Perc": func(min, max, v float64) float64 {
		if max == min {
			return 0
		}
		return (v - min) / (max - min)
	},
	"DerefInt": func(i *int) int {
		if i == nil {
			return 0
		}
		return *i
	},
	"DisplayTime": func(t time.Time) string {
		return t.Format(DefaultDisplayTime)
	},
	"DisplayTechTime": func(t time.Time) string {
		return t.Format("2006-01-02 15:04:05.000 MST")
	},
	"Since": func(t time.Time) time.Duration {
		return time.Since(t)
	},
	"UILocales": func() []string {
		return locale.List()
	},
	"LocaleParam": func(loc string) string {
		if loc == locale.Default() {
			return "/"
		}
		return fmt.Sprintf("?locale=%s", loc)
	},
	"FormatDuration": func(d time.Duration) string {
		z := time.Unix(0, 0).UTC()
		return z.Add(d).Format("15:04")
	},
	"Title": func(t string) string {
		return cases.Title(language.Make(locale.Default())).String(t)
	},
	"RandomEmoji": func() string {
		return "🍏"
	},
	"max": func(a, b int) int {
		if a > b {
			return a
		}
		return b
	},
	"min": func(a, b int) int {
		if a < b {
			return a
		}
		return b
	},
	"seq": func(start, end int) []int {
		s := make([]int, end-start+1)
		for i := range s {
			s[i] = start + i
		}
		return s
	},
	"sub": func(i1, i2 int) int {
		return i1 - i2
	},
	"add": func(i1, i2 int) int {
		return i1 + i2
	},
}
