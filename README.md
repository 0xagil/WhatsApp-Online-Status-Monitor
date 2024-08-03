# WhatsApp Status Monitor

WhatsApp Status Monitor is a tool that allows you to track the online/offline status of selected WhatsApp contacts. It provides a web interface to select contacts and view their status over time.

## Features

- Select WhatsApp contacts to monitor
- Real-time status updates
- Visual representation of online/offline periods
- Responsive web interface

## Prerequisites

Before you begin, ensure you have met the following requirements:

- Go 1.22 or higher installed on your system
- A WhatsApp account
- Basic knowledge of terminal/command line operations

## Installation

1. Clone the repository:

``
git clone https://github.com/0xagil/WhatsApp-Online-Status-Monitor.git
``
2. Install the required Go packages:

``
go mod tidy
``

## Usage

1. Run the application:

``
go run main.go
``

2. The application will start and automatically open your default web browser to `http://localhost:8080`. If it doesn't, manually open this URL in your browser.

3. On first run, you'll need to authenticate your WhatsApp account:
- Scan the QR code displayed in the terminal using your WhatsApp mobile app.
- Once authenticated, the QR code will disappear, and the application will start running.

4. In the web interface:
- Go to the "Select Contacts" page to choose which contacts you want to monitor.
- Use the "View Status" page to see the real-time status of your selected contacts.

5. To stop the application, press `Ctrl+C` in the terminal where it's running.

## Notes

- This application does not store or share any of your WhatsApp messages. It only tracks the online/offline status of the contacts you choose to monitor.
- Continuous use of this tool may drain your device's battery faster than normal.
- Be mindful of privacy concerns and only use this tool responsibly and with the knowledge of the contacts you're monitoring.

## Contributing

Contributions to the WhatsApp Status Monitor are welcome. Please feel free to submit a Pull Request.

## License

This project is licensed under the [MIT License](LICENSE).

## Disclaimer

This project is not affiliated with, authorized, maintained, sponsored or endorsed by WhatsApp or any of its affiliates or subsidiaries. This is an independent and unofficial software. Use at your own risk.