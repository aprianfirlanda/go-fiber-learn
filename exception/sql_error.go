package exception

type SqlError struct {
	Message string
}

func (sqlError *SqlError) Error() string {
	return sqlError.Message
}
