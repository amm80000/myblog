package controllers

type ExitController struct {
	BaseController
}

func (this *ExitController) Get() {
	this.DelSession("loginUser")
	this.Redirect("/", 302)
}
