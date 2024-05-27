package sharding

// rootRule 根命名规则
type rootRule struct {
	constPrefix string // 固定的强制前缀
	constSuffix string // 固定的强制后缀
}

// GetConstPrefix 获取固定的强制前缀
func (r *rootRule) GetConstPrefix() string {
	return r.constPrefix
}

// GetConstSuffix 获取固定的强制后缀
func (r *rootRule) GetConstSuffix() string {
	return r.constSuffix
}

// SetConstPrefix 设置固定的强制前缀
func (r *rootRule) SetConstPrefix(prefix string) RootRule {
	r.constPrefix = prefix
	return r
}

// SetConstSuffix 设置固定的强制后缀
func (r *rootRule) SetConstSuffix(suffix string) RootRule {
	r.constSuffix = suffix
	return r
}

// ForkRootRule 分裂根命名规则
func (r *rootRule) ForkRootRule() RootRule {
	return NewRootRule(r.GetConstPrefix(), r.GetConstSuffix())
}

// GenerateSubShardingRule 基于 根命名规则 和 差异化的前后缀生成规则列表 生成细分的分割规则
func (r *rootRule) GenerateSubShardingRule(prefixRuleList []FnStringModifyRule, suffixRuleList []FnStringModifyRule) SubRule {
	return r.generateNamingModificationRule(prefixRuleList, suffixRuleList)
}

// GenerateSubMutableShardingRule 基于 根命名规则 和 差异化的前后缀生成规则列表 生成细分的分割规则(可变)
func (r *rootRule) GenerateSubMutableShardingRule(prefixRuleList []FnStringModifyRule, suffixRuleList []FnStringModifyRule) MutableSubRule {
	return r.generateNamingModificationRule(prefixRuleList, suffixRuleList)
}

// GenerateSubShardingRuleWithTime 生成基于时间的分割规则
func (r *rootRule) GenerateSubShardingRuleWithTime(cfg *RuleWithTimeConfig) (SubRuleWithTime, error) {
	return r.generateShardingRuleWithTime(cfg)
}

// GenerateSubShardingRuleWithGroup 生成基于分组的子分割规则
func (r *rootRule) GenerateSubShardingRuleWithGroup(cfg *RuleWithGroupConfig) (SubRuleWithGroup, error) {
	return r.generateShardingRuleWithGroup(cfg)
}
