package presets

import (
	"context"
	"fmt"
	"net/http"

	"github.com/G-Core/gcorelabscdn-go/gcore"
)

var _ PresetsService = (*Service)(nil)

type Service struct {
	r gcore.Requester
}

func NewService(r gcore.Requester) *Service {
	return &Service{r: r}
}

func (s *Service) Apply(ctx context.Context, presetID int64, req *ApplyRequest) (*AppliedPreset, error) {
	preset := AppliedPreset{
		PresetID: int(presetID),
		ObjectID: req.ObjectID,
	}

	path := fmt.Sprintf("/cdn/presets/%d/applied", presetID)
	if err := s.r.Request(ctx, http.MethodPost, path, req, nil); err != nil {
		return nil, fmt.Errorf("request: %w", err)
	}

	return &preset, nil
}

func (s *Service) GetAppliedPreset(ctx context.Context, presetID, objectID int64) (*AppliedPreset, error) {
	var response GetAppliedPresetResponse
	preset := AppliedPreset{
		PresetID: int(presetID),
		ObjectID: int(objectID),
	}

	path := fmt.Sprintf("/cdn/presets/%d/applied", presetID)
	if err := s.r.Request(ctx, http.MethodGet, path, nil, &response); err != nil {
		return nil, fmt.Errorf("request: %w", err)
	}

	// check if objectID is in response.ObjectIDs list, if no - return error
	for _, id := range response.ObjectIDs {
		if id == int(objectID) {
			return &preset, nil
		}
	}

	return nil, fmt.Errorf("error: preset with id %d is not applied to the object with id %d", presetID, objectID)
}

func (s *Service) Unapply(ctx context.Context, presetID, objectID int64) error {
	path := fmt.Sprintf("/cdn/presets/%d/applied/%d", presetID, objectID)

	if err := s.r.Request(ctx, http.MethodDelete, path, nil, nil); err != nil {
		return fmt.Errorf("request: %w", err)
	}

	return nil
}
