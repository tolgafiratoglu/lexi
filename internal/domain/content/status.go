package content

import (
	"encoding/json"
	"fmt"
)

// Status, ContentItem için yaşam döngüsü durumunu temsil eder.
type Status string

// Tanımlı durum sabitleri (enum gibi)
const (
	StatusDraft               Status = "draft"
	StatusEnrichmentRequested Status = "enrichment_requested"
	StatusEnriched            Status = "enriched"
	StatusPublished           Status = "published"
	StatusFailed              Status = "failed"
)

// IsValidStatus, verilen değerin geçerli bir durum olup olmadığını kontrol eder.
func IsValidStatus(s Status) bool {
	switch s {
	case StatusDraft,
		StatusEnrichmentRequested,
		StatusEnriched,
		StatusPublished,
		StatusFailed:
		return true
	default:
		return false
	}
}

// MarshalJSON, Status'ü JSON'a string olarak encode eder.
func (s Status) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(s))
}

// UnmarshalJSON, JSON'dan Status okur ve geçerliyse atar.
func (s *Status) UnmarshalJSON(data []byte) error {
	var val string
	if err := json.Unmarshal(data, &val); err != nil {
		return err
	}

	status := Status(val)
	if !IsValidStatus(status) {
		return fmt.Errorf("invalid status: %s", val)
	}

	*s = status
	return nil
}
