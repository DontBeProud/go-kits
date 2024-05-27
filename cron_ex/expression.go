package cron_ex

import (
	"fmt"
	"strconv"
)

// NewExpressionGenerator 创建cron表达式生成器接口, 默认 * * * * * ?(每一秒)
func NewExpressionGenerator() ExpressionGenerator {
	return &specGenerator{}
}

// TimeRange 表达式时间范围
type TimeRange struct {
	Start *uint32
	End   *uint32
}

func NewTimeRange() *TimeRange {
	return &TimeRange{}
}

// ExpressionGenerator cron表达式生成器接口
// 封装这个生成器纯粹为了开发过程中可以偷懒, 内部不负责有效性验证, 可自行至https://www.pppet.net/等网站验证spec的有效性
type ExpressionGenerator interface {
	// GenExpression 生成spec
	GenExpression() string

	GetSecond() string
	GetMinute() string
	GetHour() string
	GetDay() string
	GetMonth() string
	GetWeek() string

	SetSecondAny() ExpressionGenerator
	SetMinuteAny() ExpressionGenerator
	SetHourAny() ExpressionGenerator
	SetDayAny() ExpressionGenerator
	SetMonthAny() ExpressionGenerator
	SetWeekAny() ExpressionGenerator

	SetWorkDay(val uint32) ExpressionGenerator

	SetSecondConstValue(val uint32) ExpressionGenerator
	SetMinuteConstValue(val uint32) ExpressionGenerator
	SetHourConstValue(val uint32) ExpressionGenerator
	SetDayConstValue(val uint32) ExpressionGenerator
	SetMonthConstValue(val uint32) ExpressionGenerator
	SetWeekConstValue(val uint32) ExpressionGenerator

	SetSecondsConstValueList(l []uint32) ExpressionGenerator
	SetMinutesConstValueList(l []uint32) ExpressionGenerator
	SetHoursConstValueList(l []uint32) ExpressionGenerator
	SetDaysConstValueList(l []uint32) ExpressionGenerator
	SetMonthsConstValueList(l []uint32) ExpressionGenerator
	SetWeeksConstValueList(l []uint32) ExpressionGenerator

	SetSecondRange(tRange TimeRange, secInterval *uint32) ExpressionGenerator
	SetMinuteRange(tRange TimeRange, minInterval *uint32) ExpressionGenerator
	SetHourRange(tRange TimeRange, hourInterval *uint32) ExpressionGenerator
	SetDayRange(tRange TimeRange, dayInterval *uint32) ExpressionGenerator
	SetMonthRange(tRange TimeRange, monthInterval *uint32) ExpressionGenerator
	SetWeekRange(tRange TimeRange, dayInterval *uint32) ExpressionGenerator
}

// cron spec 生成器
type specGenerator struct {
	second *string
	minute *string
	hour   *string
	day    *string
	month  *string
	week   *string
}

// GenExpression 生成spec
func (g *specGenerator) GenExpression() string {
	return fmt.Sprintf("%s %s %s %s %s %s", g.GetSecond(), g.GetMinute(), g.GetHour(), g.GetDay(), g.GetMonth(), g.GetWeek())
}

func (g *specGenerator) SetWorkDay(val uint32) ExpressionGenerator {
	s := fmt.Sprintf("%dW", val)
	g.day = &s
	return g
}

func (g *specGenerator) SetSecondRange(tRange TimeRange, secInterval *uint32) ExpressionGenerator {
	g.second = _parseRange(tRange, secInterval)
	return g
}

func (g *specGenerator) SetMinuteRange(tRange TimeRange, minInterval *uint32) ExpressionGenerator {
	g.minute = _parseRange(tRange, minInterval)
	return g
}

func (g *specGenerator) SetHourRange(tRange TimeRange, hourInterval *uint32) ExpressionGenerator {
	g.hour = _parseRange(tRange, hourInterval)
	return g
}

func (g *specGenerator) SetDayRange(tRange TimeRange, dayInterval *uint32) ExpressionGenerator {
	g.day = _parseRange(tRange, dayInterval)
	return g
}

func (g *specGenerator) SetMonthRange(tRange TimeRange, monthInterval *uint32) ExpressionGenerator {
	g.month = _parseRange(tRange, monthInterval)
	return g
}

func (g *specGenerator) SetWeekRange(tRange TimeRange, dayInterval *uint32) ExpressionGenerator {
	g.week = _parseRange(tRange, dayInterval)
	return g
}

