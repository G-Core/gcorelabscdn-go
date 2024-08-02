package invalidations

import "context"

type InvalidationsService interface {
	// Clears the cache that matches the pattern. If any error occurs, it returns an error, otherwise it returns nil.
	PurgeByPattern(ctx context.Context, resourceID int, req *PurgeByPatternRequest) error
}

type PurgeByPatternRequest struct {
	Paths []string `json:"paths"`
}
