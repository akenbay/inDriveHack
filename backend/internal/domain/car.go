package domain

type CarConditionResp struct {
	Intactness         float64 `json:"intactness"`
	IntactnessContent  string  `json:"intactness_content"`
	Cleanliness        float64 `json:"cleanliness"`
	CleanlinessContent string  `json:"cleanliness_content"`
	Conclusion         string  `json:"conclusion"`
}
