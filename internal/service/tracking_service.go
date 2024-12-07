package service

import (
	"affluo/ent"
	"affluo/ent/track"
	"affluo/internal/dto"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/google/uuid"
	"net"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/ua-parser/uap-go/uaparser"
)

type TrackingService struct {
	client *ent.Client
	redis  *redis.Client
	ua     *uaparser.Parser
}

func NewTrackingService(client *ent.Client, redisClient *redis.Client) *TrackingService {
	uaParser, _ := uaparser.New("path/to/regexes.yaml")
	return &TrackingService{
		client: client,
		redis:  redisClient,
		ua:     uaParser,
	}
}

// GenerateDeviceFingerprint creates a probabilistic device fingerprint
func (s *TrackingService) GenerateDeviceFingerprint(ctx context.Context, request *dto.TrackingRequest) string {
	// Combine multiple attributes for fingerprinting
	components := []string{
		request.IPAddress,
		request.UserAgent,
		request.ScreenResolution,
		request.Timezone,
		request.InstalledFonts,
	}

	hash := sha256.New()
	hash.Write([]byte(strings.Join(components, "|")))
	return hex.EncodeToString(hash.Sum(nil))
}

// DetectFraudulentActivity checks for suspicious tracking activities
func (s *TrackingService) DetectFraudulentActivity(ctx context.Context, track *dto.TrackingRequest) bool {
	// Multiple fraud detection mechanisms
	if s.isBlacklisted(ctx, track.IPAddress) {
		return true
	}

	if s.isDuplicateClick(ctx, track) {
		return true
	}

	if s.hasExcessiveClicks(ctx, track) {
		return true
	}

	return false
}

func (s *TrackingService) isBlacklisted(ctx context.Context, ipAddress string) bool {
	// Check IP and device blacklist in Redis
	ip := net.ParseIP(ipAddress)
	if ip == nil {
		return true
	}

	// Check Redis blacklist
	blacklisted, _ := s.redis.SIsMember(ctx, "ip_blacklist", ipAddress).Result()
	return blacklisted
}

func (s *TrackingService) isDuplicateClick(ctx context.Context, track *dto.TrackingRequest) bool {
	// Use Redis to check for duplicate clicks within a timeframe
	key := fmt.Sprintf("click:%s:%s", track.CampaignID, track.DeviceFingerprint)
	exists, _ := s.redis.Exists(ctx, key).Result()
	return exists == 1
}

func (s *TrackingService) hasExcessiveClicks(ctx context.Context, track *dto.TrackingRequest) bool {
	// Rate limiting logic using Redis
	key := fmt.Sprintf("clicks:%s", track.IPAddress)
	clicks, _ := s.redis.Incr(ctx, key).Result()
	s.redis.Expire(ctx, key, 1*time.Hour)

	return clicks > 100 // Adjust threshold as needed
}

// Track records a tracking event
func (s *TrackingService) Track(ctx context.Context, request *dto.TrackingRequest) (*ent.Track, error) {
	// Analyze user agent
	client := s.ua.Parse(request.UserAgent)

	// Generate device fingerprint
	fingerprint := s.GenerateDeviceFingerprint(ctx, request)

	// Detect fraudulent activity
	isFraudulent := s.DetectFraudulentActivity(ctx, request)

	// Create tracking record
	trackStatus := "valid"
	if isFraudulent {
		trackStatus = "suspected_fraud"
	}

	return s.client.Track.Create().
		SetIPAddress(request.IPAddress).
		SetUserAgent(request.UserAgent).
		SetDeviceFingerprint(fingerprint).
		SetType(track.Type(request.TrackType)).
		SetStatus(track.Status(trackStatus)).
		SetIsUniqueClick(request.IsUniqueClick).
		SetAdditionalMetadata(map[string]interface{}{
			"browser": client.UserAgent.Family,
			"os":      client.Os.Family,
			"device":  client.Device.Family,
		}).
		Save(ctx)
}

// GenerateTrackingScript creates a unique tracking script
func (s *TrackingService) GenerateTrackingScript(campaignID string) string {
	trackingCode := uuid.New().String()

	return fmt.Sprintf(`
<script>
(function() {
	window.trackEvent = function(eventType) {
		fetch('/track', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({
				campaignId: '%s',
				trackingCode: '%s',
				eventType: eventType,
				timestamp: new Date().toISOString()
			})
		});
	};

	// Auto-track page impression
	trackEvent('impression');

	// Track outbound links
	document.addEventListener('click', function(e) {
		if (e.target.tagName === 'A') {
			trackEvent('click');
		}
	});
})();
</script>
`, campaignID, trackingCode)
}
