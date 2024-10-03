# Define variables
BINARY_NAME=load_balancer
BOOKS_API=books_api
STUDENTS_API=students_api
MOVIES_API=movies_api
PORT=8080
BOOKS_PORT=8081
STUDENTS_PORT=8082
MOVIES_PORT=8083

# Default target to build all binaries
all: build

# Build the load balancer binary
build: $(BINARY_NAME) $(BOOKS_API) $(STUDENTS_API) $(MOVIES_API)

$(BINARY_NAME):
	go build -o $(BINARY_NAME) algorithm/main.go

$(BOOKS_API):
	go build -o $(BOOKS_API) algorithm/books_api.go

$(STUDENTS_API):
	go build -o $(STUDENTS_API) algorithm/students_api.go

$(MOVIES_API):
	go build -o $(MOVIES_API) algorithm/movies_api.go

# Run the load balancer
run: build
	@echo "Starting the load balancer on port $(PORT)..."
	./$(BINARY_NAME) &

# Run the Books API
run-books:
	@echo "Starting the Books API on port $(BOOKS_PORT)..."
	./$(BOOKS_API)

# Run the Students API
run-students:
	@echo "Starting the Students API on port $(STUDENTS_PORT)..."
	./$(STUDENTS_API)

# Run the Movies API
run-movies:
	@echo "Starting the Movies API on port $(MOVIES_PORT)..."
	./$(MOVIES_API)

# Clean up binaries
clean:
	rm -f $(BINARY_NAME) $(BOOKS_API) $(STUDENTS_API) $(MOVIES_API)

# Help command
help:
	@echo "Makefile commands:"
	@echo "  make all            - Build all binaries"
	@echo "  make run            - Run the load balancer"
	@echo "  make run-books      - Run the Books API"
	@echo "  make run-students   - Run the Students API"
	@echo "  make run-movies     - Run the Movies API"
	@echo "  make clean          - Remove all binaries"
	@echo "  make help           - Show this help message"
