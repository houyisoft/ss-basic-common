package utils

import (
	"ss-basic-common/utils/cast"
	"strings"

)


type Operator int

const (
	EQ Operator = iota
	STRING_EQ
	GE
	LE
	DATE_GE
	DATE_LE
	TIME_GE
	TIME_LE
	LIKE
	LIKE_LEFT
	LIKE_RIGHT
	UNEQUAL
)

func (o Operator) string() string {
	switch o {
	case EQ:
		return "="
	case STRING_EQ:
		return "="
	case GE:
		return ">="
	case LE:
		return "<="
	case TIME_GE:
		return ">="
	case TIME_LE:
		return "<="
	case DATE_GE:
		return ">="
	case DATE_LE:
		return "<="
	case LIKE:
		return "%"
	case LIKE_LEFT:
		return "%"
	case LIKE_RIGHT:
		return "%"
	case UNEQUAL:
		return "!="
	default:
		return "="
	}

}

func GetOperator(oper string) Operator {
	switch oper {
	case "EQ":
		return EQ
	case "STRING_EQ":
		return STRING_EQ
	case "GE":
		return GE
	case "LE":
		return LE
	case "TIME_GE":
		return TIME_GE
	case "TIME_LE":
		return TIME_LE
	case "DATE_GE":
		return DATE_GE
	case "DATE_LE":
		return DATE_LE
	case "LIKE":
		return LIKE
	case "LIKE_LEFT":
		return LIKE_LEFT
	case "LIKE_RIGHT":
		return LIKE_RIGHT
	case "UNEQUAL":
		return UNEQUAL
	default:
		return EQ
	}
}


type Filter struct {
	Name     string
	Operator Operator
	Val      string
}


func GeneratorConditionSql(filters []Filter) string {
	if len(filters) == 0 {
		return ""
	}
	whereSql := ""
	for _, v := range filters {
		operator := v.Operator
		val := v.Val
		if operator == EQ {
			whereSql += " AND " + v.Name + "=" + val
		} else if operator == STRING_EQ {
			whereSql += " AND " + v.Name + "='" + val + "'"
		} else if operator == GE {
			whereSql += " AND " + v.Name + "<=" + val
		} else if operator == LE {
			whereSql += " AND " + v.Name + ">=" + val
		} else if operator == DATE_GE {
			val = val + " 00:00:00"
			iFromTimestamp := GetInt64FromTime(val)
			whereSql += " AND " + v.Name + "<=" + cast.ToString(iFromTimestamp)
		} else if operator == DATE_LE {
			val = val + " 00:00:00"
			iFromTimestamp := GetInt64FromTime(val)
			whereSql += " AND " + v.Name + ">=" + cast.ToString(iFromTimestamp)
		} else if operator == TIME_GE {
			iFromTimestamp := GetInt64FromTime(val)
			whereSql += " AND " + v.Name + "<=" + cast.ToString(iFromTimestamp)
		} else if operator == TIME_LE {
			iFromTimestamp := GetInt64FromTime(val)
			whereSql += " AND " + v.Name + ">=" + cast.ToString(iFromTimestamp)
		} else if operator == LIKE {
			whereSql += " AND " + v.Name + " like '%" + val + "%'"
		} else if operator == LIKE_LEFT {
			whereSql += " AND " + v.Name + " like '%" + val
		} else if operator == LIKE_RIGHT {
			whereSql += " AND " + v.Name + " like '" + val + "%'"
		}
	}
	if whereSql != "" {
		replace := strings.Replace(whereSql, " AND", "", 1)
		whereSql = " WHERE " + replace
	}
	return whereSql
}



