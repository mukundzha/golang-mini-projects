# CLI Calculator

A simple interactive command-line calculator written in Go. It supports addition, subtraction, multiplication, and division on floating-point numbers, with basic input validation.

## Features

- REPL-style loop: perform as many calculations as you want without restarting
- Supports `+`, `-`, `*`, and `/`
- Handles decimal numbers
- Protects against division by zero
- Reports invalid operators
- Exit by entering `0` as the first number

## Requirements

- Go 1.26.8 or later

## Installation

Clone the repository and build the binary:

```sh
git clone <repository-url>
cd cli-calculator
go build -o cli-calculator
```

Or run it directly without building:

```sh
go run main.go
```

## Usage

Start the program:

```sh
./cli-calculator
```

You will be prompted to enter the first number, an operator, and the second number:

```
Welcome to Calculator!
Enter first number (or 0 to exit):
10
Enter operator (+, -, *, /):
+
Enter second number:
5
Result: 15
```

### Exiting

Enter `0` when prompted for the first number:

```
Enter first number (or 0 to exit):
0
Goodbye!
```

> Note: since `0` is the exit signal, the first number cannot be zero. To work around this, either use a different first operand or restructure the calculation (e.g., `5 * 0`).

## Example Calculations

| Input           | Result |
| --------------- | ------ |
| `10 + 5`        | 15     |
| `10 - 5`        | 5      |
| `10 * 5`        | 50     |
| `10 / 5`        | 2      |
| `10.5 + 2.25`   | 12.75  |
| `10 / 0`        | Error: Cannot divide by zero! |

## Project Structure

```
cli-calculator/
├── go.mod
├── main.go
└── README.md
```

- `main.go` — entry point containing the calculator loop and arithmetic logic
- `go.mod` — Go module definition (`module cli-calculator`)

## License

This project is open source and available under the MIT License.
