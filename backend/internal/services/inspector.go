package services

import (
	"inDriveHack/internal/domain"
	"mime/multipart"
)

type InspectorService struct {
	analyzer domain.ImageAnalyzer
}

func NewInspectorService(analyzer domain.ImageAnalyzer) *InspectorService {
	return &InspectorService{analyzer: analyzer}
}

func (s *InspectorService) Inspect(file multipart.File) (*domain.CarConditionResp, error) {
	analyzeRes, err := s.analyzer.Analyze(file)
	if err != nil {
		return nil, err
	}

	result := domain.CarConditionResp{
		Intactness:  analyzeRes.Intactness,
		Cleanliness: analyzeRes.Cleanliness,
	}

	if result.Intactness >= 0.8 {
		result.IntactnessContent = "Машина не сильно повреждена."
	} else if result.Intactness >= 0.5 {
		result.IntactnessContent = "Машина повреждена"
	} else {
		result.IntactnessContent = "Машина сильно повреждена"
	}

	if result.Cleanliness >= 0.8 {
		result.CleanlinessContent = "Машина не сильно загрязнена."
	} else if result.Cleanliness >= 0.5 {
		result.CleanlinessContent = "Машина загрязнена"
	} else {
		result.CleanlinessContent = "Машина сильно загрязнена"
	}

	// Add a business rule: compute conclusion
	if result.Intactness >= 0.8 && result.Cleanliness >= 0.8 {
		result.Conclusion = "Машина в отличном состоянии"
	} else if result.Intactness >= 0.5 && result.Cleanliness >= 0.5 {
		result.Conclusion = "Машина в среднем состоянии"
	} else {
		result.Conclusion = "Машина в ужасном состоянии"
	}

	return &result, nil
}
