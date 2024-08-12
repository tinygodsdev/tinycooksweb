package moderation

import (
	"context"

	"github.com/tinygodsdev/tinycooksweb/pkg/recipe"
)

const (
	// Moderation statuses
	// NB: for moderation statuses that are supposed to be queried by the store
	// a view with the same name should be created in the Airtable base
	ModerationStatusPending     = "pending"
	ModerationStatusApproved    = "approved"
	ModerationStatusRejected    = "rejected"
	ModerationStatusNeedsChange = "needsChange"
	ModerationStatusFinished    = "finished"
	ModerationStatusErrored     = "errored"
)

type ModerationStore interface {
	Get(ctx context.Context, moderationStatus string) ([]RecipeModerationInstance, error)
	GetApproved(ctx context.Context) ([]RecipeModerationInstance, error)
	Save(ctx context.Context, recipes []*recipe.Recipe) error
}

type RecipeModerationInstance interface {
	Recipe() *recipe.Recipe
	Approve(context.Context) error
	Reject(context.Context) error
	NeedsChange(context.Context) error
	Finish(context.Context) error
	Errored(context.Context, error) error
}
