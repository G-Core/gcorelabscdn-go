package logsuploader

import (
	"time"
)

type StorageType string

const (
	S3GcoreStorageType  StorageType = "s3_gcore"
	S3AmazonStorageType StorageType = "s3_amazon"
	S3OssStorageType    StorageType = "s3_oss"
	S3OtherStorageType  StorageType = "s3_other"
	S3V1StorageType     StorageType = "s3_v1"
	FtpStorageType      StorageType = "ftp"
	SftpStorageType     StorageType = "sftp"
	HttpStorageType     StorageType = "http"
)

type Target struct {
	ID                     int64                  `json:"id"`
	ClientID               int64                  `json:"client_id"`
	Created                time.Time              `json:"created"`
	Updated                time.Time              `json:"updated"`
	Name                   string                 `json:"name"`
	Description            string                 `json:"description"`
	StorageType            StorageType            `json:"storage_type"`
	Config                 map[string]interface{} `json:"config"`
	RelatedUploaderConfigs []int64                `json:"related_uploader_configs"`
}

type TargetCreateRequest struct {
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	StorageType StorageType            `json:"storage_type"`
	Config      map[string]interface{} `json:"config"`
}

type TargetUpdateRequest struct {
	Name        string                 `json:"name,omitempty"`
	Description string                 `json:"description"`
	StorageType StorageType            `json:"storage_type,omitempty"`
	Config      map[string]interface{} `json:"config,omitempty"`
}
