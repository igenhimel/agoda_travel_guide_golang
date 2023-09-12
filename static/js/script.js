$(document).ready(function() {
    // Get references to the select element and the return-calendar div
    var itineraryTypeSelect = $('#itinerary-type');
    var returnCalendarDiv = $('#return-calendar');

    // Initial state based on selected value on page load
    if (itineraryTypeSelect.val() === 'ROUND_TRIP') {
        returnCalendarDiv.show();
    } else {
        returnCalendarDiv.hide();
    }

    // Add an event listener for select change
    itineraryTypeSelect.on('change', function() {
        if (itineraryTypeSelect.val() === 'ROUND_TRIP') {
            returnCalendarDiv.show();
        } else {
            returnCalendarDiv.hide();
        }
    });
});

//date-picker
  // Get references to the travel-date and return-date input fields
  var travelDateInput = $('#travel-date');
  var returnDateInput = $('#return-date');

  // Initialize Pikaday for travel date with a minimum date (today)
  var travelDatePicker = new Pikaday({
      field: travelDateInput[0], // Get the DOM element from jQuery object
      toString(date, format) {
        const day = date.getDate();
        const month = date.getMonth() + 1;
        const year = date.getFullYear();
        return `${year}-${month < 10 ? '0' + month : month}-${day < 10 ? '0' + day : day}`;
    },
      minDate: new Date(), // Set the minimum date to today
      onSelect: function() {
          // Calculate the minimum date for the return date picker as the day after the selected travel date
          var selectedTravelDate = this.getDate();
          var nextDay = new Date(selectedTravelDate);
          nextDay.setDate(nextDay.getDate() + 1);
          returnDatePicker.setMinDate(nextDay);
          returnDateInput.val("");
      }
  });

  // Initialize Pikaday for return date with no initial minDate
  var returnDatePicker = new Pikaday({
      field: returnDateInput[0], // Get the DOM element from jQuery object
      toString(date, format) {
        const day = date.getDate();
        const month = date.getMonth() + 1;
        const year = date.getFullYear();
        return `${year}-${month < 10 ? '0' + month : month}-${day < 10 ? '0' + day : day}`;
    },
  });

  var passengerDropdown = $(".passenger-dropdown");

  // Function to update the #passenger-input value
  function updatePassengerInput() {
      var adultsCount = parseInt($("#adults-count").text());
      var seniorsCount = parseInt($("#seniors-count").text());

      // Ensure that adultsCount is always at least 1
      if (adultsCount < 1) {
          adultsCount = 1;
          $("#adults-count").text(adultsCount);
      }

      // Generate the string for the input field
      var inputText = "";

      if (adultsCount > 0) {
          inputText += adultsCount + " Passenger" + (adultsCount !== 1 ? "s" : "");
      }

      // Update the #passenger-input value
      $("#passenger-input").val(inputText);
  }

  // Toggle passenger dropdown on input field click
  $("#passenger-input").click(function (e) {
      e.stopPropagation(); // Prevent the click event from propagating to the document
      passengerDropdown.toggleClass("open");
  });

  // Increment and Decrement functions for adults
  $("#increment-adults").click(function () {
      var adultsCount = parseInt($("#adults-count").text());
      $("#adults-count").text(adultsCount + 1);
      updatePassengerInput(); // Update the input field
  });

  $("#decrement-adults").click(function () {
      var adultsCount = parseInt($("#adults-count").text());
      if (adultsCount > 1) {
          $("#adults-count").text(adultsCount - 1);
      }
      updatePassengerInput(); // Update the input field
  });

  // Increment and Decrement functions for seniors
  $("#increment-seniors").click(function () {
      var seniorsCount = parseInt($("#seniors-count").text());
      $("#seniors-count").text(seniorsCount + 1);
      updatePassengerInput(); // Update the input field
  });

  $("#decrement-seniors").click(function () {
      var seniorsCount = parseInt($("#seniors-count").text());
      if (seniorsCount > 0) {
          $("#seniors-count").text(seniorsCount - 1);
      }
      updatePassengerInput(); // Update the input field
  });

  // Close passenger dropdown when clicking outside
  $(document).on('click', function (e) {
      if (!passengerDropdown.is(e.target) && passengerDropdown.has(e.target).length === 0) {
          passengerDropdown.removeClass("open");
      }
  });
//pass count


$("#close-error").click(function() {
    $("#error-alert").css("height", "0");
});

