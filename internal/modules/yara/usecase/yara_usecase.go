package usecase

import (
	"be-interview-task/internal/apiclient/virustotalservice"
	"be-interview-task/internal/contracts"
	"context"
	"mime/multipart"
	"time"
)

type YaraUsecase struct {
	contextTimeout time.Duration
	yaraRepository contracts.YaraRepository
}

func NewYaraUsecase(timeout time.Duration, yr contracts.YaraRepository) contracts.YaraUsecase {
	return &YaraUsecase{
		contextTimeout: timeout,
		yaraRepository: yr,
	}
}

func (y *YaraUsecase) AnalyzeFile(ctx context.Context, xApiKey string, file *multipart.FileHeader) (map[string]interface{}, error) {
	ctx, cancel := context.WithTimeout(ctx, y.contextTimeout)
	defer cancel()

	result, err := virustotalservice.UploadFileService(xApiKey, file)
	if err != nil {
		return nil, err
	}

	err = y.yaraRepository.Save(ctx, result["data"].(map[string]interface{}))
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (y *YaraUsecase) GetResult(ctx context.Context, xApiKey string, id string) (map[string]interface{}, error) {
	_, cancel := context.WithTimeout(ctx, y.contextTimeout)
	defer cancel()

	result, err := virustotalservice.GetResultService(xApiKey, id)
	if err != nil {
		return nil, err
	}

	return result, nil
}
