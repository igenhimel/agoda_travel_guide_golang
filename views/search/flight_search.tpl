<div class="row">
    <div class="col-md-10 offset-md-1">
        <div class="card shadow">
            <div class="card-body">
                <form class="container mb-3" action="/search/flight" method="get">
                    <div class="row">
                        <div class="col">
                            <div class="input-group py-3">
                                <div class="input-group-append bg-white">
                                    <button
                                        class="btn py-3 rounded-0 rounded-start border-secondary border-end-0 search-border"
                                        type="button" id="search-button"><i class="fas fa-search"></i></button>
                                </div>
                                <input type="text"
                                    class="form-control py-3 border border-secondary border-start-0 search-border"
                                    placeholder="Source Airport" aria-label="Search" aria-describedby="search-button"
                                    name="location_from">

                            </div>

                        </div>
                        <div class="col">
                            <div class="input-group py-3">
                                <div class="input-group-append bg-white">
                                    <button
                                        class="btn py-3 rounded-0 rounded-start border-secondary border-end-0 search-border"
                                        type="button" id="search-button"><i class="fas fa-search"></i></button>
                                </div>
                                <input type="text"
                                    class="form-control py-3 border border-secondary border-start-0 search-border"
                                    placeholder="Destination Airport" aria-label="Search"
                                    aria-describedby="search-button" name="location_to">

                            </div>
                        </div>
                    </div>
                    <div class="row">
                        <div class="col">
                            <div class="input-group py-3">
                                <div class="input-group-append bg-white">
                                    <button
                                        class="btn py-3 rounded-0 rounded-start border-secondary border-end-0 search-border"
                                        type="button" id="search-button"><i class="fa-solid fa-calendar-week"></i></button>
                                </div>
                                <input type="text"
                                    class="form-control py-3 border border-secondary border-start-0 search-border"
                                    placeholder="Departure or Travel date." aria-label="Search"
                                    aria-describedby="search-button" name="departure_date" id="travel-date" autocomplete="off" readonly>
        
                             </div>
                        </div>
                        <div class="col" style="display: none;" id="return-calendar">
                            <div class="input-group py-3">
                                <div class="input-group-append bg-white">
                                    <button
                                        class="btn py-3 rounded-0 rounded-start border-secondary border-end-0 search-border"
                                        type="button" id="search-button"><i class="fa-solid fa-calendar-week"></i></button>
                                </div>
                                <input type="text"
                                    class="form-control py-3 border border-secondary border-start-0 search-border"
                                    placeholder="Return Date" aria-label="Search"
                                    aria-describedby="search-button" name="return_date" id="return-date" autocomplete="off" readonly>
        
                             </div>
                        </div>
                        <div class="col">
                            <div class="input-group py-3">
                                <div class="input-group-prepend">
                                    <button class="btn py-3 rounded-0 rounded-start border-end-0 border-secondary search-border bg-white"
                                        type="button" id="search-button"><i class="fa-solid fa-plane-up"></i></button>
                                </div>
                                <select class="form-select py-3 border border-start-0 border-secondary search-border" aria-label="Itinerary Type"
                                    name="itinerary-type" id="itinerary-type">
                                    <option value="" selected disabled>Select Itinerary Type</option>
                                    <option value="ONE_WAY">ONE WAY</option>
                                    <option value="ROUND_TRIP">ROUND TRIP</option>
                                </select>
                            </div>
                        </div>
                       
                    </div>
                    <div class="row">
                       
                        <div class="col">
                            <div class="input-group py-3">
                                <div class="input-group-prepend ">
                                    <button class="btn py-3 rounded-0 rounded-start border-secondary search-border bg-white"
                                        type="button" id="search-button"><i class="fa-solid fa-users"></i></button>
                                </div>
                                
                                    <input type="text" class="form-control py-3" id="passenger-input" placeholder="Passengers" readonly value="1 Passengers">
                                    <div class="passenger-dropdown">
                                        <div class="dropdown-menu">
                                            <div class="passenger-type">
                                                <span class="label">Passengers</span>
                                                <span class="float-right">
                                                    <span class="btn btn-outline-primary rounded-0" id="decrement-adults">-</span>
                                                    <span class="count" id="adults-count">1</span>
                                                    <span class="btn btn-outline-primary rounded-0" id="increment-adults">+</span>
                                                </span>
                                            </div>
                                            
                                        </div>
                                    </div>
                                
                            </div>
                        </div>
                        
                        <div class="col">
                            <div class="input-group py-3">
                                <div class="input-group-prepend">
                                    <button class="btn py-3 rounded-0 rounded-start border-secondary search-border bg-white"
                                        type="button" id="search-button"><i class="fa-solid fa-plane-departure"></i></i></button>
                                </div>
                                <select class="form-select py-3 border border-start-0 border-secondary search-border" aria-label="Itinerary Type"
                                    name="class" id="class-service">
                                    <option value="" selected disabled>Select Class</option>
                                    <option value="Economy">ECONOMY</option>
                                    <option value="Business">PREMIUM ECONOMY</option>
                                    <option value="First">BUSINESS</option>
                                    <option value="Premium">FIRST</option>
                                </select>
                            </div>
                        </div>
                    </div>
                    
                




                    <div class="bundle">
                        <span class="btn btn-danger text-light rounded-0 p-0 px-1">Bundle & Save</span>
                    </div>

                    <button class="btn btn-primary search-btn-resturant shadow">SEARCH</button>
                </form>
            </div>
            </div>
        </div>
    </div>
</div>