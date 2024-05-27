package sharding

import (
	"github.com/DontBeProud/go-kits/error_ex"
	"strconv"
)

// RuleWithGroupConfig 基于分组的分割规则配置
type RuleWithGroupConfig struct {
	GroupSize             uint64 // 分组数量
	SplitCharacter        string // 分割字符 raw + SplitCharacter + group
	PrefixMode            bool   // 分组信息是否置于命名前缀(默认位false，即采用raw - group的格式)
	IndexIncreaseFromZero bool   // 分组序号是否从0增长. true: 0.1.2.3....  false: 1.2.3.4....
}

// GenerateFnModifyRule 生成分组相关的分割修饰规则
func (cfg *RuleWithGroupConfig) GenerateFnModifyRule() FnStringModifyRule {
	return generateFnModifyRuleWithGroup(cfg.GroupSize, cfg.calcIndexIncreaseBase())
}

func (cfg *RuleWithGroupConfig) calcIndexIncreaseBase() uint64 {
	var indexIncreaseBase uint64 = 1
	if cfg.IndexIncreaseFromZero {
		indexIncreaseBase = 0
	}
	return indexIncreaseBase
}

func (r *rootRule) generateShardingRuleWithGroup(cfg *RuleWithGroupConfig) (*shardingRuleWithGroup, error) {
	errPrefix := "rootRule.generateShardingRuleWithGroup: "
	if cfg == nil {
		return nil, error_ex.NewErrorExWithPrefix(errPrefix, "cfg == nil")
	}

	var prefixRules, suffixRules []FnStringModifyRule
	modifyRule := cfg.GenerateFnModifyRule()
	if modifyRule != nil {
		var splitRule FnStringModifyRule
		if cfg.SplitCharacter != "" {
			splitRule = func(param *StringModifyParam) string { return cfg.SplitCharacter }
		}
		if cfg.PrefixMode {
			// group-raw
			prefixRules = []FnStringModifyRule{modifyRule, splitRule}
		} else {
			// raw-group
			suffixRules = []FnStringModifyRule{splitRule, modifyRule}
		}
	}

	return &shardingRuleWithGroup{
		RootRule:          r,
		cfg:               *cfg,
		indexIncreaseBase: cfg.calcIndexIncreaseBase(),
		shardingRule:      r.generateNamingModificationRule(prefixRules, suffixRules),
	}, nil
}

// 基于分组的分割规则
type shardingRuleWithGroup struct {
	RootRule
	shardingRule      *shardingRule
	cfg               RuleWithGroupConfig
	indexIncreaseBase uint64
}

// Modify 修饰
func (r *shardingRuleWithGroup) Modify(baseString string, groupIndex uint64) string {
	return r.shardingRule.Modify(baseString, NewStringModifyParam().setGroup(groupIndex))
}

// BatchModify 批量修饰
func (r *shardingRuleWithGroup) BatchModify(baseString string, groupIndexList []uint64) []string {
	result := make([]string, len(groupIndexList))
	for index, modifier := range r.BatchGenerateStringModifiers(groupIndexList) {
		result[index] = modifier.ModifyString(baseString)
	}
	return result
}

// BatchModifyByGroupIndexRange 【返回去重结果】基于传入的分组序号范围批量修饰(start/end若为null则代表起始/末尾序号)
func (r *shardingRuleWithGroup) BatchModifyByGroupIndexRange(baseString string, start *uint64, end *uint64) []string {
	nodes := r.ExpandValidGroupIdNodeList(start, end)
	if r.cfg.GroupSize == 0 && len(nodes) > 1 {
		nodes = nodes[:1]
	}
	return r.BatchModify(baseString, nodes)
}

// GenerateStringModifier 生成字符串修饰器
func (r *shardingRuleWithGroup) GenerateStringModifier(groupIndex uint64) StringModifier {
	return r.shardingRule.GenerateStringModifier(NewStringModifyParam().setGroup(groupIndex))
}

// BatchGenerateStringModifiers 批量生成字符串修饰器
func (r *shardingRuleWithGroup) BatchGenerateStringModifiers(groupIndexList []uint64) []StringModifier {
	modifiers := make([]StringModifier, len(groupIndexList))
	for index, t := range groupIndexList {
		modifiers[index] = r.GenerateStringModifier(t)
	}
	return modifiers
}

// BatchGenerateStringModifiersByGroupIdRange 【返回去重结果】基于传入的分组序号范围批量生成字符串修饰器(start/end若为null则代表起始/末尾序号)
func (r *shardingRuleWithGroup) BatchGenerateStringModifiersByGroupIdRange(start *uint64, end *uint64) []StringModifier {
	nodes := r.ExpandValidGroupIdNodeList(start, end)
	if r.cfg.GroupSize == 0 && len(nodes) > 1 {
		nodes = nodes[:1]
	}
	return r.BatchGenerateStringModifiers(nodes)
}

// ExpandValidGroupIdNodeList 根据传入的分组id区间，展开生成有效的分组id节点列表，用于后续生成修饰器列表
func (r *shardingRuleWithGroup) ExpandValidGroupIdNodeList(start *uint64, end *uint64) []uint64 {

	_start := r.indexIncreaseBase
	if start != nil && *start > _start {
		_start = *start
	}

	_end := r.indexIncreaseBase + r.cfg.GroupSize - 1
	if end != nil && *end < _end {
		_end = *end
	}

	tList := make([]uint64, 0)
	for t := _start; t <= _end; t++ {
		tList = append(tList, t)
	}

	return tList
}

// 生成分组相关的分割修饰规则
func generateFnModifyRuleWithGroup(groupSize uint64, indexIncreaseBase uint64) FnStringModifyRule {
	if groupSize == 0 {
		return nil
	}

	// 位宽
	width := len(strconv.Itoa(int(groupSize)))

	rule := func(param *StringModifyParam) string {
		ret := ""
		if param == nil {
			return ret
		}

		t := param.getGroup()
		if t == nil {
			return ret
		}

		groupIndex := strconv.Itoa(int(indexIncreaseBase + (*t)%groupSize))
		// 空位用0填充
		fillWidth := width - len(groupIndex)
		for i := 0; i < fillWidth; i++ {
			ret += "0"
		}
		ret += groupIndex

		return ret
	}
	return rule
}

const (
	defaultModifyParamKeyGroup = "__group__"
)

func (p *StringModifyParam) setGroup(group uint64) *StringModifyParam {
	return p.Set(defaultModifyParamKeyGroup, group)
}

func (p *StringModifyParam) getGroup() *uint64 {
	_g := p.Get(defaultModifyParamKeyGroup)
	if _g == nil {
		return nil
	}
	g := _g.(uint64)
	return &g
}
