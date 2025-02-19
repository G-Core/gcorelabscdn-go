package gcore

type Options struct {
	AllowedHTTPMethods          *AllowedHTTPMethods          `json:"allowedHttpMethods"`
	BotChallengeModule          *BotChallengeModule          `json:"bot_challenge_module"`
	BrotliCompression           *BrotliCompression           `json:"brotli_compression"`
	BrowserCacheSettings        *BrowserCacheSettings        `json:"browser_cache_settings"`
	CacheHttpHeaders            *CacheHttpHeaders            `json:"cache_http_headers"` // deprecated in favor of response_headers_hiding_policy
	Cors                        *Cors                        `json:"cors"`
	CountryACL                  *CountryACL                  `json:"country_acl"`
	DisableCache                *DisableCache                `json:"disable_cache"`
	DisableProxyForceRanges     *DisableProxyForceRanges     `json:"disable_proxy_force_ranges"`
	EdgeCacheSettings           *EdgeCacheSettings           `json:"edge_cache_settings"`
	FastEdge                    *FastEdge                    `json:"fastedge"`
	FetchCompressed             *FetchCompressed             `json:"fetch_compressed"`
	FollowOriginRedirect        *FollowOriginRedirect        `json:"follow_origin_redirect"`
	ForceReturn                 *ForceReturn                 `json:"force_return"`
	ForwardHostHeader           *ForwardHostHeader           `json:"forward_host_header"`
	GzipOn                      *GzipOn                      `json:"gzipOn"`
	HostHeader                  *HostHeader                  `json:"hostHeader"`
	HTTP3Enabled                *HTTP3Enabled                `json:"http3_enabled"`
	IgnoreCookie                *IgnoreCookie                `json:"ignore_cookie"`
	IgnoreQueryString           *IgnoreQueryString           `json:"ignoreQueryString"`
	ImageStack                  *ImageStack                  `json:"image_stack"`
	IPAddressACL                *IPAddressACL                `json:"ip_address_acl"`
	LimitBandwidth              *LimitBandwidth              `json:"limit_bandwidth"`
	ProxyCacheKey               *ProxyCacheKey               `json:"proxy_cache_key"`
	ProxyCacheMethodsSet        *ProxyCacheMethodsSet        `json:"proxy_cache_methods_set"`
	ProxyConnectTimeout         *ProxyConnectTimeout         `json:"proxy_connect_timeout"`
	ProxyReadTimeout            *ProxyReadTimeout            `json:"proxy_read_timeout"`
	QueryParamsBlacklist        *QueryParamsBlacklist        `json:"query_params_blacklist"`
	QueryParamsWhitelist        *QueryParamsWhitelist        `json:"query_params_whitelist"`
	QueryStringForwarding       *QueryStringForwarding       `json:"query_string_forwarding"`
	RedirectHttpsToHttp         *RedirectHttpsToHttp         `json:"redirect_https_to_http"`
	RedirectHttpToHttps         *RedirectHttpToHttps         `json:"redirect_http_to_https"`
	ReferrerACL                 *ReferrerACL                 `json:"referrer_acl"`
	RequestLimiter              *RequestLimiter              `json:"request_limiter"`
	ResponseHeadersHidingPolicy *ResponseHeadersHidingPolicy `json:"response_headers_hiding_policy"`
	Rewrite                     *Rewrite                     `json:"rewrite"`
	SecureKey                   *SecureKey                   `json:"secure_key"`
	Slice                       *Slice                       `json:"slice"`
	SNI                         *SNIOption                   `json:"sni"`
	Stale                       *Stale                       `json:"stale"`
	StaticHeaders               *StaticHeaders               `json:"staticHeaders"` // deprecated in favor of static_response_headers
	StaticRequestHeaders        *StaticRequestHeaders        `json:"staticRequestHeaders"`
	StaticResponseHeaders       *StaticResponseHeaders       `json:"static_response_headers"`
	TLSVersions                 *TLSVersions                 `json:"tls_versions"`
	UseDefaultLEChain           *UseDefaultLEChain           `json:"use_default_le_chain"`
	UseDNS01LEChallenge         *UseDNS01LEChallenge         `json:"use_dns01_le_challenge"`
	UserAgentACL                *UserAgentACL                `json:"user_agent_acl"`
	UseRSALECert                *UseRSALECert                `json:"use_rsa_le_cert"`
	WAAP                        *WAAP                        `json:"waap"`
	WAF                         *WAF                         `json:"waf"`
	WebSockets                  *WebSockets                  `json:"websockets"`
}

