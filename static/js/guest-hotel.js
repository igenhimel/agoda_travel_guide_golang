
var countRoom = document.getElementById("count-value-room");
var countAdult = document.getElementById("count-value-adult");
var countChild = document.getElementById("count-value-child");

$(document).ready(function () {
    
    // Toggle the custom dropdown on input click
    $("#guests-input").click(function () {
        $(".custom-dropdown").toggle();
    });

    // Event listener for increment and decrement buttons
    $(".increment").click(function (e) {
        e.stopPropagation(); // Prevent click event from propagating
        var option = $(this).closest(".guest-option");
        var countElement = option.find(".count");
        var count = parseInt(countElement.text());
        var min = parseInt($(this).data("min")); // Minimum value

        count = isNaN(count) ? min : count; // Handle NaN
        count++;
        countElement.text(count);

        if (option.hasClass("children-option")) {
            addChildrenAgeInputFields(count);
        }

        // If room count is incremented, also increment adults
        if (option.hasClass("rooms-option")) {
            var adultOption = $(".adults-option");
            var adultCountElement = adultOption.find(".count");
            var adultCount = parseInt(adultCountElement.text());

            if (adultCount < count) {
                adultCountElement.text(count);
            }
        }

        updateGuestsInput();
    });

    $(".decrement").click(function (e) {
        e.stopPropagation(); // Prevent click event from propagating
        var option = $(this).closest(".guest-option");
        var countElement = option.find(".count");
        var count = parseInt(countElement.text());
        var min = parseInt($(this).data("min")); // Minimum value

        count = isNaN(count) ? min : count; // Handle NaN

        if (count > min) {
            count--;
            countElement.text(count);

            if (option.hasClass("children-option")) {
                removeChildrenAgeInputFields(count);
            }

            updateGuestsInput();
        }
        
    });

    // Close the custom dropdown when clicking outside
    $(document).on("click", function (e) {
        if (!$(e.target).closest(".input-group").length) {
            $(".custom-dropdown").hide();
        }
    });

    function updateGuestsInput() {
        var rooms = parseInt($(".guest-option:contains('Rooms') .count").text());
        var adults = parseInt($(".guest-option:contains('Adults') .count").text());
        var children = parseInt($(".guest-option:contains('Children') .count").text());

        countRoom.value = rooms;
        countAdult.value = adults;
        countChild.value = children;

        // Ensure counts are at least 1 for Rooms and Adults
        rooms = Math.max(1, rooms);
        adults = Math.max(rooms, adults);

        // Ensure count is at least 0 for Children
        children = Math.max(0, children);

        var guestsInput = rooms + " rooms | " + adults + " adults";

        if (children > 0) {
            guestsInput += " | " + children + " children";

            // Add or remove child age input fields based on the count
            var existingCount = $(".child-ages .form-group").length;
            if (children > existingCount) {
                addChildrenAgeInputFields(children - existingCount);
            } else if (children < existingCount) {
                removeChildrenAgeInputFields(existingCount - children);
            }
        }

        if (adults == rooms) {
            $(".adults-option .decrement").addClass("disabled");
        } else {
            $(".adults-option .decrement").removeClass("disabled");
        }

        if(rooms<2){
            $(".rooms-option .decrement").addClass("disabled");
        }
        else{
            $(".rooms-option .decrement").removeClass("disabled");
        }

        $("#guests-input").val(guestsInput);
    }

    // Function to add select dropdowns for children's ages
    function addChildrenAgeInputFields(count) {
        var existingCount = $(".child-ages .form-group").length;

        if (count > existingCount) {
            for (var i = existingCount + 1; i <= count; i++) {
                var container = '<div class="form-group row">'; // Create a div container with a row class
                var label = '<small for="child-age-' + i + '" class="col-sm-4 col-form-label text-muted">Age of Child ' + i + '</small>';
                var selectDropdown = '<div class="col-sm-3"><select id="child-age-' + i + '" class="form-select child-age-input" onclick="showAgeDropdown(' + i + ')">';
                for (var age = 1; age <= 17; age++) {
                    selectDropdown += '<option value="' + age + '">' + age + '</option>';
                }
                selectDropdown += '</select></div>';
                container += label + selectDropdown + '</div>'; // Close the div container
                $(".child-ages").append(container);
            }
        }
    }

    // Function to remove select dropdowns for children's ages
    function removeChildrenAgeInputFields(count) {
        var existingCount = $(".child-ages .form-group").length;

        if (count <= existingCount) {
            $(".child-ages .form-group:gt(" + (count - 1) + ")").remove();
        }
        if (existingCount == 1) {
            $(".child-ages .form-group").remove();
        }
    }

    // Function to stop event propagation to prevent closing the dropdown when clicking on it
    function stopPropagation(event) {
        event.stopPropagation();
    }

    // Add a click event listener to the document for dynamically added select dropdowns
    $(document).on("click", ".child-age-input", function (event) {
        event.stopPropagation(); // Prevent click event from propagating
        var childNumber = $(this).attr("id").replace("child-age-", "");
        showAgeDropdown(childNumber);
    });

    // Initial setup to add the first select dropdown
    addChildrenAgeInputFields(0); // Set the initial count to 1 and show the first select dropdown
});