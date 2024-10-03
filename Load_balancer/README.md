# Load Balancer
This project is a basic implementation of a load balancer in Go using the Round Robin algorithm. The load balancer distributes incoming HTTP requests to a set of backend servers in a circular order. The project also has three REST APIs. The project consists of the following files:

- `main.go`: The main entry point of the load balancer.
- `api1.go`: Handles requests for API 1.
- `api2.go`: Handles requests for API 2.
- `api3.go`: Handles requests for API 3.

## Getting Started

### Prerequisites

- Go 1.16 or higher

### Installation

1. Clone the repository:
    ```sh
    git clone <name of the repo>
    cd Load_balancer
    ```

2. Build the project:
    ```sh
    go build
    ```

### Usage

Run the load balancer:
```sh
./load_balancer
```

### Project Structure

- `main.go`: Contains the main function and the load balancing logic.
- `api1.go`: Defines the endpoints and handlers for API 1.
- `api2.go`: Defines the endpoints and handlers for API 2.
- `api3.go`: Defines the endpoints and handlers for API 3.

### Running the Project

To run the load balancer, execute the following command:

```sh
go run main.go
```

This will start the load balancer, and it will begin handling requests for API 1, API 2, and API 3.

To run all three APIs, execute the following commands in separate terminal windows:

```sh
go run api1.go
go run api2.go
go run api3.go
```

### Contributing

Contributions are welcome! Please open an issue or submit a pull request.

### License

This project is licensed under the MIT License.