package xormt

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

func HandlerGin(handler MultiTenantHandlerFunc) gin.HandlerFunc{
	return func(c *gin.Context){
		mc := new(MultiTenantContext)
		mc.Context = c
		tid := xTenantIdResolver(c)
		if tdb, exist := xdbMaps[tid]; exist{
			mc.DB = tdb.db
		}else{
			fmt.Println("can not get tdb by tid", tid)
		}
		handler(mc)
	}
}

func defaultIdResolver(ctx *gin.Context)string{
	id, _ := ctx.GetQuery("tenant")
	if strings.TrimSpace(id) == ""{
		id = ctx.GetString("tenant")
	}

	return id
}
