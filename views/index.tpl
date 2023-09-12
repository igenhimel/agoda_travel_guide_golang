{{template "partials/header.tpl" .}}

{{ if .Error }}
<div id="error-alert" class="fixed-top d-flex justify-content-center align-items-center" style="height: 100vh; background-color: rgba(0, 0, 0, 0.5); z-index: 999;">
    <div class="alert alert-light alert-dismissible show" style="max-width: 400px; text-align: center; padding: 50px; border-radius: 5px; position: relative;">
        {{ .Error }}
        <button id="close-error" type="button" class="btn-close" data-bs-dismiss="alert" aria-label="Close"></button>
    </div>
</div>
{{ end }}




{{template "partials/nav.tpl" .}}
<div class="covid text-center p-2">
    <span><img src="https://img.agoda.net/images/INTTRV-45/default/Bags-heart_2021-09-30.svg" alt=""></span>
    <span>Traveling internationally? Get updated information on COVID-19 travel guidance and restrictions.</span>
    <a href="#" class="btn btn-light text-primary border">Learn more</a>
</div>



<div class="home-cover">
    <p class="h3" id="title-text">SEE THE WORLD FOR LESS</p>
    <div class="nav-btn">
        <span class="bg-white px-5 py-3 border rounded-pill shadow">
            <!-- Add an "onclick" attribute to each button to trigger a JavaScript function -->
            <button class="btn btn-active datepicker-trigger border-0" id="hotel"> <i class="fas fa-hotel"></i>
                Hotel & Homes</button>
          
            <button class="btn datepicker-trigger border-0" id="flight"> <i class="fa-solid fa-plane"></i>
                Flights</button>

        </span>
    </div>
    <div class="container search-nav">
        <section id="hotel-container">
            {{template "search/hotel_search.tpl" .}}
        </section>
        <section id="flight-container" style="display: none;">
            {{template "search/flight_search.tpl" .}}
        </section>
    </div>




</div>

{{template "partials/footer.tpl" .}}