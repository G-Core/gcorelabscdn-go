package logsuploader

import "context"

type LogsUploaderService interface {
	PolicyCreate(ctx context.Context, req *PolicyCreateRequest) (*Policy, error)
	PolicyList(ctx context.Context, limit, offset int) ([]Policy, error)
	PolicyGet(ctx context.Context, id int64) (*Policy, error)
	PolicyUpdate(ctx context.Context, id int64, req *PolicyUpdateRequest) (*Policy, error)
	PolicyDelete(ctx context.Context, policyID int64) error
}
