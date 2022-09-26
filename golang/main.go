package main
 
import (
    "github.com/astaxie/beego"  //导入beego依赖
)
 
type MainController struct {
    beego.Controller    //匿名包含beego.Controller
}
 
func (this *MainController) Get() {
    this.Ctx.WriteString("hello world")  //实现get响应
}
 //https://cloud.tencent.com/developer/article/1541826
func main() {
    beego.Router("/", &MainController{})  //设置路由
    beego.Run()   //启动Beego       
}