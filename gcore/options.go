package gcore

type Options struct {
	EdgeCacheSettings    *EdgeCacheSettings    `json:"edge_cache_settings"`
	BrowserCacheSettings *BrowserCacheSettings `json:"browser_cache_settings"`
	HostHeader           *HostHeader           `json:"hostHeader"`
	Webp                 *Webp                 `json:"webp"`
	Rewrite              *Rewrite              `json:"rewrite"`
	RedirectHttpToHttps  *RedirectHttpToHttps  `json:"redirect_http_to_https"`
	RequestLimiter       *RequestLimiter       `json:"request_limiter"`
	GzipOn               *GzipOn               `json:"gzipOn"`
	Cors                 *Cors                 `json:"cors"`
	SNI                  *SNIOption            `json:"sni"`
	IgnoreQueryString    *IgnoreQueryString    `json:"ignoreQueryString"`
	QueryParamsWhitelist *QueryParamsWhitelist `json:"query_params_whitelist"`
	QueryParamsBlacklist *QueryParamsBlacklist `json:"query_params_blacklist"`
	StaticRequestHeaders *StaticRequestHeaders `json:"staticRequestHeaders"`
	StaticHeaders        *StaticHeaders        `json:"staticHeaders"`
	CacheHttpHeaders     *CacheHttpHeaders     `json:"cache_http_headers"`
	WebSockets           *WebSockets           `json:"websockets"`
	TLSVersions          *TLSVersions          `json:"tls_versions"`
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
	Enabled bool   `json:"enabled"`
	Value   string `json:"value"`
}

type Webp struct {
	Enabled     bool `json:"enabled"`
	JPGQuality  int  `json:"jpg_quality"`
	PNGQuality  int  `json:"png_quality"`
	PNGLossless bool `json:"png_lossless"`
}

type Rewrite struct {
	Enabled bool   `json:"enabled"`
	Body    string `json:"body"`
	Flag    string `json:"flag"`
}

type RedirectHttpToHttps struct {
	Enabled bool `json:"enabled"`
	Value   bool `json:"value"`
}

type RequestLimiter struct {
	Enabled  bool   `json:"enabled"`
	Rate     int    `json:"rate"`
	Burst    int    `json:"burst"`
	RateUnit string `json:"rate_unit"`
	Delay    int    `json:"delay"`
}

type GzipOn struct {
	Enabled bool `json:"enabled"`
	Value   bool `json:"value"`
}

type Cors struct {
	Enabled bool     `json:"enabled"`
	Value   []string `json:"value"`
}

type SNIOption struct {
	Enabled        bool   `json:"enabled"`
	SNIType        string `json:"sni_type"`
	CustomHostname string `json:"custom_hostname"`
}

type IgnoreQueryString struct {
	Enabled bool `json:"enabled"`
	Value   bool `json:"value"`
}

type QueryParamsWhitelist struct {
	Enabled bool     `json:"enabled"`
	Value   []string `json:"value"`
}

type QueryParamsBlacklist struct {
	Enabled bool     `json:"enabled"`
	Value   []string `json:"value"`
}

type StaticRequestHeaders struct {
	Enabled bool              `json:"enabled"`
	Value   map[string]string `json:"value"`
}

type StaticHeaders struct {
	Enabled bool              `json:"enabled"`
	Value   map[string]string `json:"value"`
}

type CacheHttpHeaders struct {
	Enabled bool     `json:"enabled"`
	Value   []string `json:"value"`
}

type WebSockets struct {
	Enabled bool `json:"enabled"`
	Value   bool `json:"value"`
}

type TLSVersions struct {
	Enabled bool     `json:"enabled"`
	Value   []string `json:"value"`
}
