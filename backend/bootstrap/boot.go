package bootstrap

func Init() {

	// 加载配置文件
	InitConfig()

	// 初始化日志
	InitLog()

	// 初始化Redis
	InitRedis()

	// 初始化数据库
	InitGorm()

	// 初始化表单验证
	InitValidator()

	// 初始化Gin
	InitGin()
}
