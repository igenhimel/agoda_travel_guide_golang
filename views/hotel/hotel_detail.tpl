{{template "partials/header.tpl" .}}

{{template "partials/nav.tpl" .}}
<!-- Photo URIs Section -->
<h2 class="text-center bg-secondary shadow text-light">Hotel Description for: {{.HotelTitle}}</h2>
<div class="section">
    <div class="container">
      
        <h1>Hotel Photos</h1>
        <div class="row">
            {{ range .PhotoURIs }}
            <div class="col-md-3 mb-4">
                <div class="photo-card">
                    <img src="https://cf.bstatic.com{{ . }}" class="card-img-top" alt="Hotel Photo">
                </div>
            </div>
            {{ end }}
        </div>
    </div>
</div>

<!-- Description Section -->
<div class="section">
    <div class="container">
        <h1>Hotel Description</h1>
        <div class="row">
            <div class="col">
                <div class="description">
                    <p>{{ .Description }}</p>
                </div>
            </div>
        </div>
    </div>
</div>

<!-- Facility Titles Section -->
<div class="section">
    <div class="container">
        <h1>Amenities</h1>
        <div class="facility-titles">
            <ul class="facility-list">
                {{ range .FacilityTitles }}
                <li style="display: inline-block; margin-right: 10px;">{{ . }}</li>
                {{ end }}
            </ul>
        </div>
    </div>
</div>

{{template "partials/footer.tpl" .}}