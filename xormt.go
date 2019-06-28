package xormt

import (
	"github.com/go-xorm/xorm"
	"strings"
	"sync"
)

var(
 	xdbMaps map[string]*TenantDBInfo = nil
 	xdbMutex sync.Mutex

 	xSyncModels []interface{} = nil
 	xSyncModelsMutex sync.Mutex

 	xTenantIdResolver TenantIdResolver = nil
)

func init(){
	xdbMaps = make(map[string]*TenantDBInfo)
	xSyncModels = make([]interface{}, 0)
}

/**
 * @brief: 初始化
 * @param1 provider: 租户数据库提供者
 * @return: 错误信息
 */
func Init(provider TenantDBProvider, idResolver TenantIdResolver)error {
	if provider == nil {
		return &ErrParamEmpty{"provider"}
	}
	if idResolver == nil{
		xTenantIdResolver = defaultIdResolver
	}else{
		xTenantIdResolver = idResolver
	}

	tdbs := provider()
	hasDefault := false
	for i := range tdbs {
		Add(tdbs[i])
		if tdbs[i].Tid == "default" {
			hasDefault = true
		}
	}
	if !hasDefault {
		return &ErrDeaultTendarMissing{}
	}
	return nil
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

		xSyncModelsMutex.Lock()
		defer xSyncModelsMutex.Unlock()
		for i := range xSyncModels{
			syncModel(tdb, xSyncModels[i])
		}
	}

	return nil
}

/**
 * @brief: 添加同步的实体
 * @param1 model: 实体对象（例如：new(User))
 */
func AddModel(model interface{})error{
	if model == nil{
		return &ErrParamEmpty{"model"}
	}
	xSyncModelsMutex.Lock()
	defer xSyncModelsMutex.Unlock()

	xSyncModels = append(xSyncModels, model)

	xdbMutex.Lock()
	defer xdbMutex.Unlock()
	for _, v := range xdbMaps{
		syncModel(v, model)
	}

	return nil
}

/**
 * @brief: sync model
 * @param1 tenant: 租户链接
 * @param2 model: 实体对象
 * @return: 错误信息
 */
func syncModel(tenant *TenantDBInfo, model interface{})error{
	if tenant == nil{
		return &ErrParamEmpty{"tenant"}
	}
	if tenant.db == nil{
		return &ErrFieldEmpty{"tenant.db"}
	}
	if model == nil{
		return &ErrParamEmpty{"model"}
	}
	tenant.db.Sync2(model)

	return nil
}
