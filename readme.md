# Go Conference Ticket Booker CLI

A simple command-line interface (CLI) application built in Go for booking tickets to the "Go Conference".

## Description

This application simulates a ticket booking system. Users can input their details to book a certain number of tickets for the "Go Conference". The system validates the input, manages available ticket inventory, and simulates sending a confirmation ticket asynchronously.

## Features

*   **Command-Line Interface:** All interactions happen through the terminal.
*   **User Input:** Collects first name, last name, email address, and the number of tickets to book.
*   **Input Validation:**
    *   Checks if first and last names are at least 2 characters long.
    *   Checks if the email address contains an "@" symbol.
    *   Checks if the requested number of tickets is positive and does not exceed the remaining available tickets.
*   **Ticket Booking:**
    *   Stores booking information (name, email, tickets booked).
    *   Decrements the count of remaining tickets.
*   **Inventory Management:** Tracks the total and remaining tickets for the conference.
*   **Sold Out Notification:** Informs the user when the conference is fully booked.
*   **Asynchronous Ticket Sending:** Simulates sending tickets in the background using goroutines (with a 5-second delay for demonstration).
*   **Display Bookings:** Shows a list of first names of users who have booked tickets.

## Prerequisites

*   [Go](https://golang.org/doc/install) (version 1.x or higher) installed on your system.

## How to Run

1.  **Clone or Download:**
    If this project were in a repository, you would clone it. For now, ensure you have `main.go` and `helper.go` in the same directory.

2.  **Navigate to the Directory:**
    Open your terminal or command prompt and navigate to the directory where you saved the Go files.
    ```bash
    cd path/to/your/CLI-ticket-booking-Go
    ```

3.  **Run the Application:**
    You can run the application directly using:
    ```bash
    go run main.go helper.go
    ```
    Alternatively, you can build an executable:
    ```bash
    go build -o ticket-booker main.go helper.go
    ```
    And then run the executable:
    ```bash
    ./ticket-booker  # On Linux/macOS
    .\ticket-booker.exe # On Windows
    ```

4.  **Interact with the Application:**
    The application will greet you and prompt for your details:
    *   Enter your first name
    *   Enter your last name
    *   Enter your email address
    *   Enter the number of tickets you want to book

    Follow the on-screen prompts. If your input is valid, your booking will be confirmed, and a simulated ticket-sending process will begin.

## Project Structure

*   `main.go`: Contains the main application logic, user interaction, booking functions, and goroutine management.
*   `helper.go`: Contains helper functions, primarily for input validation.

## Potential Future Enhancements

*   More robust error handling for user input (e.g., non-integer input for tickets).
*   Persistent storage for bookings (e.g., saving to a file like CSV or JSON, or a database).
*   Option to view all booking details (not just first names).
*   Option to cancel bookings.
*   Improved UI/UX for error messages (e.g., consolidating all validation errors).