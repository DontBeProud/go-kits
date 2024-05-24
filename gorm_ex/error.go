package gorm_ex

import (
	"errors"
	"fmt"
	"github.com/go-sql-driver/mysql"
)

// IsErrorDuplicateKey 是否是键值重复的错误
func IsErrorDuplicateKey(err error) bool {
	if err == nil {
		return false
	}

	var _err *mysql.MySQLError
	return errors.As(err, &_err) && _err != nil && _err.Number == 1062 && fmt.Sprintf("%s", _err.SQLState) == "23000"
}
