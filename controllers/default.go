package controllers

import (
    beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
    beego.Controller
}

func (c *MainController) Get() {

    if GlobalError != "" {
        c.Data["Error"] = GlobalError
        GlobalError = "" // Clear the error message after retrieval
    }
  
    c.TplName = "index.tpl" // Use the name of your HTML template
}

