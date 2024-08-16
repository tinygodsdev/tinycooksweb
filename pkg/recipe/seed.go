package recipe

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"gopkg.in/yaml.v3"
)

const (
	defaultDataPath  = "data/seed"
	dataIgnorePrefix = '_' // ignore files starting with this prefix
)

func SeedData() []*Recipe {
	return []*Recipe{
		PumpkinBuns(),
		PumpkinSoup(),
	}
}

// SeedToYAML saves each seed recipe to a file (using Recipe.Slug as file name)
func SeedToYAML(path string) error {
	if path == "" {
		path = defaultDataPath
	}

	recipes := SeedData()
	for _, recipe := range recipes {
		filePath := filepath.Join(path, recipe.Slug+".yaml")

		file, err := os.Create(filePath)
		if err != nil {
			return fmt.Errorf("error creating file %s: %w", filePath, err)
		}
		defer file.Close()

		encoder := yaml.NewEncoder(file)
		defer encoder.Close()

		if err := encoder.Encode(recipe); err != nil {
			return fmt.Errorf("error encoding recipe %s to YAML: %w", recipe.Slug, err)
		}
	}

	return nil
}

func LoadFromYAML(path string) ([]*Recipe, error) {
	if path == "" {
		path = defaultDataPath
	}

	var recipes []*Recipe

	files, err := os.ReadDir(path)
	if err != nil {
		return nil, fmt.Errorf("error reading directory %s: %w", path, err)
	}

	for _, file := range files {
		if file.Name()[0] == dataIgnorePrefix || filepath.Ext(file.Name()) != ".yaml" {
			continue
		}

		filePath := filepath.Join(path, file.Name())

		content, err := os.ReadFile(filePath)
		if err != nil {
			return nil, fmt.Errorf("error reading file %s: %w", filePath, err)
		}

		var recipe Recipe
		err = yaml.Unmarshal(content, &recipe)
		if err != nil {
			return nil, fmt.Errorf("error unmarshaling YAML from file %s: %w", filePath, err)
		}

		recipes = append(recipes, &recipe)
	}

	for _, recipe := range recipes {
		recipe.PostProcess()
	}

	return recipes, nil
}

func PumpkinBuns() *Recipe {
	recipe := &Recipe{
		Name:        "Тыквенные булочки",
		Lang:        LangRu,
		Rating:      4.7,
		Description: "Шикарные мягкие тыквенные булочки, которые можно подавать вместо хлеба, использовать для бургеров и сэндвичей.",
		Text:        `Новый тыквенный рецепт, пока сезон тыквы идет. Эти булочки готовятся безопарным методом, что сокращает время приготовления. Тесто готовится без яиц. Булочки получаются мягкими и имеют красивый цвет благодаря тыквенному пюре.`,
		Ingredients: []*Ingredient{
			{Product: &Product{Name: "молоко"}, Quantity: "150", Unit: "грамм"},
			{Product: &Product{Name: "сухие быстродействующие дрожжи"}, Quantity: "5", Unit: "грамм"},
			{Product: &Product{Name: "сахар"}, Quantity: "40", Unit: "грамм"},
			{Product: &Product{Name: "соль"}, Quantity: "1/2", Unit: "ч.л."},
			{Product: &Product{Name: "сливочное масло"}, Quantity: "50", Unit: "грамм"},
			{Product: &Product{Name: "тыквенное пюре"}, Quantity: "150", Unit: "грамм"},
			{Product: &Product{Name: "пшеничная мука"}, Quantity: "400", Unit: "грамм"},
			{Product: &Product{Name: "желтки"}, Quantity: "1", Unit: "шт."},
			{Product: &Product{Name: "семена чиа, лен, кунжут или мак (по желанию)"}, Quantity: "", Unit: "", Optional: true},
		},
		Instructions: []Instruction{
			{Text: "Слегка подогрейте молоко до температуры не выше 45 градусов, добавьте дрожжи и дайте постоять 3-4 минуты. Перемешайте до растворения дрожжей."},
			{Text: "Добавьте соль, сахар, тыквенное пюре и растопленное сливочное масло (должно быть теплым)."},
			{Text: "Добавьте 80% от указанного количества муки и начните замешивать тесто. Постепенно добавляйте оставшуюся муку, пока тесто не станет мягким и гладким."},
			{Text: "Накройте тесто и дайте ему подойти вдвое (1-2 часа)."},
			{Text: "Разделите подошедшее тесто на 8 частей и сформируйте круглые заготовки. Выложите их на противень, накройте пленкой и оставьте для финальной расстойки на 30 минут."},
			{Text: "Смажьте заготовки смесью желтка с молоком (2 ст. ложки) и по желанию посыпьте семенами. Выпекайте в предварительно разогретой до 180 градусов духовке 20-25 минут."},
		},
		Equipment: []*Equipment{
			{Name: "печь"},
			{Name: "противень"},
			{Name: "миксер"},
		},
		Tags: []*Tag{
			{Name: "выпечка", Group: "тип блюда"},
			{Name: "вегетарианское", Group: "диета"},
			{Name: "легкое", Group: "сложность"},
			{Name: "сладкое", Group: "вкус"},
		},
		Ideas: []*Idea{
			{Text: "Добавьте в тыквенное пюре чеснок, розмарин или тимьян для пряного вкуса."},
			{Text: "Используйте курагу вместе сахаром для более сладких булочек."},
		},
		Time:     160 * time.Minute,
		Servings: 8,
		Sources: []*Source{
			{
				Name:        "vkusnyblog.com",
				Description: "Вкусный блог",
				URL:         "https://www.vkusnyblog.com/recipe/tykvennye-bulochki/",
			},
		},
		Nutrition: &Nutrition{
			Calories:  200,
			Fat:       8,
			Carbs:     30,
			Protein:   5,
			Precision: NutritionPrecisionApprox,
			Benefits:  []string{"Витамины", "Минералы"},
		},
		Meta: map[string]string{
			"type":      "seed",
			"createdBy": "Dani Polani",
		},
	}

	recipe.PostProcess()
	return recipe
}

