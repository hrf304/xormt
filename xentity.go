package xormt

import (
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
