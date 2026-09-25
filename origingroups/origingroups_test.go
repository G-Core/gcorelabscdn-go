package origingroups

import (
	"encoding/json"
	"testing"
)

func TestS3ConfigMarshalGcoreOmitsCredentials(t *testing.T) {
	b, err := json.Marshal(S3Config{S3Type: "gcore", S3BucketName: "bucket", S3AuthType: "awsSignatureV4", StorageID: 123})
	if err != nil {
		t.Fatal(err)
	}

	var got map[string]interface{}
	if err := json.Unmarshal(b, &got); err != nil {
		t.Fatal(err)
	}

	for _, key := range []string{"s3_access_key_id", "s3_secret_access_key", "s3_region", "s3_storage_hostname"} {
		if _, ok := got[key]; ok {
			t.Errorf("%s must be omitted, got %s", key, b)
		}
	}
	if got["storage_id"] != float64(123) {
		t.Errorf("storage_id = %v, want 123", got["storage_id"])
	}
}

func TestS3ConfigMarshalManualOmitsStorageID(t *testing.T) {
	b, err := json.Marshal(S3Config{S3Type: "amazon", S3BucketName: "bucket", S3Region: "us-east-1", S3AccessKeyID: "ak", S3SecretAccessKey: "sk"})
	if err != nil {
		t.Fatal(err)
	}

	var got map[string]interface{}
	if err := json.Unmarshal(b, &got); err != nil {
		t.Fatal(err)
	}

	if _, ok := got["storage_id"]; ok {
		t.Errorf("storage_id must be omitted, got %s", b)
	}
	if got["s3_access_key_id"] != "ak" || got["s3_secret_access_key"] != "sk" {
		t.Errorf("credentials not marshalled: %s", b)
	}
}

func TestS3ConfigUnmarshalGcoreResponse(t *testing.T) {
	var cfg S3Config
	if err := json.Unmarshal([]byte(`{"s3_type":"gcore","storage_id":42,"s3_bucket_name":"b","s3_auth_type":"awsSignatureV4"}`), &cfg); err != nil {
		t.Fatal(err)
	}
	if cfg.StorageID != 42 || cfg.S3Type != "gcore" || cfg.S3BucketName != "b" {
		t.Errorf("unexpected config: %+v", cfg)
	}
}
