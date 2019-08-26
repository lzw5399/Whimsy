package main

// 这个路径是gopath之下src下面的路径
// 这是单一gopath的好处
import "github.com/zhiwen-kooboo/GolangDemo/go_project_demo/test_package"

func main() {
	test_package.Hello()
}
