package motorpart

import (
	"fmt"
	"context"
)

type loggingService struct {
	service Service
}

func NewLoggingService(service Service) Service {
    fmt.Println("motorpart.newLoggingService")
	return &loggingService{
		service: service,
	}
}

func (l *loggingService) AddPart(ctx context.Context, in *AddPartInput) (*AddPartOutput, error) {
	fmt.Printf("loggingService.AddPart: event=%s\n", in.Part)
	out, err := l.service.AddPart(ctx, in)
	fmt.Printf("loggingService.AddPart: %s %s\n", out, err)
	return out, err
}

func (l *loggingService) AddParts(ctx context.Context, in *AddPartsInput) (*AddPartsOutput, error) {
	fmt.Printf("loggingService.AddParts: parts=%s\n", in.Parts)
	out, err := l.service.AddParts(ctx, in)
	fmt.Printf("loggingService.AddParts: %s %s\n", out, err)
	return out, err
}

func (l *loggingService) ListParts(ctx context.Context, in *ListPartsInput) (*ListPartsOutput, error) {
    fmt.Printf("loggingService.ListParts: filters=%s\n", in.Filters)
	out, err := l.service.ListParts(ctx, in)
    // Super verbose
	//fmt.Printf("loggingService.ListParts: %s %s\n", out, err)
	return out, err
}

func (l *loggingService) ListFilterValues(ctx context.Context, in *ListFilterValuesInput) (*ListFilterValuesOutput, error) {
    fmt.Printf("loggingService.ListFilterValues: fieldname=%s\n", in.FieldName)
	out, err := l.service.ListFilterValues(ctx, in)
    // Super verbose
	//fmt.Printf("loggingService.ListFilterValues: %s %s\n", out, err)
	return out, err
}
