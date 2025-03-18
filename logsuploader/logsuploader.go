package logsuploader

import "context"

type LogsUploaderService interface {
	PolicyCreate(ctx context.Context, req *PolicyCreateRequest) (*Policy, error)
	PolicyList(ctx context.Context, limit, offset int) ([]Policy, error)
	PolicyGet(ctx context.Context, id int64) (*Policy, error)
	PolicyUpdate(ctx context.Context, id int64, req *PolicyUpdateRequest) (*Policy, error)
	PolicyDelete(ctx context.Context, policyID int64) error
	TargetCreate(ctx context.Context, req *TargetCreateRequest) (*Target, error)
	TargetList(ctx context.Context, limit, offset int) ([]Target, error)
	TargetGet(ctx context.Context, id int64) (*Target, error)
	TargetUpdate(ctx context.Context, id int64, req *TargetUpdateRequest) (*Target, error)
	TargetDelete(ctx context.Context, targetID int64) error
}
