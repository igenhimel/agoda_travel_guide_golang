<nav class="navbar-light bg-light">

    <form class="navbar-form container" action="/search/hotellist" method="get">
        <div class="row">
            <!-- First Search Bar -->
            <div class="col-3">
                <div class="input-group py-2">
                    <div class="input-group-append bg-white">
                        <button class="btn py-2 rounded-0 rounded-start border-secondary border-end-0 search-border"
                            type="button" id="search-button"><i class="fas fa-search"></i></button>
                    </div>
                    <input type="text"
                        class="form-control py-2 border border-secondary border-start-0 search-border"
                        placeholder="Enter a Destination or Property" aria-label="Search"
                        aria-describedby="search-button" name="destination" id="destination">
                </div>
            </div>

            <!-- Second Search Bar -->
            <div class="col-5 mt-1">
                <div class="t-datepicker">
                    <div class="t-check-in bg-white py-1"></div>
                    <div class="t-check-out bg-white py-1"></div>
                </div>
            </div>

            <!-- Third Search Bar -->
            <div class="col-3 mt-2">
                <div class="input-group">
                    <!-- Input field for number of guests -->
                    <div class="input-group-append bg-white">
                        <button
                            class="btn py-2 rounded-0 rounded-start border-secondary border-end-0 search-border"
                            type="button" id="search-button"><i class="fa-solid fa-users"></i></button>
                    </div>
                    <input autocomplete="off"
                        class="form-control py-2 border-secondary search-border border-start-0"
                        placeholder="Guests" aria-label="Guests" id="guests-input">
                    <!-- Dropdown for selecting the number of guests -->
                    <div class="custom-dropdown container">
                        <div class="guest-option rooms-option">
                            <span>Rooms</span>
                            <span class="count-controls">
                                <span class="btn btn-sm btn-outline-primary decrement decrement-room disabled" data-min="1">-</span>
                                <span class="count">1</span>
                                <span class="btn btn-sm btn-outline-primary increment increment-room" data-min="1">+</span>
                            </span>
                            <input type="hidden" name="count-value-room" value="1" id="count-value-room">
                        </div>
                        <div class="guest-option adults-option">
                            <span>Adults <small class="text-muted">(Ages 18 or above)</small></span>
                            <span class="count-controls">
                                <span class="btn btn-sm btn-outline-primary decrement decrement-adult disabled" data-min="1">-</span>
                                <span class="count">1</span>
                                <span class="btn btn-sm btn-outline-primary increment increment-adult" data-min="1">+</span>
                            </span>
                            <input type="hidden" name="count-value-adult" value="1" id="count-value-adult">
                        </div>
                        <div class="guest-option children-option">
                            <span>Children <small class="text-muted">(Ages 1-17)</small></span>
                            <span class="count-controls">
                                <span class="btn btn-sm btn-outline-primary decrement decrement-children" data-min="0">-</span>
                                <span class="count">0</span>
                                <span class="btn btn-sm btn-outline-primary increment increment-children" data-min="0">+</span>
                            </span>
                            <input type="hidden" name="count-value-child" value="0" id="count-value-child">
                        </div>
                        <div class="child-ages">
                            <!-- Children ages input fields will be added here -->
                        </div>
                    </div>



                </div>
            </div>

            <div class="col mt-2">
                <button class="btn btn-outline-primary w-100" id="searchButton">SEARCH</button>
            </div>
        </div>


    </form>

    </div>
</nav>