package sslcerts

import (
	"context"
	"time"
)

type SSLCertService interface {
	Create(ctx context.Context, req *CreateRequest) (*Cert, error)
	Get(ctx context.Context, id int64) (*Cert, error)
	Delete(ctx context.Context, id int64) error
}

type Cert struct {
	ID                  int64     `json:"id"`
	Name                string    `json:"name"`
	Deleted             bool      `json:"deleted"`
	CertIssuer          string    `json:"cert_issuer"`
	CertSubjectCN       string    `json:"cert_subjeck_cn"`
	ValidityNotBefore   time.Time `json:"validity_not_before"`
	ValidityNotAfter    time.Time `json:"validity_not_after"`
	HasRelatedResources bool      `json:"hasRelatedResources"`
	Automated           bool      `json:"automated"`
}

type CreateRequest struct {
	Name       string `json:"name"`
	Cert       string `json:"sslCertificate"`
	PrivateKey string `json:"sslPrivateKey"`
	Automated  bool   `json:"automated"`
}
