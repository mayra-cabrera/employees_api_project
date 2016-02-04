package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["employees_api_project/controllers:DepartmentsController"] = append(beego.GlobalControllerRouter["employees_api_project/controllers:DepartmentsController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["employees_api_project/controllers:DepartmentsController"] = append(beego.GlobalControllerRouter["employees_api_project/controllers:DepartmentsController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["employees_api_project/controllers:DepartmentsController"] = append(beego.GlobalControllerRouter["employees_api_project/controllers:DepartmentsController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["employees_api_project/controllers:DepartmentsController"] = append(beego.GlobalControllerRouter["employees_api_project/controllers:DepartmentsController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["employees_api_project/controllers:DepartmentsController"] = append(beego.GlobalControllerRouter["employees_api_project/controllers:DepartmentsController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["employees_api_project/controllers:EmployeesController"] = append(beego.GlobalControllerRouter["employees_api_project/controllers:EmployeesController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["employees_api_project/controllers:EmployeesController"] = append(beego.GlobalControllerRouter["employees_api_project/controllers:EmployeesController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["employees_api_project/controllers:EmployeesController"] = append(beego.GlobalControllerRouter["employees_api_project/controllers:EmployeesController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["employees_api_project/controllers:EmployeesController"] = append(beego.GlobalControllerRouter["employees_api_project/controllers:EmployeesController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["employees_api_project/controllers:EmployeesController"] = append(beego.GlobalControllerRouter["employees_api_project/controllers:EmployeesController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

}
