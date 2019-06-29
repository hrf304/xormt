# xormt
基于xorm和gin多租户实现

## 实例仓库
https://github.com/hrf304/ukid.git

## 使用步骤
1、在main函数中调用Init进行初始化，并传入相关TenantDBProvider和TenantIdResolver的具体实现

```
func main() {
	ginEngine := gin.Default()
	xormt.Init(xormtext.GetTenants, xormtext.GetTenantId)
	router.Register(ginEngine)
	ginEngine.Run(":8080")
}
```
2、需要同步实体类中调用AddModel

```
func init(){
	xormt.AddModel(new (User))
}
```
3、映射router，此时需要使用xormt.HandlerGin返回gin.HandlerFunc

```
v1.GET("/users/:id", xormt.HandlerGin(ctrl.Get))
```
4、实现controller处理函数，参数为*xormt.MultiTenantContext

```
func (c *UserController) Get(ctx *xormt.MultiTenantContext) {
	i++

	user := &entity.User{}
	user.Id = util.UUID()
	user.Name = fmt.Sprintf("huangrf%d", i)
	user.LoginId = fmt.Sprintf("huangrf%d", i)
	user.Major = fmt.Sprintf("major%d", i)

	_, err := ctx.DB.InsertOne(user)
	if err != nil{
		ctx.JSON(500, &entity.Resp{500, err.Error(), nil})
	}else{
		ctx.JSON(http.StatusOK, &entity.Resp{http.StatusOK, "", user.Id})
	}
}
```
