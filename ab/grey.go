package ab

type Condition struct {
	Key      string
	Value    interface{}
	Operator string
}

type GreyConfig struct {
	WhiteList  []Condition
	Conditions []Condition
	RateRatio  int // 千分比
}

type TestTool struct {
	Properties map[string]string
	Pid        string
}

type Match interface {
	IsMatch(value string, dstValue interface{}) bool
}

type EqualMatch struct {
}

func (m *EqualMatch) IsMatch(value string, dstValue interface{}) bool {
	if _, ok := dstValue.(string); !ok {
		return false
	}
	return value == dstValue.(string)
}

func (c Condition) Hit(value string) bool {
	var match Match
	if c.Operator == "=" {
		match = &EqualMatch{}
	} else {
		return false
	}
	return match.IsMatch(value, c.Value)
}

func (t *TestTool) HitToggle(cfg GreyConfig) {

}
