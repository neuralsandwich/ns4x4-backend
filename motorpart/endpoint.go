package motorpart

type ListPartsRequest struct {}

type ListPartsResponse struct {
	Parts []Part `json:"parts"`
}

type ListFilterValuesRequest struct {}

type ListFilterValuesResponse struct {
	Values []string `json:"values"`
}

type PostPartRequest struct {
	Part Part `json:"part"`
}

type PostPartResponse struct {
	Error string `json:"error,omitempty"`
}

type PostPartsRequest struct {
	Parts []Part `json:"parts"`
}

type PostPartsResponse struct {
	Error string `json:"error,omitempty"`
}
