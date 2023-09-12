# Beego Travel Guide Project

This repository contains a Beego-based web application that serves as a travel guide. It allows users to search for hotels based on their destination, check-in, and check-out dates. The application leverages the Booking.com API to fetch hotel information and display it to the user.

## Prerequisites

Before you can run this project, ensure you have the following prerequisites installed on your system:

- Go: [Installation Guide](https://golang.org/doc/install)
- Beego: [Installation Guide](https://beego.me/docs/install)
- Booking.com API Key: You need to sign up on RapidAPI to obtain an API key. Replace `"88eef088ffmshe0050347753f521p138692jsnb84bd17606db"` with your actual API key in the code.

## Getting Started

Follow these steps to get the project up and running:

1. Clone this repository to your local machine:

   ```shell
   git clone https://github.com/igenhimel/beego_travel_guide_project.git
   cd beego_travel_guide_project/travel_guide
   ```

2. Install the project dependencies using Go Modules:

   ```shell
   go mod tidy
   ```

3. Run the Beego development server:

   ```shell
   bee run
   ```

   The application should now be running locally at `http://localhost:8080`.

## Usage

1. Open a web browser and navigate to `http://localhost:8080`.
2. You will be presented with a search form where you can enter the following information:
   - **Destination**: Enter the name of a country, city, airport, neighborhood, landmark, or property.
   - **Check-In Date**: Choose your desired check-in date.
   - **Check-Out Date**: Choose your desired check-out date.

3. Click the "Search" button to fetch hotel data based on your input.

## Features

- Search for hotels based on destination and dates.
- Displays a list of hotels including images, titles, and prices.
- Handles errors gracefully and provides informative error messages.

## Code Structure

The code for this project is structured as follows:

- `controllers/HotelController.go`: Defines the HotelController, which handles hotel search and display.
- `models/HotelInfo.go`: Defines the data structures for parsing JSON responses from the Booking.com API.
- `views/hotel/hotelList.tpl`: Defines the HTML template for displaying hotel information.


## Acknowledgments

- [Beego](https://beego.me/) - A powerful and flexible Go web framework.
- [RapidAPI](https://rapidapi.com/) - Provides access to a variety of APIs, including the Booking.com API used in this project.

Feel free to contribute to this project or report any issues you encounter. Happy coding!