package main

import (
    _ "travel_guide/routers" // Import your routers package
    beego "github.com/beego/beego/v2/server/web"
)

func main() {
    // Initialize Beego application
    beego.BConfig.WebConfig.Session.SessionOn = true
    beego.BConfig.WebConfig.Session.SessionProvider = "memory"
    beego.BConfig.WebConfig.Session.SessionName = "gosessionid"
    beego.BConfig.WebConfig.Session.SessionGCMaxLifetime = 3600 // Set your desired lifetime in seconds

    // Start the Beego application
    beego.Run()
}