type AllowedHTTPMethods struct {
	Enabled bool     `json:"enabled"`
	Value   []string `json:"value"`
}

type BotChallengeModule struct {
	Enabled bool `json:"enabled"`
	Value   bool `json:"value"`
}

type BrotliCompression struct {
	Enabled bool     `json:"enabled"`
	Value   []string `json:"value"`
}

type BrowserCacheSettings struct {
	Enabled bool   `json:"enabled"`
	Value   string `json:"value"`
}

// deprecated in favor of ResponseHeadersHidingPolicy
type CacheHttpHeaders struct {
	Enabled bool     `json:"enabled"`
	Value   []string `json:"value"`
}

type Cors struct {
	Enabled bool     `json:"enabled"`
	Value   []string `json:"value"`
	Always  bool     `json:"always"`
}

type CountryACL struct {
	Enabled        bool     `json:"enabled"`
	PolicyType     string   `json:"policy_type"`
	ExceptedValues []string `json:"excepted_values"`
}

type DisableCache struct {
	Enabled bool `json:"enabled"`
	Value   bool `json:"value"`
}

type DisableProxyForceRanges struct {
	Enabled bool `json:"enabled"`
	Value   bool `json:"value"`
}

type EdgeCacheSettings struct {
	Enabled      bool              `json:"enabled"`
	Value        string            `json:"value"`
	CustomValues map[string]string `json:"custom_values"`
	Default      string            `json:"default"`
}

type FastEdgeAppConfig struct {
	Enabled          bool   `json:"enabled"`
	AppID            string `json:"app_id"`
	InterruptOnError bool   `json:"interrupt_on_error"`
}

type FastEdge struct {
	Enabled          bool               `json:"enabled"`
	OnRequestHeaders *FastEdgeAppConfig `json:"on_request_headers"`
}

type FetchCompressed struct {
	Enabled bool `json:"enabled"`
	Value   bool `json:"value"`
}

type FollowOriginRedirect struct {
	Enabled bool  `json:"enabled"`
	Codes   []int `json:"codes"`
}

type ForceReturn struct {
	Enabled bool   `json:"enabled"`
	Code    int    `json:"code"`
	Body    string `json:"body"`
}

type ForwardHostHeader struct {
	Enabled bool `json:"enabled"`
	Value   bool `json:"value"`
}

type GzipOn struct {
	Enabled bool `json:"enabled"`
	Value   bool `json:"value"`
}

type HostHeader struct {
	Enabled bool   `json:"enabled"`
	Value   string `json:"value"`
}

type HTTP3Enabled struct {
	Enabled bool `json:"enabled"`
	Value   bool `json:"value"`
}

type IgnoreCookie struct {
	Enabled bool `json:"enabled"`
	Value   bool `json:"value"`
}

type IgnoreQueryString struct {
	Enabled bool `json:"enabled"`
	Value   bool `json:"value"`
}

type ImageStack struct {
	Enabled     bool `json:"enabled"`
	AvifEnabled bool `json:"avif_enabled"`
	WebpEnabled bool `json:"webp_enabled"`
	Quality     int  `json:"quality"`
	PngLossless bool `json:"png_lossless"`
}

type IPAddressACL struct {
	Enabled        bool     `json:"enabled"`
	PolicyType     string   `json:"policy_type"`
	ExceptedValues []string `json:"excepted_values"`
}

type LimitBandwidth struct {
	Enabled   bool   `json:"enabled"`
	LimitType string `json:"limit_type"`
	Speed     int    `json:"speed"`
	Buffer    int    `json:"buffer"`
}

