package sharding

import (
	"github.com/DontBeProud/go-kits/cron_ex"
	"time"
)

// NewRootRule 创建根命名规则
func NewRootRule(constPrefix string, constSuffix string) RootRule {
	return &rootRule{constPrefix, constSuffix}
}

// RootRule 根命名规则接口
type RootRule interface {
	RootRuleBase
	RootRuleModifier
	SubRuleGeneratorInterface
}

// RootRuleBase 根命名规则基础操作相关接口
type RootRuleBase interface {
	// GetConstPrefix 获取固定的强制前缀
	GetConstPrefix() string
	// GetConstSuffix 获取固定的强制后缀
	GetConstSuffix() string
	// ForkRootRule 分裂根命名规则
	ForkRootRule() RootRule
}

// RootRuleModifier 根命名规则修改操作相关接口
type RootRuleModifier interface {
	// SetConstPrefix 设置固定的强制前缀
	SetConstPrefix(prefix string) RootRule
	// SetConstSuffix 设置固定的强制后缀
	SetConstSuffix(suffix string) RootRule
}

// SubRuleGeneratorInterface 根命名规则生成子规则相关接口
type SubRuleGeneratorInterface interface {
	// GenerateSubShardingRule 基于 根命名规则 和 差异化的前后缀生成规则列表 生成细分的子分割规则
	GenerateSubShardingRule(prefixRuleList []FnStringModifyRule, suffixRuleList []FnStringModifyRule) SubRule
	// GenerateSubMutableShardingRule 基于 根命名规则 和 差异化的前后缀生成规则列表 生成细分的子分割规则(可变)
	GenerateSubMutableShardingRule(prefixRuleList []FnStringModifyRule, suffixRuleList []FnStringModifyRule) MutableSubRule
	// GenerateSubShardingRuleWithTime 生成基于时间的子分割规则
	GenerateSubShardingRuleWithTime(cfg *RuleWithTimeConfig) (SubRuleWithTime, error)
	// GenerateSubShardingRuleWithGroup 生成基于分组的子分割规则
	GenerateSubShardingRuleWithGroup(cfg *RuleWithGroupConfig) (SubRuleWithGroup, error)
}

// SubRule 分割规则(接口)
type SubRule interface {
	RootRuleBase
	// Modify 修饰
	Modify(baseString string, param *StringModifyParam) string
	// BatchModify 批量修饰
	BatchModify(baseString string, params []*StringModifyParam) []string
	// GenerateStringModifier 生成字符串修饰器
	GenerateStringModifier(param *StringModifyParam) StringModifier
	// BatchGenerateStringModifiers 批量生成字符串修饰器
	BatchGenerateStringModifiers(params []*StringModifyParam) []StringModifier
	// ForkSubRule 分裂成子规则
	ForkSubRule() SubRule
}

// MutableSubRule 可修改的分割规则(接口)
type MutableSubRule interface {
	SubRule
	RootRuleModifier
	// GetPrefixRuleList 获取前缀规则列表(不含固定的全局强制前缀规则)
	GetPrefixRuleList() []FnStringModifyRule
	// GetSuffixRuleList 获取后缀规则列表(不含固定的全局强制后缀规则)
	GetSuffixRuleList() []FnStringModifyRule
	// ResetPrefixRuleList 重置前缀规则列表(固定的全局强制前缀仍在最前)
	ResetPrefixRuleList([]FnStringModifyRule)
	// ResetSuffixRuleList 重置后缀规则列表(固定的全局强制后缀仍在最后)
	ResetSuffixRuleList([]FnStringModifyRule)
	// PushFrontPrefixRule 前序插入新规则 至 前缀规则队列(固定的全局强制前缀仍在最前)
	PushFrontPrefixRule(FnStringModifyRule)
	// PushBackPrefixRule 后序插入新规则 至 前缀规则队列
	PushBackPrefixRule(FnStringModifyRule)
	// PushFrontSuffixRule 前序插入新规则 至 后缀规则队列(固定的全局强制后缀仍在最后)
	PushFrontSuffixRule(FnStringModifyRule)
	// PushBackSuffixRule 后序插入新规则 至 后缀规则队列
	PushBackSuffixRule(FnStringModifyRule)
}

// TimeNodeQueryInterface 时间节点查询接口
type TimeNodeQueryInterface interface {
	// NextTimeNode 推算给定时间节点的下一个时间节点
	// round 周期累加轮数
	NextTimeNode(raw time.Time, round uint) time.Time
	// ExpandValidTimeNodeList 【返回去重结果】根据传入的时间区间，展开生成有效的时间节点列表，用于后续生成修饰器列表(起始时间会被自动修正为不早于EarliestValidTime的值)
	ExpandValidTimeNodeList(start *time.Time, end *time.Time) []time.Time
	// GetDefaultCronExpressionGenerator 根据时间分割规则等级，生成对应的默认计划任务表达式，默认秒和分钟设为0
	// 秒和分钟默认设为0
	// 日期在TimeLevelMonth及以上的等级默认设为对应时间周期内的第天
	GetDefaultCronExpressionGenerator() cron_ex.ExpressionGenerator
}

// SubRuleWithTime 基于时间的分割规则
type SubRuleWithTime interface {
	RootRuleBase
	RootRuleModifier
	TimeNodeQueryInterface
	// Modify 修饰
	Modify(baseString string, t time.Time) string
	// BatchModify 批量修饰
	BatchModify(baseString string, tList []time.Time) []string
	// BatchModifyByTimeRange 【返回去重结果】基于传入的时间范围批量修饰(起始时间会被自动修正为不早于EarliestValidTime的值)
	BatchModifyByTimeRange(baseString string, start *time.Time, end *time.Time) []string
	// GenerateStringModifier 生成字符串修饰器
	GenerateStringModifier(t time.Time) StringModifier
	// BatchGenerateStringModifiers 批量生成字符串修饰器
	BatchGenerateStringModifiers(tList []time.Time) []StringModifier
	// BatchGenerateStringModifiersByTimeRange 【返回去重结果】基于传入的时间范围批量生成字符串修饰器(起始时间会被自动修正为不早于EarliestValidTime的值)
	BatchGenerateStringModifiersByTimeRange(start *time.Time, end *time.Time) []StringModifier
}

// SubRuleWithGroup 基于分组的分割规则
type SubRuleWithGroup interface {
	RootRuleBase
	RootRuleModifier
	// Modify 修饰
	Modify(baseString string, groupIndex uint64) string
	// BatchModify 批量修饰
	BatchModify(baseString string, groupIndexList []uint64) []string
	// BatchModifyByGroupIndexRange 【返回去重结果】基于传入的分组序号范围批量修饰(start/end若为null则代表起始/末尾序号)
	BatchModifyByGroupIndexRange(baseString string, start *uint64, end *uint64) []string
	// GenerateStringModifier 生成字符串修饰器
	GenerateStringModifier(groupIndex uint64) StringModifier
	// BatchGenerateStringModifiers 批量生成字符串修饰器
	BatchGenerateStringModifiers(groupIndexList []uint64) []StringModifier
	// BatchGenerateStringModifiersByGroupIdRange 【返回去重结果】基于传入的分组序号范围批量生成字符串修饰器(start/end若为null则代表起始/末尾序号)
	BatchGenerateStringModifiersByGroupIdRange(start *uint64, end *uint64) []StringModifier
	// ExpandValidGroupIdNodeList 根据传入的分组id区间，展开生成有效的分组id节点列表，用于后续生成修饰器列表
	ExpandValidGroupIdNodeList(start *uint64, end *uint64) []uint64
}
