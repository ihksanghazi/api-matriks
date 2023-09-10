package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ihksanghazi/api-matriks/services"
)

type MatrixOperationRequest struct {
	MatrixA [][]int `json:"matrixA"`
	MatrixB [][]int `json:"matrixB"`
}

type MatrixOperationSingleRequest struct {
	Matrix [][]float64 `json:"matrix"`
}

type MatrixOperationResponse struct {
	Result interface{} `json:"result"`
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

func MatriksTranspose(w http.ResponseWriter, r *http.Request) {
	var request MatrixOperationSingleRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result := services.TransposeMatrix(request.Matrix)

	resultOperation := MatrixOperationResponse{Result: result}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resultOperation)
}

func MatriksInverse(w http.ResponseWriter, r *http.Request) {
	var request MatrixOperationSingleRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, errRes := services.InverseMatriks(request.Matrix)
	if errRes != nil {
		http.Error(w, errRes.Error(), http.StatusInternalServerError)
		return
	}

	resultOperation := MatrixOperationResponse{Result: result}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resultOperation)
}

func MatriksDeterminan(w http.ResponseWriter, r *http.Request) {
	var request MatrixOperationSingleRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result := services.DeterminantMatriks(request.Matrix)

	resultOperation := MatrixOperationResponse{Result: result}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resultOperation)
}

func MatriksReduce(w http.ResponseWriter, r *http.Request) {
	var request MatrixOperationSingleRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result := services.ReduceMatrix(request.Matrix)

	resultOperation := MatrixOperationResponse{Result: result}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resultOperation)
}

func MatriksCreateIdentity(w http.ResponseWriter, r *http.Request) {
	size := r.URL.Query().Get("size")

	value, err := strconv.Atoi(size)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result := services.CreateIdentityMatrix(value)

	resultOperation := MatrixOperationResponse{Result: result}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resultOperation)
}

func MatriksCreateDiagonal(w http.ResponseWriter, r *http.Request) {
	size := r.URL.Query().Get("size")
	diagonal := r.URL.Query().Get("diagonal")

	valueSize, errSize := strconv.Atoi(size)
	if errSize != nil {
		http.Error(w, errSize.Error(), http.StatusBadRequest)
		return
	}

	valueDiagonal, errDiagonal := strconv.Atoi(diagonal)
	if errDiagonal != nil {
		http.Error(w, errDiagonal.Error(), http.StatusBadRequest)
		return
	}

	result := services.CreateDiagonalMatrix(valueSize, valueDiagonal)

	resultOperation := MatrixOperationResponse{Result: result}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resultOperation)
}
