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
	UseNext           bool            `json:"use_next"`
	Sources           []SourceRequest `json:"sources,omitempty"`
	Auth              *AuthS3         `json:"auth,omitempty"`
	ProxyNextUpstream []string        `json:"proxy_next_upstream"`
}

type SourceRequest struct {
	Source             string    `json:"source,omitempty"`
	Backup             bool      `json:"backup"`
	Enabled            bool      `json:"enabled"`
	OriginType         string    `json:"origin_type,omitempty"`
	Config             *S3Config `json:"config,omitempty"`
	HostHeaderOverride *string   `json:"host_header_override,omitempty"`
}

type OriginGroup struct {
	ID                int64    `json:"id"`
	Name              string   `json:"name"`
	AuthType          string   `json:"auth_type,omitempty"`
	UseNext           bool     `json:"use_next"`
	Sources           []Source `json:"sources,omitempty"`
	Auth              *AuthS3  `json:"auth,omitempty"`
	ProxyNextUpstream []string `json:"proxy_next_upstream"`
}

type Source struct {
	Source             string    `json:"source,omitempty"`
	Backup             bool      `json:"backup"`
	Enabled            bool      `json:"enabled"`
	OriginType         string    `json:"origin_type,omitempty"`
	Config             *S3Config `json:"config,omitempty"`
	HostHeaderOverride *string   `json:"host_header_override,omitempty"`
}

type S3Config struct {
	S3Type            string `json:"s3_type"`
	S3BucketName      string `json:"s3_bucket_name"`
	S3Region          string `json:"s3_region,omitempty"`
	S3StorageHostname string `json:"s3_storage_hostname,omitempty"`
	S3AuthType        string `json:"s3_auth_type,omitempty"`
	S3AccessKeyID     string `json:"s3_access_key_id"`
	S3SecretAccessKey string `json:"s3_secret_access_key"`
}

type AuthS3 struct {
	S3Type            string `json:"s3_type"`
	S3AccessKeyID     string `json:"s3_access_key_id"`
	S3SecretAccessKey string `json:"s3_secret_access_key"`
	S3BucketName      string `json:"s3_bucket_name"`
	S3StorageHostname string `json:"s3_storage_hostname,omitempty"`
	S3Region          string `json:"s3_region,omitempty"`
}
