package design_patterns

import "testing"

func TestFactory(t *testing.T) {

	var factoryInf FactoryInterfaca
	factoryInf = new(wuhanFactory)
	factoryInf.GenPhone("xiaomi").work()
	factoryInf.GenPhone("huawei").work()




	factoryInf = new(shanghaiFactory)
	factoryInf.GenPhone("xiaomi").work()
	factoryInf.GenPhone("huawei").work()

}