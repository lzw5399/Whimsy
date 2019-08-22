package package_for_test

type Cat struct {
	Name  string
	Color string
}

// 模拟一个【子类】
// Note: 这里的【父struct】不需要给【成员名】，然后就可以支持以下【两种】写法
// 1、【blackCat.Name】 2、【blackCat.Cat.Name】
// 如果给了【成员名】，则只可以用这种写法
// 1、【blackCat.成员名.Name】
type BlackCat struct {
	Cat // 不写成员名，这种用法叫【嵌入】
	Hobby string
}

// 由于go不支持重载，所以这边构造函数的名字不能一样啊
func NewCatByName(name string) *Cat {
	return &Cat{Name: name}
}

// 父struct构造函数
func NewCat(name string, color string) *Cat {
	return &Cat{Name: name, Color: color}
}

// 子struct构造函数
func NewBlackCat(color string) *BlackCat {
	// cat := new(BlackCat) 或者用new也一样
	cat := &BlackCat{
		Cat:   Cat{
			Name:  "dsds",
			Color: color,
		},
		Hobby: "ggg",
	}
	return cat
}
