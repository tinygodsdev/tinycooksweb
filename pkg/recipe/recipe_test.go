package recipe

import (
	"testing"
	"time"

	"github.com/tinygodsdev/tinycooksweb/pkg/locale"
	"gotest.tools/v3/assert"
)

func TestRecipe_addTimeTag(t *testing.T) {
	tests := []struct {
		name          string
		lang          string
		time          time.Duration
		expectedTag   string
		expectedGroup string
	}{
		{
			name:          "Fast in Russian",
			lang:          locale.Ru,
			time:          25 * time.Minute,
			expectedTag:   "Быстро",
			expectedGroup: "Время",
		},
		{
			name:          "Medium in Russian",
			lang:          locale.Ru,
			time:          45 * time.Minute,
			expectedTag:   "Средне",
			expectedGroup: "Время",
		},
		{
			name:          "Long in Russian",
			lang:          locale.Ru,
			time:          75 * time.Minute,
			expectedTag:   "Долго",
			expectedGroup: "Время",
		},
		{
			name:          "Very long in Russian",
			lang:          locale.Ru,
			time:          2 * time.Hour,
			expectedTag:   "Очень долго",
			expectedGroup: "Время",
		},
		{
			name:          "Fast in English",
			lang:          locale.En,
			time:          25 * time.Minute,
			expectedTag:   "Fast",
			expectedGroup: "Time",
		},
		{
			name:          "Medium in English",
			lang:          locale.En,
			time:          45 * time.Minute,
			expectedTag:   "Medium",
			expectedGroup: "Time",
		},
		{
			name:          "Long in English",
			lang:          locale.En,
			time:          75 * time.Minute,
			expectedTag:   "Long",
			expectedGroup: "Time",
		},
		{
			name:          "Very long in English",
			lang:          locale.En,
			time:          2 * time.Hour,
			expectedTag:   "Very long",
			expectedGroup: "Time",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			recipe := &Recipe{
				Lang: tt.lang,
				Time: tt.time,
			}

			recipe.PostProcess()

			assert.Equal(t, 1, len(recipe.Tags))
			assert.Equal(t, tt.expectedTag, recipe.Tags[0].Name)
			assert.Equal(t, tt.expectedGroup, recipe.Tags[0].Group)
		})
	}
}
