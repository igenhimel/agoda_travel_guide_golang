<!DOCTYPE html>
<html>
<head>
    <title>Flight Information</title>
</head>
<body>
    <h1>Flight Information</h1>
    
    {{range .FlightData.Flights}}
        <div class="flight-info">
            <p>ID: {{.ID}}</p>
            <p>Shareable URL: {{.ShareableURL}}</p>
            <!-- Display other flight information as needed -->
        </div>
    {{end}}
</body>
</html>
