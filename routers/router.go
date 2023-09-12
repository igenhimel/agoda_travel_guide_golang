// routers/router.go

package routers

import (
    beego "github.com/beego/beego/v2/server/web"
    "travel_guide/controllers"
    
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/search/hotel", &controllers.HotelController{}, "get:Get")
    beego.Router("/search/hotel/details", &controllers.HotelControllerDetails{})
    beego.Router("/search/flight", &controllers.FlightController{},"get:Get")

   


}
