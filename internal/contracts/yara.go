package contracts

import (
	"context"
	"mime/multipart"
)

type YaraUsecase interface {
	AnalyzeFile(ctx context.Context, xApiKey string,file *multipart.FileHeader) (map[string]interface{}, error)
	GetResult(ctx context.Context, xApiKey string, id string) (map[string]interface{}, error)
}

type YaraRepository interface {
	Save(ctx context.Context, data map[string]interface{}) error
}