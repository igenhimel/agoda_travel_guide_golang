$(document).ready(function() {
    // Show hotel container and hide others by default
    $("#hotel").click(function() {
        $("#hotel-container").show();
        $("#resturant-container").hide();
        $("#flight-container").hide();
        $('.btn-active').removeClass('btn-active');
        $(this).addClass('btn-active');
        $("#title-text").html("BOOK A HOME ON AGODA HOMES<br><small>More spacious. More local. More of why you travel.</small>");
    });

    $("#resturant").click(function() {
        $("#hotel-container").hide();
        $("#flight-container").hide();
        $("#resturant-container").show();
        $('.btn-active').removeClass('btn-active');
        $(this).addClass('btn-active');
        $("#title-text").html("HOTELS, RESORTS, HOSTELS & MORE<br><small>Get the best prices on 2,000,000+ properties, worldwide</small>");
    });

    $("#flight").click(function() {
        $("#hotel-container").hide();
        $("#resturant-container").hide();
        $("#flight-container").show();
        $('.btn-active').removeClass('btn-active');
        $(this).addClass('btn-active');
        $("#title-text").html("FLIGHTS<br><small>Book your flights here.</small>");
    });
});
