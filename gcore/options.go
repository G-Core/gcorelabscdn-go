package gcore

type Options struct {
	EdgeCacheSettings *EdgeCacheSettings `json:"edge_cache_settings,omitempty"`
	BrowserCacheSettings *BrowserCacheSettings `json:"browser_cache_settings,omitempty"`
	HostHeader        *HostHeader        `json:"hostHeader,omitempty"`
	Webp			*Webp      	  `json:"webp,omitempty"`
	Rewrite			*Rewrite      	  `json:"rewrite,omitempty"`
	RedirectHttpToHttps			*RedirectHttpToHttps      	  `json:"redirect_http_to_https,omitempty"`
	GzipOn			*GzipOn      	  `json:"gzipOn,omitempty"`
	Cors			*Cors      	  `json:"cors,omitempty"`
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

type BrowserCacheSettings struct {
	Enabled		bool	`json:"enabled"`
	Value       string  `json:"value"`
}

type Webp struct {
	Enabled 	bool	`json:"enabled"`
	JPGQuality  int		`json:"jpg_quality"`
	PNGQuality  int		`json:"png_quality"`
	PNGLossless	bool 	`json:"png_lossless"`
}

type Rewrite struct {
	Enabled		bool    `json:"enabled"`
	Body        string  `json:"body"`
	Flag		string	`json:"flag"`
}

type RedirectHttpToHttps struct {
	Enabled     bool	`json:"enabled"`
	Value       bool    `json:"value"`
}

type GzipOn struct {
	Enabled     bool	`json:"enabled"`
	Value       bool    `json:"value"`
}

type Cors struct {
	Enabled     bool		`json:"enabled"`
	Value       []string    `json:"value"`
}
