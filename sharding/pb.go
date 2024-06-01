package sharding

import (
	"github.com/DontBeProud/go-kits/sharding/sharding_pb"
)

// NewRootRuleWithPb 基于pb创建根命名规则
func NewRootRuleWithPb(cfg *sharding_pb.ShardingRootConfig) RootRule {
	if cfg == nil {
		return &rootRule{}
	}
	return &rootRule{
		constPrefix: cfg.GetConstPrefix(),
		constSuffix: cfg.GetConstSuffix(),
	}
}

// GenerateSubShardingRuleWithTimeWithPb 生成基于时间的子分割规则(通过pb)
func (r *rootRule) GenerateSubShardingRuleWithTimeWithPb(cfg *sharding_pb.ShardingWithTimeConfig) (SubRuleWithTime, error) {
	return r.GenerateSubShardingRuleWithTime(ParsesShardingWithTimeConfigPb(cfg))
}

// GenerateSubShardingRuleWithGroupWithPb 生成基于分组的子分割规则(通过pb)
func (r *rootRule) GenerateSubShardingRuleWithGroupWithPb(cfg *sharding_pb.RuleWithGroupConfig) (SubRuleWithGroup, error) {
	return r.GenerateSubShardingRuleWithGroup(ParseRuleWithGroupConfigWithPb(cfg))
}

func ParseRuleWithGroupConfigWithPb(cfg *sharding_pb.RuleWithGroupConfig) *RuleWithGroupConfig {
	if cfg == nil {
		return nil
	}

	return &RuleWithGroupConfig{
		GroupSize:             cfg.GetGroupSize(),
		SplitCharacter:        cfg.GetSplitCharacter(),
		PrefixMode:            cfg.GetPrefixMode(),
		IndexIncreaseFromZero: cfg.GetIndexIncreaseFromZero(),
	}
}

// ParsesShardingWithTimeConfigPb 解析ShardingWithTimeConfigPb
func ParsesShardingWithTimeConfigPb(cfg *sharding_pb.ShardingWithTimeConfig) *RuleWithTimeConfig {
	if cfg == nil {
		return nil
	}

	res := &RuleWithTimeConfig{
		Level:              TimeLevel(cfg.GetTimeLevel().Number()),
		SplitCharacter:     cfg.GetSplitCharacter(),
		TimeSplitCharacter: cfg.GetTimeSplitCharacter(),
		PrefixMode:         cfg.GetPrefixMode(),
	}

	if t := cfg.GetEarliestValidTime(); t != nil {
		res.EarliestValidTime = cfg.GetEarliestValidTime().AsTime()
	}

	return res
}
