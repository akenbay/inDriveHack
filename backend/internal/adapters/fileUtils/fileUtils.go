package fileUtils

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"strings"

	"inDriveHack/internal/domain"
)

type FileUtils struct{}

var _ domain.FileUtils = (*FileUtils)(nil)

func NewFileUtils() *FileUtils {
	return &FileUtils{}
}

func (f *FileUtils) ValidateImage(fileHeader *multipart.FileHeader) error {
	if fileHeader == nil || fileHeader.Size == 0 {
		return fmt.Errorf("invalid file content")
	}

	// 1. Check file size (without reading content)
	if fileHeader.Size > 5<<20 { // 5MB limit
		return fmt.Errorf("file too large (max 5MB)")
	}

	// 2. Check MIME type (reads only first 512 bytes)
	file, err := fileHeader.Open()
	if err != nil {
		return err
	}
	defer file.Close()

	buf := make([]byte, 512)
	n, err := file.Read(buf)
	if err != nil && err != io.EOF {
		return fmt.Errorf("invalid file content")
	}

	mimeType := http.DetectContentType(buf[:n])
	if !strings.HasPrefix(mimeType, "image/") {
		return fmt.Errorf("only images allowed")
	}

	return nil
}

func (f *FileUtils) FileHeaderToBytes(fileHeader *multipart.FileHeader) ([]byte, error) {
	if fileHeader == nil {
		return []byte{}, fmt.Errorf("nil fileheader")
	}
	file, err := fileHeader.Open()
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	// Read entire file (after validation)
	buf := new(bytes.Buffer)
	if _, err := io.Copy(buf, file); err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	return buf.Bytes(), nil
}
