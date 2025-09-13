package iliyasAPI

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"inDriveHack/internal/domain"
)

type IliyasAPI struct {
	port int
}

var _ domain.ImageAnalyzer = (*IliyasAPI)(nil)

func NewIliyasAPI(port int) *IliyasAPI {
	return &IliyasAPI{
		port: port,
	}
}

func (t *IliyasAPI) Analyze(imageData []byte) (*domain.AnalysisResult, error) {

	analyzeImageURL := "http://iliyasAPI:" + fmt.Sprint(t.port) + "/analyze"

	analyzeImageReq, err := http.NewRequest(http.MethodPost, analyzeImageURL, bytes.NewReader(imageData))
	if err != nil {
		return nil, err
	}

	// Send request
	client := &http.Client{}

	analyzeResp, err := client.Do(analyzeImageReq)
	if err != nil {
		slog.Error("Error when analyzing image", "error", err)
		return nil, err
	}
	defer analyzeResp.Body.Close()

	if analyzeResp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(analyzeResp.Body)
		return nil, fmt.Errorf("analyzing failed (status %d): %s", analyzeResp.StatusCode, string(body))
	}

	var result domain.AnalysisResult
	if err := json.NewDecoder(analyzeResp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode analysis response: %w", err)
	}

	return &result, nil
}
