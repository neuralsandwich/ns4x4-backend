package inmemory

import (
    "fmt"
	"github.com/halfsystems/ns4x4-backend/motorpart"
)

type repository struct {
	parts map[string]motorpart.Part
}

func New() motorpart.Repository {
    fmt.Println("inmemory.New()")
	return &repository{
		parts: map[string]motorpart.Part{},
	}
}

func (r *repository) ListParts() ([]motorpart.Part, error) {
    vals := make([]motorpart.Part, len(r.parts))

    i := 0
    for _, v := range r.parts {
        vals[i] = v
        i++
    }

	return vals, nil
}

func (r *repository) AddPart(p motorpart.Part) error {
	if p.Id == "" {
		return motorpart.ErrBadId
	}

    _, ok := r.parts[p.Id]
    if ok {
	    return motorpart.ErrPartExists
    }

	r.parts[p.Id] = p

	return nil
}
