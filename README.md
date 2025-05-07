# HTML to Image Converter

A high-performance web service that converts HTML content to images using ChromeDP (Chrome DevTools Protocol). This service provides a simple REST API endpoint to convert HTML strings to PNG images with customizable dimensions and quality.

## Features

- Convert HTML to PNG images
- Customizable image dimensions
- Adjustable image quality
- Full-page capture support
- Base64 encoded image output
- CORS enabled for cross-origin requests
- Health check endpoint
- High-performance rendering using ChromeDP

## Prerequisites

- Go 1.21 or later
- Chrome/Chromium browser
- Windows, Linux, or macOS

## Installation

### Windows

1. Clone the repository:
```bash
git clone https://github.com/yourusername/html-to-image.git
cd html-to-image
```

2. Run the setup script:
```bash
setup.bat
```

### Linux/macOS

1. Clone the repository:
```bash
git clone https://github.com/yourusername/html-to-image.git
cd html-to-image
```

2. Make the setup script executable:
```bash
chmod +x setup.sh
```

3. Run the setup script:
```bash
./setup.sh
```

## Usage

### Starting the Server

#### Windows
```bash
run.bat
```

#### Linux/macOS
```bash
./html-to-image
```

The server will start on port 3000 by default. You can change the port by setting the `PORT` environment variable.

### API Endpoints

#### Convert HTML to Image
```http
POST /html-to-image
Content-Type: application/json

{
    "html": "<div>Your HTML content here</div>",
    "width": 800,
    "height": 600,
    "full_page": false,
    "quality": 90
}
```

Parameters:
- `html` (required): The HTML content to convert
- `width` (optional): Image width in pixels (default: 800)
- `height` (optional): Image height in pixels (default: 600)
- `full_page` (optional): Whether to capture the full page (default: false)
- `quality` (optional): Image quality (1-100, default: 90)

Response:
```json
{
    "success": true,
    "image_data": "base64_encoded_image_data"
}
```

#### Health Check
```http
GET /
```

Response:
```json
{
    "status": "healthy"
}
```

### Example Usage with cURL

```bash
curl -X POST http://localhost:3000/html-to-image \
  -H "Content-Type: application/json" \
  -d '{
    "html": "<div style=\"background: red; width: 100px; height: 100px;\"></div>",
    "width": 200,
    "height": 200
  }'
```

## Error Handling

The API returns appropriate HTTP status codes and error messages:

- 400 Bad Request: Invalid input parameters
- 500 Internal Server Error: Server-side errors

Error response format:
```json
{
    "success": false,
    "message": "Error description"
}
```

## Development

### Project Structure
```
html-to-image/
├── main.go           # Main application code
├── go.mod           # Go module file
├── go.sum           # Go module checksum
├── setup.sh         # Linux/macOS setup script
├── setup.bat        # Windows setup script
└── README.md        # This file
```

### Dependencies

- [Fiber](https://github.com/gofiber/fiber) - Web framework
- [ChromeDP](https://github.com/chromedp/chromedp) - Chrome DevTools Protocol

## License

MIT License

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request 