type ProxyCacheKey struct {
	Enabled bool   `json:"enabled"`
	Value   string `json:"value"`
}

type ProxyCacheMethodsSet struct {
	Enabled bool `json:"enabled"`
	Value   bool `json:"value"`
}

type ProxyConnectTimeout struct {
	Enabled bool   `json:"enabled"`
	Value   string `json:"value"`
}

type ProxyReadTimeout struct {
	Enabled bool   `json:"enabled"`
	Value   string `json:"value"`
}

type QueryParamsBlacklist struct {
	Enabled bool     `json:"enabled"`
	Value   []string `json:"value"`
}

type QueryParamsWhitelist struct {
	Enabled bool     `json:"enabled"`
	Value   []string `json:"value"`
}

type QueryStringForwarding struct {
    Enabled              bool     `json:"enabled"`
    ForwardFromFileTypes []string `json:"forward_from_file_types"`
    ForwardToFileTypes   []string `json:"forward_to_file_types"`
}

type RedirectHttpsToHttp struct {
	Enabled bool `json:"enabled"`
	Value   bool `json:"value"`
}

type RedirectHttpToHttps struct {
	Enabled bool `json:"enabled"`
	Value   bool `json:"value"`
}

type ReferrerACL struct {
	Enabled        bool     `json:"enabled"`
	PolicyType     string   `json:"policy_type"`
	ExceptedValues []string `json:"excepted_values"`
}

type RequestLimiter struct {
	Enabled  bool   `json:"enabled"`
	Rate     int    `json:"rate"`
	Burst    int    `json:"burst"`
	RateUnit string `json:"rate_unit"`
	Delay    int    `json:"delay"`
}

type ResponseHeadersHidingPolicy struct {
	Enabled  bool     `json:"enabled"`
	Mode     string   `json:"mode"`
	Excepted []string `json:"excepted"`
}

type Rewrite struct {
	Enabled bool   `json:"enabled"`
	Body    string `json:"body"`
	Flag    string `json:"flag"`
}

type SecureKey struct {
	Enabled bool   `json:"enabled"`
	Key     string `json:"key"`
	Type    int    `json:"type"`
}

type Slice struct {
	Enabled bool `json:"enabled"`
	Value   bool `json:"value"`
}

type SNIOption struct {
	Enabled        bool   `json:"enabled"`
	SNIType        string `json:"sni_type"`
	CustomHostname string `json:"custom_hostname"`
}

type Stale struct {
	Enabled bool     `json:"enabled"`
	Value   []string `json:"value"`
}

// deprecated in favor of StaticResponseHeaders
type StaticHeaders struct {
	Enabled bool              `json:"enabled"`
	Value   map[string]string `json:"value"`
}

type StaticRequestHeaders struct {
	Enabled bool              `json:"enabled"`
	Value   map[string]string `json:"value"`
}

type StaticResponseHeadersItem struct {
	Name   string   `json:"name"`
	Value  []string `json:"value"`
	Always bool     `json:"always"`
}

type StaticResponseHeaders struct {
	Enabled bool                        `json:"enabled"`
	Value   []StaticResponseHeadersItem `json:"value"`
}

type TLSVersions struct {
	Enabled bool     `json:"enabled"`
	Value   []string `json:"value"`
}

type UseDefaultLEChain struct {
	Enabled bool `json:"enabled"`
	Value   bool `json:"value"`
}

type UseDNS01LEChallenge struct {
	Enabled bool `json:"enabled"`
	Value   bool `json:"value"`
}

type UserAgentACL struct {
	Enabled        bool     `json:"enabled"`
	PolicyType     string   `json:"policy_type"`
	ExceptedValues []string `json:"excepted_values"`
}

type UseRSALECert struct {
	Enabled bool `json:"enabled"`
	Value   bool `json:"value"`
}

type WAAP struct {
	Enabled bool `json:"enabled"`
	Value   bool `json:"value"`
}

type WAF struct {
	Enabled bool `json:"enabled"`
	Value   bool `json:"value"`
}

type WebSockets struct {
	Enabled bool `json:"enabled"`
	Value   bool `json:"value"`
}
