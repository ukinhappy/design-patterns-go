package design_patterns

import "fmt"

//工厂接口
type FactoryInterfaca interface {
	GenPhone(name string) Phone
}

//实体work
type Phone interface {
	work()
}

// -----
//工厂实例1
type wuhanFactory struct {
}

func (w *wuhanFactory) GenPhone(name string) Phone {
	switch name {
	case "xiaomi":
		return wuhanxiaomo{}
	case "huawei":
		return wuhanhuawei{}
	}
	return nil
}

type shanghaiFactory struct {
}

func (w *shanghaiFactory) GenPhone(name string) Phone {
	switch name {
	case "xiaomi":
		return shanghaixiaomo{}
	case "huawei":
		return shanghaihuawei{}
	}
	return nil
}

// --- 工厂实例---

type wuhanxiaomo struct {
}

func (x wuhanxiaomo) work() {
	fmt.Println("wuhanxiaomo")
}

type wuhanhuawei struct {
}

func (x wuhanhuawei) work() {
	fmt.Println("wuhanhuawei")
}

type shanghaixiaomo struct {
}

func (x shanghaixiaomo) work() {
	fmt.Println("shanghaixiaomo")
}

type shanghaihuawei struct {
}

func (x shanghaihuawei) work() {
	fmt.Println("shanghaihuawei")
}
