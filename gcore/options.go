package gcore

type Options struct {
	EdgeCacheSettings *EdgeCacheSettings `json:"edge_cache_settings,omitempty"`
	HostHeader        *HostHeader        `json:"hostHeader,omitempty"`
}

type EdgeCacheSettings struct {
	Enabled      bool              `json:"enabled"`
	Value        string            `json:"value"`
	CustomValues map[string]string `json:"custom_values"`
	Default      string            `json:"default"`
}

type HostHeader struct {
	Enabled bool   `json:"enabled"`
	Value   string `json:"value"`
}
