package repository

import (
	"be-interview-task/internal/contracts"
	"be-interview-task/internal/utils/database/manager"
	"context"
	"encoding/json"
)

type YaraRepository struct {
	conn  manager.DBTx
	Table string
}

func NewYaraRepository(conn manager.DBTx) contracts.YaraRepository {
	return &YaraRepository{
		conn:  conn,
		Table: "analyzed_files",
	}
}

func (y *YaraRepository) connection(ctx context.Context) manager.DBTx {
	conn, ok := manager.DBTransactionFromContext(ctx)
	if !ok {
		conn = y.conn
	}

	return conn
}

func (y *YaraRepository) Save(ctx context.Context, data map[string]interface{}) error {
	db := y.connection(ctx)
	self := map[string]interface{}{
		"self": data["links"].(map[string]interface{})["self"],
	}

	selfJson, err := json.Marshal(self)
	if err != nil {
		return err
	}

	values := map[string]interface{}{
		"uuid":  data["id"],
		"links": selfJson,
		"type":  data["type"],
	}

	err = db.WithContext(ctx).Table(y.Table).Create(values).Error
	if err != nil {
		return err
	}

	return nil
}
