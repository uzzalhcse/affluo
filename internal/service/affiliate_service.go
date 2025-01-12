// internal/service/affiliate_service.go
package service

import (
	"affluo/ent"
	"affluo/ent/affiliate"
	"affluo/ent/user"
	"affluo/internal/dto"
	"context"
	"fmt"
)

type AffiliateService struct {
	client *ent.Client
}

func NewAffiliateService(client *ent.Client) *AffiliateService {
	return &AffiliateService{
		client: client,
	}
}

func (s *AffiliateService) CreateAffiliate(ctx context.Context, req *dto.CreateAffiliateRequest) (*ent.Affiliate, error) {
	// Check if user exists
	exists, err := s.client.User.Query().
		Where(user.IDEQ(req.UserID)).
		Exist(ctx)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, fmt.Errorf("user with ID %d not found", req.UserID)
	}

	return s.client.Affiliate.Create().
		SetSource(affiliate.Source(req.Source)).
		SetCommission(req.Commission).
		SetUserID(req.UserID).
		SetAffiliateUserID(req.AffiliateUserID).
		SetTrackingCode(req.TrackingCode).
		Save(ctx)
}

func (s *AffiliateService) GetUserAffiliates(ctx context.Context, userID int64) ([]*ent.Affiliate, error) {
	return s.client.Affiliate.Query().
		Where(affiliate.HasUserWith(user.IDEQ(userID))).
		All(ctx)
}

func (s *AffiliateService) UpdateAffiliate(ctx context.Context, id int64, req *dto.UpdateAffiliateRequest) (*ent.Affiliate, error) {
	return s.client.Affiliate.UpdateOneID(id).
		SetSource(affiliate.Source(req.Source)).
		SetCommission(req.Commission).
		Save(ctx)
}

func (s *AffiliateService) DeleteAffiliate(ctx context.Context, id int64) error {
	return s.client.Affiliate.DeleteOneID(id).Exec(ctx)
}
