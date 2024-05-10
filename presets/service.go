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

func (s *Service) Apply(ctx context.Context, presetID int, req *ApplyRequest) (*AppliedPreset, error) {
	preset := AppliedPreset{
		PresetID: presetID,
		ObjectID: req.ObjectID,
	}

	path := fmt.Sprintf("/cdn/presets/%d/applied", presetID)
	if err := s.r.Request(ctx, http.MethodPost, path, req, nil); err != nil {
		return nil, fmt.Errorf("request: %w", err)
	}

	return &preset, nil
}

func (s *Service) GetAppliedPreset(ctx context.Context, presetID, objectID int) (*AppliedPreset, error) {
	var response GetAppliedPresetResponse
	preset := AppliedPreset{
		PresetID: presetID,
		ObjectID: objectID,
	}

	path := fmt.Sprintf("/cdn/presets/%d/applied", presetID)
	if err := s.r.Request(ctx, http.MethodGet, path, nil, &response); err != nil {
		return nil, fmt.Errorf("request: %w", err)
	}

	// check if objectID is in response.ObjectIDs list, if no - return error
	for _, id := range response.ObjectIDs {
		if id == objectID {
			return &preset, nil
		}
	}

	return nil, nil
}

func (s *Service) Unapply(ctx context.Context, presetID, objectID int) error {
	path := fmt.Sprintf("/cdn/presets/%d/applied/%d", presetID, objectID)

	if err := s.r.Request(ctx, http.MethodDelete, path, nil, nil); err != nil {
		return fmt.Errorf("request: %w", err)
	}

	return nil
}

func (s *Service) Get(ctx context.Context, presetID int) (*Preset, error) {
	path := fmt.Sprintf("/cdn/presets/%d", presetID)
	var preset Preset

	if err := s.r.Request(ctx, http.MethodGet, path, nil, &preset); err != nil {
		return nil, fmt.Errorf("unable to retrieve preset with id=%d. Make sure the ID provided is correct. Details: %w", presetID, err)
	}

	return &preset, nil
}
