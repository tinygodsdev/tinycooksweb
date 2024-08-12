package moderation

import (
	"context"
	"fmt"
	"time"

	"github.com/mehanizm/airtable"
	"github.com/tinygodsdev/datasdk/pkg/logger"
	"github.com/tinygodsdev/tinycooksweb/pkg/recipe"
)

const (
	defaultUploadBatchSize = 10

	// Airtable schema
	Name         = "Name"
	Description  = "Description"
	Text         = "Text"
	Lang         = "Lang"
	Servings     = "Servings"
	Time         = "Time"
	Tags         = "Tags"
	Ingredients  = "Ingredients"
	Equipment    = "Equipment"
	Rating       = "Rating"
	Ideas        = "Ideas"
	Sources      = "Sources"
	Instructions = "Instructions"
	Nutrition    = "Nutrition"
	Moderation   = "Moderation"
	Error        = "Error"
)

type Config struct {
	APIKey string
	BaseID string
	Table  string
}

type recipeWithModerationRecord struct {
	r *recipe.Recipe
	m *airtable.Record
}

type AirtableModerationStore struct {
	cfg    Config
	client *airtable.Client
	log    logger.Logger
	table  *airtable.Table
	schema *airtable.Tables
}

func NewAirtableModerationStore(cfg Config, log logger.Logger) (*AirtableModerationStore, error) {
	client := airtable.NewClient(cfg.APIKey)

	schema, err := client.GetBaseSchema(cfg.BaseID).Do()
	if err != nil {
		return nil, err
	}

	table := client.GetTable(cfg.BaseID, cfg.Table)
	if table == nil {
		return nil, fmt.Errorf("table not found: %s", cfg.Table)
	}

	return &AirtableModerationStore{
		cfg:    cfg,
		log:    log,
		client: client,
		table:  table,
		schema: schema,
	}, nil
}

func (s *AirtableModerationStore) Get(
	ctx context.Context,
	moderationStatus string,
) ([]RecipeModerationInstance, error) {
	records, err := s.table.GetRecords().
		FromView(moderationStatus).
		Do()
	if err != nil {
		return nil, err
	}

	var res []RecipeModerationInstance
	for _, record := range records.Records {
		recipe, err := s.parseRecord(record)
		if err != nil {
			return nil, err
		}

		res = append(res, &recipeWithModerationRecord{
			r: recipe,
			m: record,
		})
	}

	return res, nil
}

func (s *AirtableModerationStore) Save(ctx context.Context, recipes []*recipe.Recipe) error {
	for i := 0; i < len(recipes); i += defaultUploadBatchSize {
		end := i + defaultUploadBatchSize
		if end > len(recipes) {
			end = len(recipes)
		}

		records := &airtable.Records{Typecast: true}
		for _, r := range recipes[i:end] {
			record := &airtable.Record{
				Fields: map[string]any{
					Name:         r.Name,
					Description:  r.Description,
					Text:         r.Text,
					Lang:         r.Lang,
					Servings:     r.Servings,
					Tags:         r.TagsJSONString(),
					Ingredients:  r.IngredientsJSONString(),
					Equipment:    r.EquipmentJSONString(),
					Ideas:        r.IdeasJSONString(),
					Sources:      r.SourcesJSONString(),
					Instructions: r.InstructionsJSONString(),
					Rating:       r.Rating,
					Time:         r.Time.Minutes(),
					Nutrition:    r.Nutrition.JSONString(),
					Moderation:   ModerationStatusPending,
				},
			}
			records.Records = append(records.Records, record)
		}

		_, err := s.table.AddRecords(records)
		if err != nil {
			s.log.Error("failed to add records", "error", err)
			continue
		}
		time.Sleep(50 * time.Millisecond)
	}

	s.log.Info("saving to Airtable done")
	return nil
}

func (s *AirtableModerationStore) GetApproved(ctx context.Context) ([]RecipeModerationInstance, error) {
	return s.Get(ctx, ModerationStatusApproved)
}

func (r *recipeWithModerationRecord) Recipe() *recipe.Recipe {
	return r.r
}

func (r *recipeWithModerationRecord) updateModerationStatus(
	ctx context.Context,
	status string,
	err error,
) error {
	errStr := ""
	if err != nil {
		errStr = err.Error()
	}
	_, err = r.m.UpdateRecordPartial(map[string]any{Moderation: status, Error: errStr})
	if err != nil {
		return err
	}

	return nil
}

func (r *recipeWithModerationRecord) Approve(ctx context.Context) error {
	return r.updateModerationStatus(ctx, ModerationStatusApproved, nil)
}

func (r *recipeWithModerationRecord) Reject(ctx context.Context) error {
	return r.updateModerationStatus(ctx, ModerationStatusRejected, nil)
}

func (r *recipeWithModerationRecord) NeedsChange(ctx context.Context) error {
	return r.updateModerationStatus(ctx, ModerationStatusNeedsChange, nil)
}

func (r *recipeWithModerationRecord) Finish(ctx context.Context) error {
	return r.updateModerationStatus(ctx, ModerationStatusFinished, nil)
}

func (r *recipeWithModerationRecord) Errored(ctx context.Context, err error) error {
	return r.updateModerationStatus(ctx, ModerationStatusErrored, err)
}

func (s *AirtableModerationStore) parseRecord(record *airtable.Record) (*recipe.Recipe, error) {
	r := &recipe.Recipe{
		Name:             getFieldString(record, Name),
		Description:      getFieldString(record, Description),
		Text:             getFieldString(record, Text),
		Lang:             getFieldString(record, Lang),
		Servings:         getFieldInt(record, Servings),
		Rating:           getFieldFloat32(record, Rating),
		Tags:             recipe.TagsFromJSONString(getFieldString(record, Tags)),
		Ingredients:      recipe.IngredientsFromJSONString(getFieldString(record, Ingredients)),
		Equipment:        recipe.EquipmentFromJSONString(getFieldString(record, Equipment)),
		Ideas:            recipe.IdeasFromJSONString(getFieldString(record, Ideas)),
		Sources:          recipe.SourcesFromJSONString(getFieldString(record, Sources)),
		Instructions:     recipe.InstructionsFromJSONString(getFieldString(record, Instructions)),
		Nutrition:        recipe.NutritionFromJSONString(getFieldString(record, Nutrition)),
		ModerationStatus: getFieldString(record, Moderation),
		Time:             getFieldMinutes(record, Time),
	}

	r.SlugifyAll()
	return r, nil
}

func getFieldString(record *airtable.Record, field string) string {
	if record.Fields[field] == nil {
		return ""
	}
	return record.Fields[field].(string)
}

func getFieldInt(record *airtable.Record, field string) int {
	// airtable returns float64 for integers
	if record.Fields[field] == nil {
		return 0
	}

	return int(record.Fields[field].(float64))
}

func getFieldFloat32(record *airtable.Record, field string) float32 {
	if record.Fields[field] == nil {
		return 0
	}

	return float32(record.Fields[field].(float64))
}

func getFieldMinutes(record *airtable.Record, field string) time.Duration {
	if record.Fields[field] == nil {
		return 0
	}

	return time.Duration(int(record.Fields[field].(float64))) * time.Minute
}
