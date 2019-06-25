package xormt

/**
 * @brief: 租户数据库连接提供接口
 */
type TenantDBProvider interface{
	/**
	 * @brief: 提供接口
	 */
	Provide()[]TenantDBInfo
}