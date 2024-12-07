package dto

type TrackingRequest struct {
	CampaignID         string                 `json:"campaign_id"`
	TrackType          string                 `json:"track_type"` // click, impression, conversion
	IPAddress          string                 `json:"ip_address"`
	UserAgent          string                 `json:"user_agent"`
	Referrer           string                 `json:"referrer,omitempty"`
	ScreenResolution   string                 `json:"screen_resolution,omitempty"`
	Timezone           string                 `json:"timezone,omitempty"`
	InstalledFonts     string                 `json:"installed_fonts,omitempty"`
	IsUniqueClick      bool                   `json:"is_unique_click"`
	DeviceFingerprint  string                 `json:"device_fingerprint,omitempty"`
	AdditionalMetadata map[string]interface{} `json:"additional_metadata,omitempty"`
}
