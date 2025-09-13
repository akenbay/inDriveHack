package domain

type AnalysisResult struct {
	Intactness  float64 `json:"intactness"`
	Cleanliness float64 `json:"cleanliness"`
}

type ImageAnalyzer interface {
	Analyze(imageData []byte) (*AnalysisResult, error)
}
