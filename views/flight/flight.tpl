{{template "partials/header.tpl" .}}
{{template "partials/nav.tpl" .}}

<div class="container bg-light">
   

    <div class="mt-4">
        
  
    
        <h2>Showing {{.Flight.FlightsCount}} flights from 
            {{range $index, $route := .Flight.Routes}}
            {{$route.Origin.CityCode}}
            to
            {{$route.Destination.CityCode}}
            {{end}}
        </h2>
        <p>Price per passenger includes tax and fees</p>
    </div>

    <div class="mt-4">
        {{ $travelDate := .travel_date }}
        {{ $returnDate := .return_date }}
        {{range $index, $flight := .Flight.Data}}
        <div class="card w-100 mx-2 mt-2">
            <div class="card-header bg-danger text-light small">Available Flight</div>
            <div class="card-body">
                <div class="row">
                    <div class="col-8">
                        <div class="icon">
                            <span class="btn btn-outline-light me-2 text-dark border">Eco Saver</span>
                            <span><i class="fas fa-suitcase-rolling me-2"></i></span>
                            <span><i class="fas fa-plane-departure text-success me-2"></i></span>
                            <span><i class="fas fa-chair me-2"></i></span>
                        </div>

                        <div class="time">
                            <p class="small ms-2">Travel Date : {{ $travelDate }}</p>
                            {{if $returnDate}}
                            <p class="small ms-2">Return Date : {{ $returnDate }}</p>
                            {{end}}
                        </div>
                        <p class="text-success me-2 small"><i class="fas fa-check me-2"></i>Free Cancellation within 42 hours of booking completion</p>
                    </div>
                    
                    <div class="col-4">
                        <div>
                            {{range $j, $travelerPrice := $flight.TravelPrices}}
                            <span class="currency-code text-danger lead fw-bold" style="vertical-align: middle;">
                                <sub>{{$travelerPrice.Price.Price.Currency.Code}}</sub>
                                <span class="fs-3" style="vertical-align: sub;">{{$travelerPrice.Price.Price.Value}}</span>
                            </span>                            
                            {{end}}
                        </div>
                        
                        <div class="url">
                            <a href="{{$flight.ShareableUrl}}" class="text-primary text-decoration-none"><i class="fa-solid fa-caret-down me-2"></i> View Details</a>
                        </div>
                    </div>
                </div>
            </div>
        </div>
        {{end}}
    </div>
</div>

{{template "partials/footer.tpl" .}}
