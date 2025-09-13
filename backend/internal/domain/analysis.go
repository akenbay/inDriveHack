package domain

import "mime/multipart"

type AnalysisResult struct {
	Intactness  float64 `json:"intactness"`
	Cleanliness float64 `json:"cleanliness"`
}

type ImageAnalyzer interface {
	Analyze(file multipart.File) (*AnalysisResult, error)
}
