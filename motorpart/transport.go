package motorpart

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type ListPartsHandler struct {
	service Service
}

func (h *ListPartsHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    fmt.Println("ListPartsHandler.ServeHTTP")
	// ServeHTTP aka NewServer
	ctx := req.Context()

	input := &ListPartsInput {
        Filters: req.URL.Query(),
    }
	output, err := h.service.ListParts(ctx, input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	listPartsResponse := &ListPartsResponse {
		Parts: output.Parts,
	}

	// encode response
    enableCors(&w)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(listPartsResponse)
}

type ListFilterValuesHandler struct {
	service Service
}

func (h *ListFilterValuesHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    fmt.Println("ListFilterValuesHandler.ServeHTTP")
	ctx := req.Context()

    queryValues := req.URL.Query()
    
	input := &ListFilterValuesInput{
        FieldName: queryValues.Get("field"),
    }
	output, err := h.service.ListFilterValues(ctx, input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	listFilterValuesResponse := &ListFilterValuesResponse{
		Values: output.Values,
	}

	// encode response
    enableCors(&w)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(listFilterValuesResponse)
}

type PostPartHandler struct {
	service Service
}

func (h *PostPartHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    fmt.Println("PostPartHandler.ServeHTTP")
	// ServeHTTP aka NewServer
	ctx := req.Context()

	// decode Request
	if req.Body == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var postPartRequest PostPartRequest
	if err := json.NewDecoder(req.Body).Decode(&postPartRequest); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}
	fmt.Println("Id", postPartRequest.Part.Id)
	fmt.Println("Id", postPartRequest.Part.PartName)

	input := &AddPartInput {
		Part: postPartRequest.Part,
	}
	_, err := h.service.AddPart(ctx, input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}

	postPartResponse := &PostPartResponse {}

	// encode response
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(postPartResponse)
}

type PostPartsHandler struct {
	service Service
}

func (h *PostPartsHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    fmt.Println("PostPartsHandler.ServeHTTP")
	// ServeHTTP aka NewServer
	ctx := req.Context()

	// decode Request
	if req.Body == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var postPartsRequest PostPartsRequest
	if err := json.NewDecoder(req.Body).Decode(&postPartsRequest); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}

	input := &AddPartsInput {
		Parts: postPartsRequest.Parts,
	}
	_, err := h.service.AddParts(ctx, input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}

	postPartsResponse := &PostPartsResponse {}

	// encode response
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(postPartsResponse)
}
func MakeHTTPHandler(s Service) http.Handler {
    fmt.Println("motorpart.MakeHTTPHandler")
	mux := mux.NewRouter()

	listPartsHandler := ListPartsHandler{s}
	mux.HandleFunc("/", listPartsHandler.ServeHTTP).Methods("GET")
	
    listFilterValuesHandler := ListFilterValuesHandler{s}
	mux.HandleFunc("/filtervalues", listFilterValuesHandler.ServeHTTP).Methods("GET")

	postPartHandler := PostPartHandler{s}
	mux.HandleFunc("/{id}", postPartHandler.ServeHTTP).Methods("POST")

	postPartsHandler := PostPartsHandler{s}
	mux.HandleFunc("/", postPartsHandler.ServeHTTP).Methods("POST")

	return mux
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
