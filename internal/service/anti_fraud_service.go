// service/anti_fraud_service.go
package service

import (
	"affluo/ent"
	"affluo/ent/banner"
	"affluo/ent/lead"
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"golang.org/x/time/rate"
	"strconv"
	"sync"
	"time"
)

type AntiFraudService struct {
	client     *ent.Client
	redis      *redis.Client
	ipLimiters sync.Map // Map for IP-based rate limiting
}

type FraudCheckResult struct {
	IsValid   bool
	Reason    string
	RiskScore float64
}

type ActivityLog struct {
	IP          string
	UserAgent   string
	Timestamp   time.Time
	EventType   string
	BannerID    int64
	Fingerprint string
	Metadata    map[string]interface{}
}

func NewAntiFraudService(client *ent.Client, redis *redis.Client) *AntiFraudService {
	return &AntiFraudService{
		client: client,
		redis:  redis,
	}
}

func (s *AntiFraudService) CheckActivity(ctx context.Context, activity ActivityLog) (*FraudCheckResult, error) {
	// 1. Rate Limiting Check
	if !s.checkRateLimit(ctx, activity.IP, activity.EventType) {
		return &FraudCheckResult{
			IsValid:   false,
			Reason:    "rate_limit_exceeded",
			RiskScore: 1.0,
		}, nil
	}

	// 2. Check for suspicious patterns
	riskScore, err := s.calculateRiskScore(ctx, activity)
	if err != nil {
		return nil, err
	}

	// 3. Validate based on risk score
	if riskScore >= 0.7 {
		return &FraudCheckResult{
			IsValid:   false,
			Reason:    "high_risk_activity",
			RiskScore: riskScore,
		}, nil
	}

	return &FraudCheckResult{
		IsValid:   true,
		RiskScore: riskScore,
	}, nil
}

func (s *AntiFraudService) checkRateLimit(ctx context.Context, ip string, eventType string) bool {
	key := fmt.Sprintf("rate_limit:%s:%s", ip, eventType)

	// Get or create limiter for this IP
	limiter, _ := s.ipLimiters.LoadOrStore(key, rate.NewLimiter(rate.Every(1*time.Minute), 60))
	return limiter.(*rate.Limiter).Allow()
}

func (s *AntiFraudService) calculateRiskScore(ctx context.Context, activity ActivityLog) (float64, error) {
	var riskScore float64

	// 1. Check click velocity
	clickCount, err := s.redis.Get(ctx, fmt.Sprintf("clicks:%s:%d", activity.IP, activity.BannerID)).Int()
	if err != nil && err != redis.Nil {
		return 0, err
	}
	if clickCount > 10 {
		riskScore += 0.3
	}

	// 2. Check for suspicious timing patterns
	timingScore, err := s.checkTimingPatterns(ctx, activity)
	if err != nil {
		return 0, err
	}
	riskScore += timingScore

	// 3. Check for multiple conversions
	if activity.EventType == "lead" {
		convCount, err := s.checkMultipleConversions(ctx, activity)
		if err != nil {
			return 0, err
		}
		if convCount > 0 {
			riskScore += 0.4
		}
	}

	return riskScore, nil
}

func (s *AntiFraudService) checkTimingPatterns(ctx context.Context, activity ActivityLog) (float64, error) {
	key := fmt.Sprintf("timing:%s:%d", activity.IP, activity.BannerID)

	// Store current timestamp
	err := s.redis.RPush(ctx, key, activity.Timestamp.UnixNano()).Err()
	if err != nil {
		return 0, err
	}

	// Keep only last 10 timestamps
	s.redis.LTrim(ctx, key, -10, -1)

	// Get timestamps for analysis
	timestamps, err := s.redis.LRange(ctx, key, 0, -1).Result()
	if err != nil {
		return 0, err
	}

	// Analyze time differences
	if len(timestamps) >= 3 {
		var score float64
		var lastTime int64

		for i, ts := range timestamps {
			if i == 0 {
				lastTime, _ = strconv.ParseInt(ts, 10, 64)
				continue
			}

			currentTime, _ := strconv.ParseInt(ts, 10, 64)
			timeDiff := time.Duration(currentTime - lastTime)

			// If clicks are too regular (bot-like)
			if timeDiff < 500*time.Millisecond {
				score += 0.1
			}

			lastTime = currentTime
		}

		return score, nil
	}

	return 0, nil
}

func (s *AntiFraudService) checkMultipleConversions(ctx context.Context, activity ActivityLog) (int, error) {
	count, err := s.client.Lead.Query().
		Where(
			lead.HasBannerWith(banner.IDEQ(activity.BannerID)),
			lead.IPAddressEQ(activity.IP),
			lead.CreatedAtGT(time.Now().Add(-24*time.Hour)),
		).
		Count(ctx)

	return count, err
}

// Add to your tracking middleware
func (s *AntiFraudService) ValidateAndLogActivity(activity ActivityLog) error {
	result, err := s.CheckActivity(context.Background(), activity)
	if err != nil {
		return err
	}

	if !result.IsValid {
		return fmt.Errorf("suspicious activity detected: %s", result.Reason)
	}

	return nil
}
