// service/commission_service.go
package service

import (
	"affluo/ent"
	"affluo/ent/commissionplan"
	"affluo/ent/user"
	"affluo/internal/dto"
	"context"
)

type CommissionService struct {
	client *ent.Client
}

func NewCommissionService(client *ent.Client) *CommissionService {
	return &CommissionService{
		client: client,
	}
}

func (s *CommissionService) CreateCommissionPlan(ctx context.Context, req *dto.CreateCommissionPlanRequest) (*ent.CommissionPlan, error) {

	return s.client.CommissionPlan.Create().
		SetName(req.Name).
		SetDescription(req.Description).
		SetType(commissionplan.Type(req.Type)).
		SetClickCommission(req.ClickCommission).
		SetImpressionCommission(req.ImpressionCommission).
		SetFirstLeadCommission(req.FirstLeadCommission).
		SetRepeatLeadCommission(req.RepeatLeadCommission).
		SetValidMonths(req.ValidMonths).
		SetMinimumPayout(req.MinimumPayout).
		SetIsDefault(req.IsDefault).
		SetIsActive(req.IsActive).
		Save(ctx)
}

func (s *CommissionService) AssignCommissionPlanToPublisher(ctx context.Context, planID, publisherID int64) error {
	return s.client.CommissionPlan.UpdateOneID(int(planID)).
		AddPublisherIDs(publisherID).
		Exec(ctx)
}

func (s *CommissionService) GetPublisherCommission(ctx context.Context, publisherID int64) (*ent.CommissionPlan, error) {
	return s.client.CommissionPlan.Query().
		Where(
			commissionplan.HasPublishersWith(user.IDEQ(publisherID)),
			commissionplan.IsActive(true),
		).
		First(ctx)
}
func (s *CommissionService) GetAllCommissionPlans(ctx context.Context) ([]*ent.CommissionPlan, error) {
	return s.client.CommissionPlan.Query().
		Order(ent.Desc(commissionplan.FieldID)).
		All(ctx)
}

func (s *CommissionService) GetCommissionPlanByID(ctx context.Context, planID int64) (*ent.CommissionPlan, error) {
	return s.client.CommissionPlan.Query().
		Where(commissionplan.IDEQ(int(planID))).
		First(ctx)
}

func (s *CommissionService) UpdateCommissionPlan(ctx context.Context, planID int64, req *dto.CreateCommissionPlanRequest) (*ent.CommissionPlan, error) {

	return s.client.CommissionPlan.UpdateOneID(int(planID)).
		SetName(req.Name).
		SetDescription(req.Description).
		SetType(commissionplan.Type(req.Type)).
		SetClickCommission(req.ClickCommission).
		SetImpressionCommission(req.ImpressionCommission).
		SetFirstLeadCommission(req.FirstLeadCommission).
		SetRepeatLeadCommission(req.RepeatLeadCommission).
		SetValidMonths(req.ValidMonths).
		SetMinimumPayout(req.MinimumPayout).
		SetIsDefault(req.IsDefault).
		SetIsActive(req.IsActive).
		Save(ctx)
}
