package cmds

/**
 * Author  :  xiean
 * EMail   :  xiepaup@163.com 
 * GitHub  :  https://github.com/xiepaup
 * Date    :  2018-08-05 11:53
 * Project :  dbatools
 **/

type RedisCMD struct {
}

const (
	REDIS_DATA_TYPE_HASH_NAME = "hash"
	REDIS_DATA_TYPE_STRINGS_NAME = "string"
	REDIS_DATA_TYPE_LIST_NAME = "list"
	REDIS_DATA_TYPE_SET_NAME  = "set"
	REDIS_DATA_TYPE_ZSET_NAME = "zset"
	REDIS_DATA_TYPE_OTHER_NAME = "unkonwn"
)

const (
	REDIS_DATA_TYPE_HAHS_CODE = iota
	REDIS_DATA_TYPE_STRINGS_CODE
	REDIS_DATA_TYPE_LIST_CODE
	REDIS_DATA_TYPE_SET_CODE
	REDIS_DATA_TYPE_ZSET_CODE
	REDIS_DATA_TYPE_OTHER_CODE
)


const (
	UNKOWN_CMD_TYPE     = 8888

	HASH_CMD_TYPE_READ  = 9101
	HASH_CMD_TYPE_WRITE = 9102

	STR_CMD_TYPE_READ   = 9201
	STR_CMD_TYPE_WRITE  = 9202

	LIST_CMD_TYPE_READ  = 9301
	LIST_CMD_TYPE_WRITE = 9302

	ZSET_CMD_TYPE_READ  = 9401
	ZSET_CMD_YTPE_WRITE = 9402

	SET_CMD_TYPE_READ  = 9501
	SET_CMD_TYPE_WRITE = 9502
)



func  GetDataTypeByCmd(c string)  {

}