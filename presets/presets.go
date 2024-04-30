package presets

import (
	"context"
)

type PresetsService interface {
	Create(ctx context.Context, presetID int64, req *CreateRequest) (*Preset, error)
	Get(ctx context.Context, presetID, objectID int64) (*Preset, error)
	Delete(ctx context.Context, presetID, objectID int64) error
}

type CreateRequest struct {
	ObjectID int `json:"object_id"`
}

type GetResponse struct {
	ObjectType *string `json:"object_type"`
	ObjectIDs  []int   `json:"object_ids"`
}

type Preset struct {
	ID       int
	ObjectID int
}
