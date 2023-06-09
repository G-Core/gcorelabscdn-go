package gcore

type Options struct {
	BrowserCacheSettings *BrowserCacheSettings `json:"browser_cache_settings"`
	CacheHttpHeaders     *CacheHttpHeaders     `json:"cache_http_headers"` // deprecated in favor of response_headers_hiding_policy
	Cors                 *Cors                 `json:"cors"`
	EdgeCacheSettings    *EdgeCacheSettings    `json:"edge_cache_settings"`
	ForceReturn          *ForceReturn          `json:"force_return"`
	GzipOn               *GzipOn               `json:"gzipOn"`
	HostHeader           *HostHeader           `json:"hostHeader"`
	IgnoreQueryString    *IgnoreQueryString    `json:"ignoreQueryString"`
	QueryParamsBlacklist *QueryParamsBlacklist `json:"query_params_blacklist"`
	QueryParamsWhitelist *QueryParamsWhitelist `json:"query_params_whitelist"`
	RedirectHttpToHttps  *RedirectHttpToHttps  `json:"redirect_http_to_https"`
	RequestLimiter       *RequestLimiter       `json:"request_limiter"`
	Rewrite              *Rewrite              `json:"rewrite"`
	SNI                  *SNIOption            `json:"sni"`
	StaticHeaders        *StaticHeaders        `json:"staticHeaders"` // deprecated in favor of static_response_headers
	StaticRequestHeaders *StaticRequestHeaders `json:"staticRequestHeaders"`
	TLSVersions          *TLSVersions          `json:"tls_versions"`
	UseRSALECert         *UseRSALECert         `json:"use_rsa_le_cert"`
	Webp                 *Webp                 `json:"webp"` // deprecated option
	WebSockets           *WebSockets           `json:"websockets"`
}

type BrowserCacheSettings struct {
	Enabled bool   `json:"enabled"`
	Value   string `json:"value"`
}

// deprecated in favor of response_headers_hiding_policy
type CacheHttpHeaders struct {
	Enabled bool     `json:"enabled"`
	Value   []string `json:"value"`
}

type Cors struct {
	Enabled bool     `json:"enabled"`
	Value   []string `json:"value"`
}

type EdgeCacheSettings struct {
	Enabled      bool              `json:"enabled"`
	Value        string            `json:"value"`
	CustomValues map[string]string `json:"custom_values"`
	Default      string            `json:"default"`
}

type ForceReturn struct {
	Enabled bool   `json:"enabled"`
	Code    int    `json:"code"`
	Body    string `json:"body"`
}

type GzipOn struct {
	Enabled bool `json:"enabled"`
	Value   bool `json:"value"`
}

type HostHeader struct {
	Enabled bool   `json:"enabled"`
	Value   string `json:"value"`
}

type IgnoreQueryString struct {
	Enabled bool `json:"enabled"`
	Value   bool `json:"value"`
}

type QueryParamsBlacklist struct {
	Enabled bool     `json:"enabled"`
	Value   []string `json:"value"`
}

type QueryParamsWhitelist struct {
	Enabled bool     `json:"enabled"`
	Value   []string `json:"value"`
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

type Rewrite struct {
	Enabled bool   `json:"enabled"`
	Body    string `json:"body"`
	Flag    string `json:"flag"`
}

type SNIOption struct {
	Enabled        bool   `json:"enabled"`
	SNIType        string `json:"sni_type"`
	CustomHostname string `json:"custom_hostname"`
}

// deprecated in favor of static_response_headers
type StaticHeaders struct {
	Enabled bool              `json:"enabled"`
	Value   map[string]string `json:"value"`
}

type StaticRequestHeaders struct {
	Enabled bool              `json:"enabled"`
	Value   map[string]string `json:"value"`
}

type TLSVersions struct {
	Enabled bool     `json:"enabled"`
	Value   []string `json:"value"`
}

type UseRSALECert struct {
	Enabled bool `json:"enabled"`
	Value   bool `json:"value"`
}

// deprecated option
type Webp struct {
	Enabled     bool `json:"enabled"`
	JPGQuality  int  `json:"jpg_quality"`
	PNGQuality  int  `json:"png_quality"`
	PNGLossless bool `json:"png_lossless"`
}

type WebSockets struct {
	Enabled bool `json:"enabled"`
	Value   bool `json:"value"`
}
