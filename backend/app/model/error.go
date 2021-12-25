package model

const (
	/*
		poetry project error
	*/

	/*
		database error code ( -2000 ~ -2009 )
	*/
	ERR_DATABASE_CREATE = -2000
	ERR_DATABASE_READ   = -2001
	ERR_DATABASE_UPDATE = -2002
	ERR_DATABASE_DELETE = -2003
	ERR_DATABASE_COUNT  = -2004
	/*
		http error code ( -2010 ~  )
	*/
	ERR_INVALID_PARAM = -2010
)

var ERR_MAP = map[int]string{
	ERR_DATABASE_CREATE: "database create error",
	ERR_DATABASE_READ:   "database read error",
	ERR_DATABASE_UPDATE: "database update error",
	ERR_DATABASE_DELETE: "database delete error",
	ERR_DATABASE_COUNT:  "database count error",
	ERR_INVALID_PARAM:   "invalid param",
}
