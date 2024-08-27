package handler

import (
	"fmt"
	"html/template"
	"path/filepath"
	"time"
	"unicode"

	"github.com/bradfitz/iter"
	"github.com/tinygodsdev/tinycooksweb/pkg/locale"
	"golang.org/x/exp/rand"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// this function might panic if no names are provided
// it is intended to be used only on server startup
// func (h *Handler) template(names ...string) *template.Template {
// 	if len(names) == 0 {
// 		panic("template: no names provided")
// 	}

// 	base := names[0]

// 	return template.Must(template.New(base).Funcs(funcMap()).ParseFiles(
// 		lo.Map(names, func(n string, _ int) string { return h.t + n })...,
// 	))
// }

// this function might panic if no names are provided
// it is intended to be used only on server startup
func (h *Handler) template(base string, dir string) *template.Template {
	tmpl := template.Must(template.New(base).Funcs(funcMap()).ParseFiles(h.t + base))
	pattern := filepath.Join(h.t, dir, "*.html")
	files, err := filepath.Glob(pattern)
	if err != nil {
		panic("template: error while finding files in directory " + dir)
	}

	if len(files) == 0 {
		panic("template: no files found in directory " + dir)
	}

	return template.Must(tmpl.ParseFiles(files...))
}

func funcMap() template.FuncMap {
	return template.FuncMap{
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
		"Capitalize": func(s string) string {
			if s == "" {
				return ""
			}

			r := []rune(s)
			r[0] = unicode.ToUpper(r[0])
			return string(r)
		},
		"RandomEmoji": func() string {
			foodEmojis := []string{
				"🍏", "🍎", "🍐", "🍊", "🍋", "🍌", "🍉", "🍇", "🍓", "🫐", "🍈", "🍒", "🍑", "🥭", "🍍", "🥥", "🥝",
				"🍅", "🍆", "🥑", "🥦", "🥬", "🥒", "🌶", "🫑", "🌽", "🥕", "🫒", "🧄", "🧅", "🥔", "🍠", "🥐", "🥯",
				"🍞", "🥖", "🥨", "🧀", "🥚", "🍳", "🧈", "🥞", "🧇", "🥓", "🥩", "🍗", "🍖", "🌭", "🍔", "🍟", "🍕",
				"🫓", "🥪", "🥙", "🧆", "🌮", "🌯", "🫔", "🥗", "🥘", "🫕", "🍝", "🍜", "🍲", "🍛", "🍣", "🍱", "🥟",
				"🍤", "🍙", "🍚", "🍘", "🍥", "🥠", "🥮", "🍢", "🍡", "🍧", "🍨", "🍦", "🥧", "🧁", "🍰", "🎂", "🍮",
				"🍭", "🍬", "🍫", "🍿", "🧋", "🍩", "🍪", "🌰", "🥜", "🍯", "🥛", "🍼", "☕", "🍵", "🧃", "🥤", "🍶",
				"🍺", "🍻", "🥂", "🍷", "🥃", "🍸", "🍹", "🧉", "🍾", "🧊",
			}

			rand.Seed(uint64(time.Now().UnixNano()))
			return foodEmojis[rand.Intn(len(foodEmojis))]
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
}
