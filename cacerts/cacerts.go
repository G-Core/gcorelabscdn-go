package cacerts

import (
	"context"
	"time"
)

type CACertService interface {
	Create(ctx context.Context, req *CreateRequest) (*CACert, error)
	Get(ctx context.Context, id int64) (*CACert, error)
	Update(ctx context.Context, id int64, req *UpdateRequest) (*CACert, error)
	Delete(ctx context.Context, id int64) error
}

type CACert struct {
	ID                  int64     `json:"id"`
	Name                string    `json:"name"`
	Deleted             bool      `json:"deleted"`
	CertIssuer          string    `json:"cert_issuer"`
	CertSubjectCN       string    `json:"cert_subject_cn"`
	CertSubjectAlt      string    `json:"cert_subject_alt"`
	ValidityNotBefore   time.Time `json:"validity_not_before"`
	ValidityNotAfter    time.Time `json:"validity_not_after"`
	HasRelatedResources bool      `json:"hasRelatedResources"`
}

type CreateRequest struct {
	Name string `json:"name"`
	Cert string `json:"sslCertificate"`
}

type UpdateRequest struct {
	Name string `json:"name"`
}
