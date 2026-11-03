package db

type LikeString interface {
	~string
}

func MakePlaceHolder(len int) []string {
	placeholders := make([]string, len)
	for i := 0; i < len; i++ {
		placeholders[i] = "?"
	}
	return placeholders
}

func MakeColm[T ~string](rows ...T) []string {
	res := make([]string, len(rows))
	for i, row := range rows {
		res[i] = string(row)
	}
	return res
}

func ToArgs[T any](items []T) []interface{} {
	args := make([]interface{}, len(items))
	for i, v := range items {
		args[i] = v
	}
	return args
}

func MakeSet(rows []Collom) []string {
	set := make([]string, len(rows))
	for i, row := range rows {
		set[i] = string(row) + " = ?"
	}
	return set
}
