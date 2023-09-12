package main

import (
    _ "travel_guide/routers" // Import your routers package
    beego "github.com/beego/beego/v2/server/web"
    "github.com/beego/beego/v2/server/web/session"
)
var globalSessions *session.Manager

func main() {
    // Initialize Beego application
    sessionConfig := &session.ManagerConfig{
        CookieName:      "sessionID",
        EnableSetCookie: true,
        Gclifetime:      3600,    // Session lifetime in seconds (1 hour)
        Maxlifetime:     3600,    // Max session lifetime in seconds (1 hour)
        Secure:          false,   // Set to true if using HTTPS
        CookieLifeTime:  3600,    // Cookie lifetime in seconds (1 hour)
        ProviderConfig:  "./tmp", // Path to store session files
    }
    globalSessions, _ := session.NewManager("memory", sessionConfig)
    go globalSessions.GC()

    // Start the Beego application
    beego.Run()
}
