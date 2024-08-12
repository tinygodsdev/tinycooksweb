package recipe

import "encoding/json"

const (
	indent = "  "
)

func (r *Recipe) TagsJSONString() string {
	data, err := json.MarshalIndent(r.Tags, "", indent)
	if err != nil {
		return "[]"
	}

	return string(data)
}

func TagsFromJSONString(data string) []*Tag {
	var tags []*Tag
	err := json.Unmarshal([]byte(data), &tags)
	if err != nil {
		return nil
	}

	for _, t := range tags {
		t.Slugify()
	}

	return tags
}

func (r *Recipe) IngredientsJSONString() string {
	data, err := json.MarshalIndent(r.Ingredients, "", indent)
	if err != nil {
		return "[]"
	}

	return string(data)
}

func IngredientsFromJSONString(data string) []*Ingredient {
	var ingredients []*Ingredient
	err := json.Unmarshal([]byte(data), &ingredients)
	if err != nil {
		return nil
	}

	for _, i := range ingredients {
		i.Slugify()
	}

	return ingredients
}

func (r *Recipe) EquipmentJSONString() string {
	data, err := json.MarshalIndent(r.Equipment, "", indent)
	if err != nil {
		return "[]"
	}

	return string(data)
}

func EquipmentFromJSONString(data string) []*Equipment {
	var equipment []*Equipment
	err := json.Unmarshal([]byte(data), &equipment)
	if err != nil {
		return nil
	}

	for _, e := range equipment {
		e.Slugify()
	}

	return equipment
}

func (r *Recipe) IdeasJSONString() string {
	data, err := json.MarshalIndent(r.Ideas, "", indent)
	if err != nil {
		return "[]"
	}

	return string(data)
}

func IdeasFromJSONString(data string) []*Idea {
	var ideas []*Idea
	err := json.Unmarshal([]byte(data), &ideas)
	if err != nil {
		return nil
	}

	return ideas
}

func (r *Recipe) SourcesJSONString() string {
	data, err := json.MarshalIndent(r.Sources, "", indent)
	if err != nil {
		return "[]"
	}

	return string(data)
}

func SourcesFromJSONString(data string) []*Source {
	var sources []*Source
	err := json.Unmarshal([]byte(data), &sources)
	if err != nil {
		return nil
	}

	return sources
}

func (r *Recipe) InstructionsJSONString() string {
	data, err := json.MarshalIndent(r.Instructions, "", indent)
	if err != nil {
		return "[]"
	}

	return string(data)
}

func InstructionsFromJSONString(data string) []Instruction {
	var instructions []Instruction
	err := json.Unmarshal([]byte(data), &instructions)
	if err != nil {
		return nil
	}

	return instructions
}

func (n *Nutrition) JSONString() string {
	data, err := json.MarshalIndent(n, "", indent)
	if err != nil {
		return "{}"
	}

	return string(data)
}

func NutritionFromJSONString(data string) *Nutrition {
	var nutrition Nutrition
	err := json.Unmarshal([]byte(data), &nutrition)
	if err != nil {
		return nil
	}

	return &nutrition
}

func (r *Recipe) JSONString() string {
	data, err := json.MarshalIndent(r, "", indent)
	if err != nil {
		return "{}"
	}

	return string(data)
}
