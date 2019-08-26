package main

// 由于go module的前缀是github.com/zhiwen-kooboo/GolangDemo
// 所以import是这样的目录结构
import "github.com/zhiwen-kooboo/GolangDemo/go_project_demo/test_package"

func main() {
	test_package.Hello()
}
