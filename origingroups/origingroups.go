package origingroups

import (
	"context"
)

type OriginGroupService interface {
	Create(ctx context.Context, req *GroupRequest) (*OriginGroup, error)
	Get(ctx context.Context, id int64) (*OriginGroup, error)
	Update(ctx context.Context, id int64, req *GroupRequest) (*OriginGroup, error)
	Delete(ctx context.Context, id int64) error
}

type GroupRequest struct {
	Name              string          `json:"name"`
	AuthType          string          `json:"auth_type,omitempty"`
	UseNext           bool            `json:"use_next,omitempty"`
	Sources           []SourceRequest `json:"sources,omitempty"`
	Auth              *AuthS3         `json:"auth,omitempty"`
	ProxyNextUpstream []string        `json:"proxy_next_upstream"`
}

type SourceRequest struct {
	Source  string `json:"source"`
	Backup  bool   `json:"backup"`
	Enabled bool   `json:"enabled"`
}

type OriginGroup struct {
	ID                int64    `json:"id"`
	Name              string   `json:"name"`
	AuthType          string   `json:"auth_type,omitempty"`
	UseNext           bool     `json:"use_next,omitempty"`
	Sources           []Source `json:"sources,omitempty"`
	Auth              *AuthS3  `json:"auth,omitempty"`
	ProxyNextUpstream []string `json:"proxy_next_upstream"`
}

type Source struct {
	Source  string `json:"source"`
	Backup  bool   `json:"backup"`
	Enabled bool   `json:"enabled"`
}

type AuthS3 struct {
	S3Type            string `json:"s3_type"`
	S3AccessKeyID     string `json:"s3_access_key_id"`
	S3SecretAccessKey string `json:"s3_secret_access_key"`
	S3BucketName      string `json:"s3_bucket_name"`
	S3StorageHostname string `json:"s3_storage_hostname,omitempty"`
	S3Region          string `json:"s3_region,omitempty"`
}
