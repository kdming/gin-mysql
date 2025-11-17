package sqlUtil

type sqlUtil struct {
	sql    string
	params []interface{}
}

func NewSqlUtil() *sqlUtil {
	return &sqlUtil{
		sql:    "",
		params: []interface{}{},
	}
}

func (s *sqlUtil) AddParams(filed string, val interface{}) {
	if s.sql != "" {
		s.sql += " AND " + filed + " = ?"
	} else {
		s.sql += filed + " = ?"
	}
	s.params = append(s.params, val)
}

func (s *sqlUtil) AddParamsAndOp(filed, op string, val interface{}) {
	if s.sql != "" {
		s.sql += " AND " + filed + " " + op + " ?"
	} else {
		s.sql += filed + " " + op + " ?"
	}
	s.params = append(s.params, val)
}

func (s *sqlUtil) Build() string {
	return s.sql
}

func (s *sqlUtil) Params() []interface{} {
	return s.params
}
