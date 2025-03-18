package Model

import (
	"database/sql/driver"
	"fmt"
	"time"
)

const timeFormat = "2006-01-02 15:04:05" // 时间格式化模板
const timezone = "Asia/Shanghai"         // 时区

// Time 是自定义的时间类型，用于处理 JSON 和数据库中的时间字段
type Time time.Time

// MarshalJSON 实现 JSON 序列化接口，将时间格式化为字符串
func (t Time) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(timeFormat)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, timeFormat) // 格式化时间
	b = append(b, '"')
	return b, nil
}

// UnmarshalJSON 实现 JSON 反序列化接口，将字符串解析为时间
func (t *Time) UnmarshalJSON(data []byte) error {
	now, err := time.ParseInLocation(`"`+timeFormat+`"`, string(data), time.Local)
	if err != nil {
		return err
	}
	*t = Time(now)
	return nil
}

// String 返回时间的字符串表示
func (t Time) String() string {
	return time.Time(t).Format(timeFormat)
}

// local 将时间转换为指定时区的时间
func (t Time) local() time.Time {
	loc, _ := time.LoadLocation(timezone)
	return time.Time(t).In(loc)
}

// Value 实现 driver.Valuer 接口，用于将时间存储到数据库
func (t Time) Value() (driver.Value, error) {
	var zeroTime time.Time
	var ti = time.Time(t)
	if ti.UnixNano() == zeroTime.UnixNano() {
		return nil, nil // 如果时间为零值，返回 nil
	}
	return ti, nil
}

// Scan 实现 sql.Scanner 接口，用于从数据库读取时间
func (t *Time) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = Time(value)
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