func PumpkinSoup() *Recipe {
	recipe := &Recipe{
		Name:        "Тыквенный суп-пюре",
		Lang:        LangRu,
		Rating:      4.5,
		Description: "Для супа лучше выбирать тыкву не сладких сортов. При приготовлении тыквы, старайтесь не переварить, иначе она потеряет вкус.",
		Text:        `Для супа лучше выбирать тыкву не сладких сортов. При приготовлении тыквы, старайтесь не переварить, иначе она потеряет вкус.`,
		Ingredients: []*Ingredient{
			{Product: &Product{Name: "тыква"}, Quantity: "400", Unit: "г"},
			{Product: &Product{Name: "лук репчатый"}, Quantity: "1", Unit: "шт."},
			{Product: &Product{Name: "молоко"}, Quantity: "200", Unit: "мл"},
			{Product: &Product{Name: "масло сливочное"}, Quantity: "20", Unit: "г"},
			{Product: &Product{Name: "петрушка"}, Quantity: "по вкусу", Unit: "", Optional: true},
			{Product: &Product{Name: "приправа"}, Quantity: "по вкусу", Unit: "", Optional: true},
			{Product: &Product{Name: "сливки 35-38%"}, Quantity: "по вкусу", Unit: ""},
			{Product: &Product{Name: "хлеб белый"}, Quantity: "по вкусу", Unit: ""},
		},
		Instructions: []Instruction{
			{Text: "Лук измельчить, тыкву нарезать кубиками."},
			{Text: "На сливочном масле слегка обжарить лук. Добавить тыкву, посолить, приправить и потушить с добавлением небольшого количества воды. Примерно 20 минут."},
			{Text: "Готовую тыкву переложить в блендер и измельчить. Добавить горячее молоко. И еще раз пробить блендером."},
			{Text: "Перелить в кастрюльку и прогреть, но не доводить до кипения."},
			{Text: "Разлить по тарелкам, добавить сухарики, приправить сметаной или жирными сливками и украсить зеленью петрушки."},
		},
		Equipment: []*Equipment{
			{Name: "кастрюля"},
			{Name: "блендер"},
		},
		Tags: []*Tag{
			{Name: "супы", Group: "тип блюда"},
			{Name: "вегетарианское", Group: "диета"},
			{Name: "легкое", Group: "сложность"},
			{Name: "соленое", Group: "вкус"},
		},
		Ideas: []*Idea{
			{Text: "Для пряного вкуса добавьте чеснок, розмарин или тимьян."},
		},
		Time:     35 * time.Minute,
		Servings: 1,
		Sources: []*Source{
			{
				Name:        "edimdoma.ru",
				Description: "ЕдимДома",
				URL:         "https://www.edimdoma.ru/retsepty/59418-tykvennyy-sup-pyure",
			},
		},
		Nutrition: &Nutrition{
			Calories:  427,
			Fat:       25,
			Carbs:     49,
			Protein:   13,
			Precision: NutritionPrecisionApprox,
			Benefits:  []string{"Белки", "Жиры", "Углеводы"},
		},
		Meta: map[string]string{
			"type":      "seed",
			"createdBy": "Dani Polani",
		},
	}

	recipe.PostProcess()
	return recipe
}
