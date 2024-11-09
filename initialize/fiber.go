package initialize

import (
	"fiber/app/api"
	"fiber/app/service"
	"fiber/global"
	"fiber/middleware"
	"fmt"
	"github.com/gofiber/fiber/v3"
	"log"
	"os"
	"os/signal"
	"syscall"
)

// swagger handler

func InitHttpServer() {
	app := middleware.Use()
	// 初始化服务
	service.InitService()

	// 初始化路由
	api.InitRouter(app)
	global.GetRouters = app.GetRoutes()
	// 打印所有路由
	Listen := fmt.Sprintf(":%d", global.Conf.Server.Port)
	fmt.Println("Server is running on port:", Listen)

	// 启动应用
	go func() {
		if err := app.Listen(Listen, fiber.ListenConfig{EnablePrefork: global.Conf.Server.EnablePrefork}); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	go func() {
		app.Hooks().OnShutdown(func() error {

			// 关闭redis
			if err := global.Redis.Close(); err != nil {
				log.Fatalf("Failed to close Redis connection: %v", err)
			}
			// 关闭数据库
			db, _ := global.Mysql.DB()
			if err := db.Close(); err != nil {
				log.Fatalf("Failed to close MySQL connection: %v", err)
			}
			fmt.Println("资源类优雅停止")
			return nil
		})
	}()

	// 设置优雅服务关闭
	shutdown(app)
}
func shutdown(app *fiber.App) {
	// 创建一个通道来捕获中断信号
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// 等待中断信号
	<-interrupt

	// 打印关闭信息
	log.Println("Shutting down server...")

	// 关闭服务器
	if err := app.Shutdown(); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}
	//	通过钩子函数关闭数据库
	//if app.Hooks() != nil {
	//	db, _ := global.Mysql.DB()
	//	if err := db.Close(); err != nil {
	//		log.Fatalf("Failed to close MySQL connection: %v", err)
	//	}
	//	global.Redis.Shutdown(context.Background())
	//}

	// 执行清理工作
	log.Println("服务端优雅关闭")
}
