package motorpart

import (
    "fmt"
    "context"
    "net/url"
    "reflect"
)

//------------------------------------------------------------------------------
// Part Service
//------------------------------------------------------------------------------

type Service interface {
	AddPart(context.Context, *AddPartInput) (*AddPartOutput, error)
	AddParts(context.Context, *AddPartsInput) (*AddPartsOutput, error)
	ListParts(context.Context, *ListPartsInput) (*ListPartsOutput, error)
    ListFilterValues(context.Context,
        *ListFilterValuesInput) (*ListFilterValuesOutput, error)
}

type service struct {
	repository Repository
}

func New(repository Repository) Service {
    fmt.Println("motorpart.New()")
	return &service{
		repository: repository,
	}
}

//------------------------------------------------------------------------------
// ListParts
//------------------------------------------------------------------------------

type ListPartsInput struct {
    Filters url.Values
}

type ListPartsOutput struct {
	Parts []Part
}

func matchesFilter(filters map[string]string, part Part) bool {
    pv := reflect.ValueOf(part)

    for k, v := range filters {
        if pv.FieldByName(k).String() != v {
            return false
        }
    }

    return true
}

func (s *service) ListParts(_ context.Context, input *ListPartsInput) (*ListPartsOutput, error) {
	parts, err := s.repository.ListParts()
	if err != nil {
		return nil, err
	}

    filters := make(map[string]string)
    for k, v := range input.Filters {
        if len(v) == 0 {
            continue
        }
        if v[0] == "" {
            continue
        }
        filters[k] = v[0]
    }
    fmt.Println(filters)

    i := 0
    for _, p := range parts {
        if matchesFilter(filters, p) {
            parts[i] = p
            i++
        }
    }
    parts = parts[:i]


	return &ListPartsOutput{
		Parts: parts,
	}, nil
}

//------------------------------------------------------------------------------
// AddPart
//------------------------------------------------------------------------------

type AddPartInput struct {
	Part Part `json:"part"`
}

type AddPartOutput struct {
	Error string `json:"error"`
}

func (s *service) AddPart(_ context.Context, input *AddPartInput) (*AddPartOutput, error) {
	err := s.repository.AddPart(input.Part)
	if err != nil {
		return nil, err
	}

	return &AddPartOutput{}, nil
}

//------------------------------------------------------------------------------
// AddParts
//------------------------------------------------------------------------------

type AddPartsInput struct {
	Parts []Part `json:"parts"`
}

type AddPartsOutput struct {
	Error string `json:"error"`
}

func (s *service) AddParts(_ context.Context, input *AddPartsInput) (*AddPartsOutput, error) {
    for _, part := range input.Parts {
	    err := s.repository.AddPart(part)
	    if err != nil {
	    	return nil, err
	    }
    }

	return &AddPartsOutput{}, nil
}

//------------------------------------------------------------------------------
// ListFilterValues
//------------------------------------------------------------------------------

type ListFilterValuesInput struct {
    FieldName string `json:"field"`
}

type ListFilterValuesOutput struct {
	Values []string
}

func (s *service) ListFilterValues(_ context.Context,
    input *ListFilterValuesInput) (*ListFilterValuesOutput, error) {

    if input.FieldName == "" {
        return nil, ErrBadField
    }

	parts, err := s.repository.ListParts()
    if err != nil {
        return nil, err
    }

    valueSet := make(map[string]bool)
    for _, p := range parts {
        pv := reflect.ValueOf(p)
        value := pv.FieldByName(input.FieldName).String()

        if exists := valueSet[value]; !exists {
            valueSet[value] = true
        }
    }

    values := make([]string, len(valueSet))
    
    i := 0
    for k := range valueSet {
        values[i] = k
        i++
    }
	return &ListFilterValuesOutput{
		Values: values,
	}, nil
}
