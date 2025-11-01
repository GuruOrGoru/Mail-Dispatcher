# Email Dispatcher

A high-performance Go application for bulk email dispatching using a producer-consumer pattern with goroutines and rate limiting.

## Features

- **Concurrent Processing**: Utilizes 200 worker goroutines for parallel email sending
- **Rate Limiting**: Configured for 1 email per second with a burst capacity of 200 emails
- **CSV Input**: Loads recipient data from CSV files
- **Template Support**: Uses Go HTML templates for email content
- **SMTP Integration**: Sends emails via SMTP (configured for localhost:1025 by default)

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/guruorgoru/email-dispatcher.git
   cd email-dispatcher
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Build the application:
   ```bash
   go build -o main main.go
   ```

## Usage

1. Prepare your email data in `emails.csv` format:
   ```
   name,email
   John Doe,john@example.com
   Jane Smith,jane@example.com
   ```

2. Customize the email template in `email.templ`

3. Run the dispatcher:
   ```bash
   ./main
   ```

## Configuration

- **SMTP Settings**: Modify `smtpHost` and `smtpPort` in `internal/consumers/consumer.go`
- **Worker Count**: Adjust `workerCount` in `main.go`
- **Rate Limiting**: Modify the rate limiter parameters in `main.go`
- **Sender Email**: Update the from address in `internal/consumers/consumer.go`

## Performance

### Architecture
- **Producer-Consumer Pattern**: Single producer loads CSV data into a buffered channel (capacity 50)
- **Worker Pool**: 200 concurrent goroutines process emails from the channel
- **Rate Limiting**: Uses `golang.org/x/time/rate` limiter with 1 request/second and burst of 200

### Benchmarks
- **Throughput**: Can process up to 200 emails immediately, then 1 email per second
- **Memory Usage**: Minimal due to streaming CSV processing and bounded channel
- **CPU Usage**: Scales with worker count; optimized for I/O-bound email sending

### Optimization Notes
- Random sleep (0-50ms) between sends prevents thundering herd
- Buffered channel prevents memory spikes during CSV loading
- Graceful error handling ensures failed sends don't block the pipeline

## Dependencies

- `golang.org/x/time/rate` - Rate limiting
- Standard library: `net/smtp`, `html/template`, `encoding/csv`, `sync`

## Development

Run tests:
```bash
go test ./...
```

Build for production:
```bash
go build -ldflags="-s -w" -o main main.go
```

## License

This project is open source under the MIT License.