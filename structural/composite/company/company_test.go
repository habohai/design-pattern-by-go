package company

import (
	"fmt"
	"testing"
)

func TestCompany(t *testing.T) {
	root := NewConcreteCompany("北京总公司")
	root.Add("总公司人力资源部", NewHRDepartment("总公司人力资源部"))
	root.Add("总公司财务", NewFinanceDepartment("总公司财务"))

	com := NewConcreteCompany("上海华东分公司")
	com.Add("上海华东分公司人力资源部", NewHRDepartment("上海华东分公司人力资源部"))
	com.Add("上海华东分公司财务", NewFinanceDepartment("上海华东分公司财务"))
	root.Add("上海华东分公司", com)

	com1 := NewConcreteCompany("南京办事处")
	com1.Add("南京办事处人力资源部", NewHRDepartment("南京办事处人力资源部"))
	com1.Add("南京办事处财务", NewFinanceDepartment("南京办事处财务"))
	com.Add("南京办事处", com1)

	com2 := NewConcreteCompany("杭州办事处")
	com2.Add("杭州办事处人力资源部", NewHRDepartment("杭州办事处人力资源部"))
	com2.Add("杭州办事处财务", NewFinanceDepartment("杭州办事处财务"))
	com.Add("杭州办事处", com2)

	fmt.Println("结构图:")
	root.Display(1)
	fmt.Println("职责:")
	root.LineOfDuty()
}
