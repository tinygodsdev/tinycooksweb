package recipe

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestLoadRecipeFromString(t *testing.T) {
	expectedRecipe := &Recipe{
		Name:        "Кофейные булочки с корицей",
		Description: "Восхитительные кофейные булочки с корицей для идеального завтрака или сладкого перекуса. Они сочетают нежное тесто, насыщенную кофейно-коричную начинку и простой кофейный глазурь.",
		Text:        "Эти кофейные булочки с корицей поразят любого любителя кофе и сладкой выпечки. Тесто, наполненное ароматом эспрессо, делает их непревзойденными. Использование масла вместо муки повышает мягкость булочек, а замешивание дрожжевого теста дает более рыхлую структуру.",
		Tags: []*Tag{
			{Name: "Интернациональная", Group: "Кухня"},
			{Name: "Хлеб и выпечка", Group: "Тип блюда"},
			{Name: "Завтрак", Group: "Тема"},
			{Name: "Вегетарианская", Group: "Диета"},
			{Name: "Запекание", Group: "Метод готовки"},
			{Name: "Сладкий", Group: "Вкус"},
			{Name: "Осень", Group: "Сезон"},
			{Name: "Мягкое", Group: "Текстура"},
			{Name: "Высококалорийные", Group: "Питательность"},
		},
		Ingredients: []*Ingredient{
			{Product: &Product{Name: "молоко"}, Quantity: "180", Unit: "мл"},
			{Product: &Product{Name: "кофе растворимый"}, Quantity: "14", Unit: "г"},
			{Product: &Product{Name: "сахар"}, Quantity: "80", Unit: "г"},
			{Product: &Product{Name: "дрожжи быстродействующие"}, Quantity: "9", Unit: "г"},
			{Product: &Product{Name: "яйцо"}, Quantity: "2", Unit: "шт."},
			{Product: &Product{Name: "сливочное масло"}, Quantity: "75", Unit: "г"},
			{Product: &Product{Name: "пшеничная мука"}, Quantity: "480", Unit: "г"},
			{Product: &Product{Name: "соль"}, Quantity: "1/2", Unit: "ч. л."},
			{Product: &Product{Name: "сахар коричневый"}, Quantity: "100", Unit: "г"},
			{Product: &Product{Name: "корица"}, Quantity: "1", Unit: "ст. л."},
			{Product: &Product{Name: "сахарная пудра"}, Quantity: "150", Unit: "г"},
			{Product: &Product{Name: "эспрессо"}, Quantity: "30", Unit: "мл"},
		},
		Equipment: []*Equipment{
			{Name: "духовка"},
			{Name: "форма для выпекания"},
			{Name: "миксер"},
			{Name: "скалка"},
			{Name: "венчик"},
			{Name: "силиконовая лопатка"},
		},
		Instructions: []Instruction{
			{Text: "Нагрейте молоко до теплого состояния, добавьте растворимый кофе и размешайте до полного растворения. Оставьте остывать до слегка теплого состояния."},
			{Text: "Добавьте сахар и дрожжи, оставьте на 10 минут до образования пузырьков."},
			{Text: "Вмешайте яйца и растопленное сливочное масло в дрожжевую смесь."},
			{Text: "Добавьте муку и соль. Быстро перемешайте тесто лопаткой."},
			{Text: "Оставьте тесто при комнатной температуре на 15 минут."},
			{Text: "Установите крюк для теста на миксер и замешивайте тесто на средней скорости 10 минут, затем еще 2-4 минуты на высокой скорости до гладкого состояния."},
			{Text: "Переложите тесто в смазанную емкость и оставьте в холодильнике на 8-10 часов до увеличения объема."},
			{Text: "Для начинки смешайте мягкое масло, коричневый сахар и корицу до получения пасты."},
			{Text: "Раскатайте тесто в прямоугольник 30x40 см на слегка присыпанной мукой поверхности."},
			{Text: "Размажьте начинку по тесту, посыпьте оставшимся растворимым кофе."},
			{Text: "Сверните тесто в рулет по длинной стороне и нарежьте его на 12 кусочков."},
			{Text: "Переложите булочки в форму для выпекания и оставьте на 30-60 минут при комнатной температуре до увеличения вдвое."},
			{Text: "Разогрейте духовку до 180°C и выпекайте 23-25 минут до золотистого цвета."},
			{Text: "Для глазури смешайте сахарную пудру с эспрессо до получения гладкой массы. Полейте глазурью остуженные булочки."},
		},
		Time:     11*time.Hour + 10*time.Minute,
		Servings: 12,
		Nutrition: &Nutrition{
			Calories:  397,
			Fat:       15,
			Carbs:     60,
			Protein:   6,
			Precision: "approx",
		},
		Meta: map[string]string{
			"createdBy": "ИИ-Помощник",
		},
	}

	expectedRecipe.PostProcess()

	// Define test cases
	testCases := []struct {
		name     string
		filePath string
		expected *Recipe
	}{
		{
			name:     "Test with quotes",
			filePath: "test_data/test01.txt",
			expected: expectedRecipe,
		},
		{
			name:     "Test without quotes",
			filePath: "test_data/test02.yaml",
			expected: expectedRecipe,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			data, err := os.ReadFile(tc.filePath)
			assert.NoError(t, err, "Failed to read file %s", tc.filePath)

			recipe, err := LoadRecipeFromString(string(data))
			assert.NoError(t, err, "Error loading recipe from %s", tc.filePath)
			assert.NotNil(t, recipe, "Recipe should not be nil")
			assert.Equal(t, tc.expected, recipe, "Loaded recipe does not match the expected recipe")
		})
	}
}
