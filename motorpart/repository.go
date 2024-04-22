package motorpart

type Repository interface {
	ListParts() ([]Part, error)
	AddPart(part Part) error
}
