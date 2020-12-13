package design_patterns

// 建造者
type Builder interface {
	NewProduct()             // 建造一个产品
	SetBrand(brand string)   //贴牌子
	GetProduct() interface{} //获取这个产品
}

type Car struct {
	Brand string // 品牌
}

//汽车建造者
type CarBuilder struct {
	car *Car
}

func (cb *CarBuilder) NewProduct() {
	cb.car = &Car{}
}

func (cb *CarBuilder) SetBrand(brand string) {
	cb.car.Brand = brand
}

func (cb *CarBuilder) GetProduct() interface{} {
	return cb.car
}

type Boos struct {
	builder Builder
}

//雇佣
func (bs *Boos) Hire(b Builder) {
	bs.builder = b
}

func (b *Boos) GenCar(carBrand string) *Car {
	b.builder.NewProduct()
	b.builder.SetBrand(carBrand)
	return b.builder.GetProduct().(*Car)
}
