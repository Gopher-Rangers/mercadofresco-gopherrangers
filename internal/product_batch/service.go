package productbatch

import (
	"context"
	"fmt"
)

type Services interface {
	Report(ctx context.Context) ([]Report, error)
	ReportByID(ctx context.Context, id int) (Report, error)
	Create(ctx context.Context, pb ProductBatch) (ProductBatch, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Services {
	s := service{r}
	return &s
}

func (s service) Create(ctx context.Context, pb ProductBatch) (ProductBatch, error) {
	_, err := s.repository.ReportByID(ctx, pb.BatchNumber)
	if err == nil {
		return ProductBatch{}, fmt.Errorf("error: batch number '%d' already exists in BD", pb.BatchNumber)
	}

	pb, err = s.repository.Create(ctx, pb)
	if err != nil {
		return ProductBatch{}, err
	}
	return pb, nil
}

func (s service) Report(ctx context.Context) ([]Report, error) {
	pb, err := s.repository.Report(ctx)
	if err != nil {
		return []Report{}, err
	}
	return pb, nil
}

func (s service) ReportByID(ctx context.Context, id int) (Report, error) {
	pb, err := s.repository.ReportByID(ctx, id)
	if err != nil {
		return Report{}, err
	}
	return pb, nil
}
