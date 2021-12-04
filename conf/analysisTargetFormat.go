package conf

type AnalysisTargetFormat int32

const (
	AnalysisTargetFormat0 AnalysisTargetFormat = 0
	AnalysisTargetFormat1 AnalysisTargetFormat = 1
	AnalysisTargetFormat2 AnalysisTargetFormat = 2
	AnalysisTargetFormat3 AnalysisTargetFormat = 3
)

func (self AnalysisTargetFormat) String() string {
	switch self {
	case AnalysisTargetFormat0:
		return "JOIN tableName ON node = node"
	case AnalysisTargetFormat1:
		return "LEFT JOIN tableName ON node = node"
	case AnalysisTargetFormat2:
		return "RIGHT JOIN tableName ON node = node"
	case AnalysisTargetFormat3:
		return "UNKNOWN"
	default:
		return "UNKNOWN"
	}
}

func (self AnalysisTargetFormat) GetKeyWords() []string {
	switch self {
	case AnalysisTargetFormat0:
		return []string{"JOIN", "ON"}
	case AnalysisTargetFormat1:
		return []string{"LEFT JOIN", "ON"}
	case AnalysisTargetFormat2:
		return []string{"RIGHT JOIN", "ON"}
	case AnalysisTargetFormat3:
		return []string{}
	default:
		return []string{}
	}
}

func (self AnalysisTargetFormat) GetFirstKeyWords() string {
	switch self {
	case AnalysisTargetFormat0:
		return "JOIN"
	case AnalysisTargetFormat1:
		return "LEFT JOIN"
	case AnalysisTargetFormat2:
		return "RIGHT JOIN"
	case AnalysisTargetFormat3:
		return ""
	default:
		return ""
	}
}

func (self AnalysisTargetFormat) GetLastKeyWords() string {
	switch self {
	case AnalysisTargetFormat0:
		return "ON"
	case AnalysisTargetFormat1:
		return "ON"
	case AnalysisTargetFormat2:
		return "ON"
	case AnalysisTargetFormat3:
		return ""
	default:
		return ""
	}
}
