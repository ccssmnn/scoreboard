# Scoreboard - A self hosted Google Hashcode Scoreboard

Disclaimer: This is not the official implementation.

## Why?

Google Hashcode problems are mostly large scale combinatorical optimization problems. Unfortunately the official scoreboard is shut down shortly after the event finishes. This project allows you to:

- prepare for the next Hashcode event
- experiment with your solutions after the event is finished
- experiment with black box optimization techniques

## Usage

### HTTP API

1. Clone the project
2. Start the server `go run main.go`
3. Run GET requests to receive problem files: `localhost:8080/static/2020/qualification/a_example.txt`
4. Run POST request to submit your solutions: `localhost:8080/score/2020/qualification/a_example` and put your submission file into the request body

### Go

Use the `scoreboard/obj` subpackage to parse problem and solution strings and compute the result.

## License

MIT
