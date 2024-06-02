package gorm_ex

import (
	"context"
	"github.com/DontBeProud/go-kits/error_ex"
	"gorm.io/gorm"
)

// TableInitMethod 表初始化方法
type TableInitMethod func(ctx context.Context, db *gorm.DB, tableName string) error

// BatchInitTable 批量初始化表
func BatchInitTable(ctx context.Context, eg *error_ex.ErrorGroupEx, db *gorm.DB, initMethod TableInitMethod,
	tableNames []string) error {

	// init execute method list
	tableNameNum := len(tableNames)
	execMethods := make([]ExecuteMethod, tableNameNum)
	for index, tableName := range tableNames {
		_tableName := tableName
		execMethods[index] = func(_ctx context.Context, _db *gorm.DB) error {
			return initMethod(_ctx, _db, _tableName)
		}
	}

	return BatchExecuteSqlCommands(ctx, eg, db, execMethods)
}
