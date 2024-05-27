package sharding

import "fmt"

//////////////////////////////////////////////////
// 主要用于分库、分表等需要针对字符串统一格式化的业务场景 //
//////////////////////////////////////////////////

// FnStringModifyRule 字符串修饰规则  根据传入的参数生成字符串修饰器(前缀/后缀)
type FnStringModifyRule func(param *StringModifyParam) string

// StringModifyParam 字符串修饰参数
type StringModifyParam map[string]interface{}

// NewStringModifyParam 新建字符串修饰规则
func NewStringModifyParam() *StringModifyParam {
	p := make(StringModifyParam)
	return &p
}

func (p *StringModifyParam) Set(key string, value interface{}) *StringModifyParam {
	(*p)[key] = value
	return p
}

func (p *StringModifyParam) Get(key string) interface{} {
	val, _ := (*p)[key]
	return val
}

// StringModifier 字符串修饰器
type StringModifier interface {
	// ModifyString 修饰字符串
	ModifyString(baseString string) string
	// BatchModifyString 批量修饰字符串
	BatchModifyString(baseStrings []string) []string
}

// NewModifier 创建字符串修饰器
func NewModifier(prefix string, suffix string) StringModifier {
	return &stringModifier{
		prefix: prefix,
		suffix: suffix,
	}
}

// stringModifier 字符串修饰器
type stringModifier struct {
	prefix string
	suffix string
}

// ModifyString 修饰字符串
func (m stringModifier) ModifyString(baseString string) string {
	return fmt.Sprintf("%s%s%s", m.prefix, baseString, m.suffix)
}

// BatchModifyString 批量修饰字符串
func (m stringModifier) BatchModifyString(baseStrings []string) []string {
	_len := len(baseStrings)
	res := make([]string, _len)
	if _len > 0 {
		for index, raw := range baseStrings {
			res[index] = m.ModifyString(raw)
		}
	}
	return res
}

// StringModifyRule 字符串修饰规则
type StringModifyRule interface {
	// GenPrefix 生成字符串前缀
	GenPrefix(param *StringModifyParam) string
	// GenSuffix 根据传入的参数生成字符串后缀
	GenSuffix(param *StringModifyParam) string
	// GetPrefixRuleList 获取前缀生成规则列表
	GetPrefixRuleList() []FnStringModifyRule
	// GetSuffixRuleList 获取后缀生成规则列表
	GetSuffixRuleList() []FnStringModifyRule
	// SetPrefixRuleList 设置前缀生成规则列表
	SetPrefixRuleList(prefixRuleList []FnStringModifyRule)
	// SetSuffixRuleList 设置后缀生成规则列表
	SetSuffixRuleList(suffixRuleList []FnStringModifyRule)
	// CreateStringModifier 创建字符串修饰器
	CreateStringModifier(param *StringModifyParam) StringModifier
	// BatchCreateStringModifiers 批量创建字符串修饰器
	BatchCreateStringModifiers(params []*StringModifyParam) []StringModifier
	// ModifyString 修饰字符串
	ModifyString(baseString string, param *StringModifyParam) string
	// BatchModifyStrings 批量修饰字符串
	BatchModifyStrings(baseStrings []string, param *StringModifyParam) []string
}

// stringModifyRule 字符串修饰规则
type stringModifyRule struct {
	prefixRuleList []FnStringModifyRule // 前缀生成规则
	suffixRuleList []FnStringModifyRule // 后缀生成规则
}

// NewStringModificationRule 创建字符串修饰规则
func NewStringModificationRule(prefixRuleList []FnStringModifyRule, suffixRuleList []FnStringModifyRule) StringModifyRule {
	return &stringModifyRule{
		prefixRuleList: rebuildRuleList(prefixRuleList),
		suffixRuleList: rebuildRuleList(suffixRuleList),
	}
}

// GenPrefix 根据传入的参数生成字符串前缀
func (r *stringModifyRule) GenPrefix(param *StringModifyParam) string {
	prefix := ""
	for _, prefixRule := range r.prefixRuleList {
		prefix += prefixRule(param)
	}
	return prefix
}

// GenSuffix 根据传入的参数生成字符串后缀
func (r *stringModifyRule) GenSuffix(param *StringModifyParam) string {
	suffix := ""
	for _, suffixRule := range r.suffixRuleList {
		suffix += suffixRule(param)
	}
	return suffix
}

// GetPrefixRuleList 获取前缀生成规则列表
func (r *stringModifyRule) GetPrefixRuleList() []FnStringModifyRule {
	return r.prefixRuleList
}

// GetSuffixRuleList 获取后缀生成规则列表
func (r *stringModifyRule) GetSuffixRuleList() []FnStringModifyRule {
	return r.suffixRuleList
}

// SetPrefixRuleList 设置前缀生成规则列表
func (r *stringModifyRule) SetPrefixRuleList(prefixRuleList []FnStringModifyRule) {
	r.prefixRuleList = rebuildRuleList(prefixRuleList)
}

// SetSuffixRuleList 设置后缀生成规则列表
func (r *stringModifyRule) SetSuffixRuleList(suffixRuleList []FnStringModifyRule) {
	r.suffixRuleList = rebuildRuleList(suffixRuleList)
}

// CreateStringModifier 创建字符串修饰器
func (r *stringModifyRule) CreateStringModifier(param *StringModifyParam) StringModifier {
	return NewModifier(r.GenPrefix(param), r.GenSuffix(param))
}

// BatchCreateStringModifiers 批量创建字符串修饰器
func (r *stringModifyRule) BatchCreateStringModifiers(params []*StringModifyParam) []StringModifier {
	_len := len(params)
	res := make([]StringModifier, _len)
	if _len > 0 {
		for index, param := range params {
			res[index] = r.CreateStringModifier(param)
		}
	}
	return res
}

// ModifyString 修饰字符串
func (r *stringModifyRule) ModifyString(baseString string, param *StringModifyParam) string {
	return r.CreateStringModifier(param).ModifyString(baseString)
}

// BatchModifyStrings 批量修饰字符串
func (r *stringModifyRule) BatchModifyStrings(baseStrings []string, param *StringModifyParam) []string {
	return r.CreateStringModifier(param).BatchModifyString(baseStrings)
}

func rebuildRuleList(ruleList []FnStringModifyRule) []FnStringModifyRule {
	l := make([]FnStringModifyRule, 0)
	for _, rule := range ruleList {
		if rule != nil {
			l = append(l, rule)
		}
	}
	return l
}
