package sharding

////////////////////
// 分库分表分字段规则 //
////////////////////

func (r *rootRule) generateNamingModificationRule(prefixRuleList []FnStringModifyRule, suffixRuleList []FnStringModifyRule) *shardingRule {
	mr := NewStringModificationRule(prefixRuleList, suffixRuleList)
	mr.SetPrefixRuleList(append([]FnStringModifyRule{func(*StringModifyParam) string { return r.constPrefix }}, mr.GetPrefixRuleList()...))
	mr.SetSuffixRuleList(append(mr.GetSuffixRuleList(), func(*StringModifyParam) string { return r.constSuffix }))
	return &shardingRule{
		RootRule:         r,
		modificationRule: mr,
	}
}

// 分割规则
type shardingRule struct {
	RootRule
	modificationRule StringModifyRule
}

// ForkSubRule 分裂成子规则
func (r *shardingRule) ForkSubRule() SubRule {
	return &shardingRule{
		RootRule:         r.RootRule,
		modificationRule: NewStringModificationRule(r.modificationRule.GetPrefixRuleList(), r.modificationRule.GetSuffixRuleList()),
	}
}

// GenerateStringModifier 创建字符串修饰器
func (r *shardingRule) GenerateStringModifier(param *StringModifyParam) StringModifier {
	return NewModifier(r.modificationRule.GenPrefix(param), r.modificationRule.GenSuffix(param))
}

// BatchGenerateStringModifiers 批量创建字符串修饰器
func (r *shardingRule) BatchGenerateStringModifiers(params []*StringModifyParam) []StringModifier {
	modifiers := make([]StringModifier, len(params))
	for index, param := range params {
		modifiers[index] = r.GenerateStringModifier(param)
	}
	return modifiers
}

// Modify 修饰
func (r *shardingRule) Modify(baseString string, param *StringModifyParam) string {
	return r.GenerateStringModifier(param).ModifyString(baseString)
}

// BatchModify 批量修饰
func (r *shardingRule) BatchModify(baseString string, params []*StringModifyParam) []string {
	result := make([]string, len(params))
	for index, modifier := range r.BatchGenerateStringModifiers(params) {
		result[index] = modifier.ModifyString(baseString)
	}
	return result
}

// GetPrefixRuleList 获取前缀规则列表(不含固定的全局强制前缀规则)
func (r *shardingRule) GetPrefixRuleList() []FnStringModifyRule {
	return r.modificationRule.GetPrefixRuleList()[1:]
}

// GetSuffixRuleList 获取后缀规则列表(不含固定的全局强制后缀规则)
func (r *shardingRule) GetSuffixRuleList() []FnStringModifyRule {
	return r.modificationRule.GetSuffixRuleList()[:len(r.modificationRule.GetSuffixRuleList())-1]
}

// ResetPrefixRuleList 重置前缀规则列表(固定的全局强制前缀仍在最前)
func (r *shardingRule) ResetPrefixRuleList(newRuleList []FnStringModifyRule) {
	r.modificationRule.SetPrefixRuleList(append([]FnStringModifyRule{func(*StringModifyParam) string { return r.GetConstPrefix() }}, newRuleList...))
}

// ResetSuffixRuleList 重置后缀规则列表(固定的全局强制后缀仍在最后)
func (r *shardingRule) ResetSuffixRuleList(newRuleList []FnStringModifyRule) {
	r.modificationRule.SetSuffixRuleList(append(newRuleList, func(*StringModifyParam) string { return r.GetConstSuffix() }))
}

// PushFrontPrefixRule 前序插入新规则 至 前缀规则队列(固定的全局强制前缀仍在最前)
func (r *shardingRule) PushFrontPrefixRule(newRule FnStringModifyRule) {
	r.ResetPrefixRuleList(append([]FnStringModifyRule{newRule}, r.GetPrefixRuleList()...))
}

// PushBackPrefixRule 后序插入新规则 至 前缀规则队列
func (r *shardingRule) PushBackPrefixRule(newRule FnStringModifyRule) {
	r.ResetPrefixRuleList(append(r.GetPrefixRuleList(), newRule))
}

// PushFrontSuffixRule 前序插入新规则 至 后缀规则队列(固定的全局强制后缀仍在最后)
func (r *shardingRule) PushFrontSuffixRule(newRule FnStringModifyRule) {
	r.ResetSuffixRuleList(append([]FnStringModifyRule{newRule}, r.GetSuffixRuleList()...))
}

// PushBackSuffixRule 后序插入新规则 至 后缀规则队列
func (r *shardingRule) PushBackSuffixRule(newRule FnStringModifyRule) {
	r.ResetSuffixRuleList(append(r.GetSuffixRuleList(), newRule))
}
