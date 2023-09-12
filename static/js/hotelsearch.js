
    // Function to fetch hotel data via AJAX
    function fetchHotelData() {
        const destination = $("#destination").val();
        const checkIn = $(".t-input-check-in").eq(0).val();
        const checkOut = $(".t-input-check-out").eq(0).val();

        const url = `/search/hotel-list/?destination=${destination}&t-start=${checkIn}&t-end=${checkOut}`;

        $.ajax({
            url: url,
            method: "GET",
            dataType: "json",
            success: function (data) {
                if (data.Error) {
                    // Handle errors here (display an error message)
                    console.error(data.Error);
                } else {
                    // Update the page with hotel data
                    const hotelList = $("#hotel-list");
                    hotelList.empty(); // Clear previous data
                    if(data.data==null){
                        const hotels = `<div class="text-center w-100">
                        <h1 class="text-danger">HOTEL NOT FOUND! SEARCH AGAIN</h1>
                        <a class="btn btn-warning" href="/">Return to Home Page</a>
                    </div>`
                    hotelList.append(hotels)
                    }
                   else{
                    data.data.forEach(function (hotel) {

                      
                        const hotelCard = `
                
                            <div class="col">
                                <div class="card h-100">
                                    <div class="img-card">
                                        <img src="https://cf.bstatic.com${hotel.basicPropertyData.photos.main.highResJpegUrl.relativeUrl}" class="card-img-top" alt="Hollywood Sign on The Hill Lazy-loaded">
                                        <span class="btn btn-light btn-money btn-price">From <i class="fa-solid fa-bangladeshi-taka-sign"></i> ${hotel.priceDisplayInfoIrene.displayPrice.amountPerStay.amountRounded}</i></span>
                                        <div class="btn-mores">
                                            <span class="btn btn-light opacity-75 btn-love btn-love"><i class="fa-regular fa-heart"></i></span>
                                            <span class="btn btn-light opacity-75 btn-map btn-map"><i class="fa-solid fa-location-dot"></i></span>
                                        </div>
                                    </div>
                                    <div class="card-body">
                                        <h6 class="card-title">${hotel.displayName.text}</h6>
                                        <p class="card-text">
                                            <i class="fa-solid fa-star"></i>
                                            <i class="fa-solid fa-star"></i>
                                            <i class="fa-solid fa-star"></i>
                                            <i class="fa-solid fa-star"></i>
                                            <i class="fa-regular fa-star"></i>
                                            683 Reviews
                                            <h6>Sleeps 85</h6>
                                            <a href="#" class="textstyle">Arizona > Lake Havasu City</a>
                                            <div class="text-end">
                                                <a href="/search/hotel/details/?IDDetail=${hotel.idDetail}" class="btn btn-warning">View Availability</a>
                                            </div>
                                            <div class="mt-3">
                                                <div class="row">
                                                    <div class="col-auto">
                                                        <small class="mt-2"><i class="fa-solid fa-check"></i> View  <i class="fa-solid fa-check"></i> Child Friendly</small>
                                                    </div>
                                                    <div class="col text-right text-end">
                                                        <small class="BC">BC - 367639</small>
                                                    </div>
                                                </div>
                                            </div>
                                        </p>
                                    </div>
                                </div>
                            </div>
                         
                        `;
                        hotelList.append(hotelCard);
                    
                    });
                }
                }
            },
            error: function (xhr, status, error) {
                // Handle network errors here
                console.error("Network error:", error);
            }
            
        });
    }

    // Attach a click event listener to a button (assuming you have a button with id "fetch-button")
    $("#fetch-button").on("click", fetchHotelData);

