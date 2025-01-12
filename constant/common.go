package constant

import "crypto/sha256"

const (
	TrackingDomain = "http://127.0.0.1:8080"
)

const (
	StatusDraft     = "draft"
	StatusActive    = "active"
	StatusInactive  = "inactive"
	StatusCompleted = "completed"
	StatusFailed    = "failed"
	StatusPending   = "pending"
	StatusCanceled  = "canceled"
	StatusRefunded  = "refunded"
	StatusPaid      = "paid"
)
const (
	GigCPL            = 0.3
	BannerCPC         = 0.05
	BannerImpressions = 0.001
)

var EncryptionKey = []byte("pmc-affiliate-tracker-key-32-bytes!")

func init() {
	key := sha256.Sum256([]byte("pmc-affiliate-tracker"))
	EncryptionKey = key[:]
}
