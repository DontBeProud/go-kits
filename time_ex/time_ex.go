package time_ex

import (
	"database/sql/driver"
	"fmt"
	"time"
)

// TODO: 支持用户自定义格式

// TimeEx json默认序列化为"2006-01-02 15:04:05"格式
type TimeEx struct {
	time.Time
}

func NewTimeX(t time.Time) *TimeEx {
	return &TimeEx{t}
}

// MarshalJSON 预设序列化规则
func (t TimeEx) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", t.Format(time.DateTime))), nil
}

// Value 用于兼容gorm的Value接口, (t TimeEx) 勿改为 (t *TimeEx)
func (t TimeEx) Value() (driver.Value, error) {
	if t.Time.IsZero() {
		return nil, nil
	}
	return t.Time, nil
}

// Scan 用于兼容gorm的Scan接口, (t *TimeEx) 勿改为 (t TimeEx)
func (t *TimeEx) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = TimeEx{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
