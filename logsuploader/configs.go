package logsuploader

import (
	"time"
)

type Config struct {
	ID              int64     `json:"id"`
	ClientID        int64     `json:"client_id"`
	Created         time.Time `json:"created"`
	Updated         time.Time `json:"updated"`
	Name            string    `json:"name"`
	Enabled         bool      `json:"enabled"`
	Policy          int64     `json:"policy"`
	Target          int64     `json:"target"`
	ForAllResources bool      `json:"for_all_resources"`
	Resources       []int64   `json:"resources"`
}

type ConfigCreateRequest struct {
	Name            string  `json:"name"`
	Enabled         bool    `json:"enabled"`
	Policy          int64   `json:"policy"`
	Target          int64   `json:"target"`
	ForAllResources bool    `json:"for_all_resources"`
	Resources       []int64 `json:"resources"`
}

type ConfigUpdateRequest struct {
	Name            string  `json:"name"`
	Enabled         bool    `json:"enabled"`
	Policy          int64   `json:"policy"`
	Target          int64   `json:"target"`
	ForAllResources bool    `json:"for_all_resources"`
	Resources       []int64 `json:"resources"`
}
