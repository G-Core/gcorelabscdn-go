package presets

import (
	"context"
)

type PresetsService interface {
	Apply(ctx context.Context, presetID int64, req *ApplyRequest) (*AppliedPreset, error)
	GetAppliedPreset(ctx context.Context, presetID, objectID int64) (*AppliedPreset, error)
	Unapply(ctx context.Context, presetID, objectID int64) error
}

type ApplyRequest struct {
	ObjectID int `json:"object_id"`
}

type GetAppliedPresetResponse struct {
	ObjectType *string `json:"object_type"`
	ObjectIDs  []int   `json:"object_ids"`
}

type AppliedPreset struct {
	PresetID int
	ObjectID int
}
