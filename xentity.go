package xormt

import (
	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"
)

/**
 * @brief 租户数据库信息
 * @field1 Name: 数据库连接名称，用于显示，可以为空
 * @field2 Tid: 租户id
 * @field3 ConnStr: 数据库连接字符串
 * @field4 DriverName: 驱动名称
 * @field5 db: 实际数据库连接对象，内部对象
 */
type TenantDBInfo struct {
	Name       string		// 数据库连接名称
	Tid        string		// 租户id
	ConnStr    string		// 连接串
	DriverName string		// 驱动名称
	db         *xorm.Engine	// 数据库连接对象
}

/**
 * @brief: 多租户上下文
 *		集成gin.Context
 */
type MultiTenantContext struct {
	*gin.Context
	DB *xorm.Engine
}

/**
 * @brief: 多租户处理函数
 */
type MultiTenantHandlerFunc func(*MultiTenantContext)

/**
 * @brief: 租户数据库提供者
 */
type TenantDBProvider func()[]*TenantDBInfo

/**
 * @brief: 住户id解析器
 */
type TenantIdResolver func(*gin.Context)string
