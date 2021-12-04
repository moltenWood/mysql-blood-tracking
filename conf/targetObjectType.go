package conf

type TargetObjectType int32

const (
	TableName TargetObjectType = 0
	FieldName TargetObjectType = 1
)

func (self TargetObjectType) String() string {
	switch self {
	case TableName:
		return "表名"
	case FieldName:
		return "某个表的字段名"
	default:
		return "UNKNOWN"
	}
}
