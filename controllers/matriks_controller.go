package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/ihksanghazi/api-matriks/services"
)

type MatrixOperationRequest struct {
	MatrixA [][]int `json:"matrixA"`
	MatrixB [][]int `json:"matrixB"`
}

type MatrixOperationResponse struct {
	Result [][]int `json:"result"`
}

func MatriksPertambahan(w http.ResponseWriter, r *http.Request) {
	var request MatrixOperationRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result := services.AddMatrix(request.MatrixA, request.MatrixB)

	resultOperation := MatrixOperationResponse{Result: result}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resultOperation)
}

func MatriksPengurangan(w http.ResponseWriter, r *http.Request) {
	var request MatrixOperationRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result := services.SubtractMatrix(request.MatrixA, request.MatrixB)

	resultOperation := MatrixOperationResponse{Result: result}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resultOperation)
}

func MatriksPerkalian(w http.ResponseWriter, r *http.Request) {
	var request MatrixOperationRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result := services.MultiplyMatrix(request.MatrixA, request.MatrixB)

	resultOperation := MatrixOperationResponse{Result: result}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resultOperation)
}
