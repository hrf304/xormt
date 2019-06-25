package xormt

import (
		"sync"
	"strings"
	"github.com/go-xorm/xorm"
)

var(
 	xdbMaps map[string]*TenantDBInfo
 	xdbMutex sync.Mutex
)

func init(){
	xdbMaps = make(map[string]*TenantDBInfo)
}

/**
 * @brief: 添加db
 * @param1 tdb: 租户数据库信息
 * @return1 错误信息
 */
func Add(tdb *TenantDBInfo)error{
	if strings.TrimSpace(tdb.Tid) == ""{
		return &ErrFieldEmpty{"Tid",}
	} else if strings.TrimSpace(tdb.ConnStr) == ""{
		return &ErrFieldEmpty{"ConnStr"}
	} else if strings.TrimSpace(tdb.DriverName) == ""{
		return &ErrFieldEmpty{"DriverName"}
	}

	var err error
	tdb.db, err = xorm.NewEngine(tdb.DriverName, tdb.ConnStr)
	if err != nil{
		return err
	}

	xdbMutex.Lock()
	defer xdbMutex.Unlock()

	if _, exist := xdbMaps[tdb.Tid]; !exist{
		xdbMaps[tdb.Tid] = tdb
	}

	return nil
}
