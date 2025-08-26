package logsuploader

import (
	"time"
)

type Policy struct {
	ID                     int64             `json:"id"`
	ClientID               int64             `json:"client_id"`
	Created                time.Time         `json:"created"`
	Updated                time.Time         `json:"updated"`
	IncludeEmptyLogs       bool              `json:"include_empty_logs"`
	IncludeShieldLogs      bool              `json:"include_shield_logs"`
	Name                   string            `json:"name"`
	Description            string            `json:"description"`
	RetryIntervalMinutes   int               `json:"retry_interval_minutes"`
	RotateIntervalMinutes  int               `json:"rotate_interval_minutes"`
	RotateThresholdMB      *int              `json:"rotate_threshold_mb"`
	RotateThresholdLines   int               `json:"rotate_threshold_lines"`
	DateFormat             string            `json:"date_format"`
	FieldDelimiter         string            `json:"field_delimiter"`
	FieldSeparator         string            `json:"field_separator"`
	Fields                 []string          `json:"fields"`
	FileNameTemplate       string            `json:"file_name_template"`
	FormatType             string            `json:"format_type"`
	Tags                   map[string]string `json:"tags"`
	RelatedUploaderConfigs []int64           `json:"related_uploader_configs"`
}

type PolicyCreateRequest struct {
	IncludeEmptyLogs      bool              `json:"include_empty_logs,omitempty"`
	IncludeShieldLogs     bool              `json:"include_shield_logs,omitempty"`
	Name                  string            `json:"name,omitempty"`
	Description           string            `json:"description,omitempty"`
	RetryIntervalMinutes  int               `json:"retry_interval_minutes,omitempty"`
	RotateIntervalMinutes int               `json:"rotate_interval_minutes,omitempty"`
	RotateThresholdMB     *int              `json:"rotate_threshold_mb"`
	RotateThresholdLines  int               `json:"rotate_threshold_lines,omitempty"`
	DateFormat            string            `json:"date_format,omitempty"`
	FieldDelimiter        string            `json:"field_delimiter,omitempty"`
	FieldSeparator        string            `json:"field_separator,omitempty"`
	Fields                []string          `json:"fields,omitempty"`
	FileNameTemplate      string            `json:"file_name_template,omitempty"`
	FormatType            string            `json:"format_type,omitempty"`
	Tags                  map[string]string `json:"tags,omitempty"`
}

type PolicyUpdateRequest struct {
	IncludeEmptyLogs      bool              `json:"include_empty_logs"`
	IncludeShieldLogs     bool              `json:"include_shield_logs"`
	Name                  string            `json:"name,omitempty"`
	Description           string            `json:"description"`
	RetryIntervalMinutes  int               `json:"retry_interval_minutes,omitempty"`
	RotateIntervalMinutes int               `json:"rotate_interval_minutes,omitempty"`
	RotateThresholdMB     *int              `json:"rotate_threshold_mb"`
	RotateThresholdLines  int               `json:"rotate_threshold_lines"`
	DateFormat            string            `json:"date_format"`
	FieldDelimiter        string            `json:"field_delimiter,omitempty"`
	FieldSeparator        string            `json:"field_separator,omitempty"`
	Fields                []string          `json:"fields,omitempty"`
	FileNameTemplate      string            `json:"file_name_template,omitempty"`
	FormatType            string            `json:"format_type"`
	Tags                  map[string]string `json:"tags"`
}