func (g *specGenerator) SetSecondAny() ExpressionGenerator {
	g.second = nil
	return g
}

func (g *specGenerator) SetMinuteAny() ExpressionGenerator {
	g.minute = nil
	return g
}

func (g *specGenerator) SetHourAny() ExpressionGenerator {
	g.hour = nil
	return g
}

func (g *specGenerator) SetDayAny() ExpressionGenerator {
	g.day = nil
	return g
}

func (g *specGenerator) SetMonthAny() ExpressionGenerator {
	g.month = nil
	return g
}

func (g *specGenerator) SetWeekAny() ExpressionGenerator {
	g.week = nil
	return g
}

func (g *specGenerator) SetSecondConstValue(val uint32) ExpressionGenerator {
	s := strconv.Itoa(int(val))
	g.second = &s
	return g
}

func (g *specGenerator) SetMinuteConstValue(val uint32) ExpressionGenerator {
	s := strconv.Itoa(int(val))
	g.minute = &s
	return g
}

func (g *specGenerator) SetHourConstValue(val uint32) ExpressionGenerator {
	s := strconv.Itoa(int(val))
	g.hour = &s
	return g
}

func (g *specGenerator) SetDayConstValue(val uint32) ExpressionGenerator {
	s := strconv.Itoa(int(val))
	g.day = &s
	return g
}

func (g *specGenerator) SetMonthConstValue(val uint32) ExpressionGenerator {
	s := strconv.Itoa(int(val))
	g.month = &s
	return g
}

func (g *specGenerator) SetWeekConstValue(val uint32) ExpressionGenerator {
	s := strconv.Itoa(int(val))
	g.week = &s
	return g
}

func (g *specGenerator) SetSecondsConstValueList(l []uint32) ExpressionGenerator {
	g.second = _parseConstList(l)
	return g
}

func (g *specGenerator) SetMinutesConstValueList(l []uint32) ExpressionGenerator {
	g.minute = _parseConstList(l)
	return g
}

func (g *specGenerator) SetHoursConstValueList(l []uint32) ExpressionGenerator {
	g.hour = _parseConstList(l)
	return g
}

func (g *specGenerator) SetDaysConstValueList(l []uint32) ExpressionGenerator {
	g.day = _parseConstList(l)
	return g
}

func (g *specGenerator) SetMonthsConstValueList(l []uint32) ExpressionGenerator {
	g.month = _parseConstList(l)
	return g
}

func (g *specGenerator) SetWeeksConstValueList(l []uint32) ExpressionGenerator {
	g.week = _parseConstList(l)
	return g
}

func (g *specGenerator) GetSecond() string {
	if g.second == nil {
		return "*"
	}
	return *g.second
}

func (g *specGenerator) GetMinute() string {
	if g.minute == nil {
		return "*"
	}
	return *g.minute
}

func (g *specGenerator) GetHour() string {
	if g.hour == nil {
		return "*"
	}
	return *g.hour
}

func (g *specGenerator) GetDay() string {
	if g.day == nil {
		return "*"
	}
	return *g.day
}

func (g *specGenerator) GetMonth() string {
	if g.month == nil {
		return "*"
	}
	return *g.month
}

func (g *specGenerator) GetWeek() string {
	if g.week == nil {
		return "?"
	}
	return *g.week
}

func _parseConstList(l []uint32) *string {
	s := ""
	for index, item := range l {
		if index > 0 {
			s += ","
		}
		s += strconv.Itoa(int(item))
	}
	return &s
}

func _parseRange(tRange TimeRange, interval *uint32) *string {
	s := tRange.parse()
	if interval != nil {
		s = fmt.Sprintf("%s/%d", s, *interval)
	}
	return &s
}

func (r *TimeRange) parse() string {
	s := "*"
	if r.Start != nil {
		s = strconv.Itoa(int(*r.Start))
		if r.End != nil {
			s = fmt.Sprintf("%s-%d", s, *r.End)
		}
	}
	return s
}

func (r *TimeRange) SetStart(s uint32) *TimeRange {
	_s := s
	r.Start = &_s
	return r
}

func (r *TimeRange) SetEnd(e uint32) *TimeRange {
	_e := e
	r.End = &_e
	return r
}

func (r *TimeRange) UnSetStart() *TimeRange {
	r.Start = nil
	return r
}

func (r *TimeRange) UnSetEnd() *TimeRange {
	r.End = nil
	return r
}
