package company

import (
	"fmt"
	"strings"
)

type HRDepartment struct {
	Name string
}

func NewHRDepartment(name string) *HRDepartment {
	return &HRDepartment{
		Name: name,
	}
}

func (hrd *HRDepartment) Display(depth int) {
	fmt.Println(strings.Repeat("-", depth), " ", hrd.Name)
}

func (hrd *HRDepartment) LineOfDuty() {
	fmt.Println(hrd.Name, "员工招聘培训管理")
}

type FinanceDepartment struct {
	Name string
}

func NewFinanceDepartment(name string) *FinanceDepartment {
	return &FinanceDepartment{
		Name: name,
	}
}

func (fd *FinanceDepartment) Display(depth int) {
	fmt.Println(strings.Repeat("-", depth), " ", fd.Name)
}

func (fd *FinanceDepartment) LineOfDuty() {
	fmt.Println(fd.Name, "公司财务收支管理")
}
