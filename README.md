# Lexi

[![Go Version](https://img.shields.io/badge/Go-1.22+-00ADD8?style=for-the-badge&logo=go)](https://golang.org/)
[![License](https://img.shields.io/badge/License-MIT-green.svg?style=for-the-badge)](LICENSE)
[![Architecture](https://img.shields.io/badge/Architecture-DDD%20%7C%20CQRS%20%7C%20Event%20Sourcing-blue?style=for-the-badge)](docs/)

**Lexi** is a high-performance, event-driven content enrichment platform built with Go. It leverages Domain-Driven Design (DDD), CQRS, and the Outbox Pattern to provide scalable, reliable content processing with AI-powered enrichment capabilities.

## Features

- ** Domain-Driven Design**: Rich domain models with aggregates, value objects, and business logic encapsulation
- ** CQRS Architecture**: Clear separation between command and query responsibilities
- ** Event-Driven**: Asynchronous processing using the Outbox Pattern with Kafka/NATS
- ** AI Integration**: Microsoft Phi integration for intelligent content enrichment
- ** Resilience**: Circuit breaker pattern for external API calls
- ** Scalable**: Microservice-ready architecture with clear service boundaries
- ** Testable**: Clean architecture enabling comprehensive unit and integration testing

## Architecture

Lexi follows a clean, layered architecture with clear separation of concerns:

```
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   REST API      │    │  Outbox Worker  │    │  Orchestrator   │
│   (cmd/api)     │    │(cmd/outbox-     │    │(cmd/orchestrator)│
│                 │    │ worker)         │    │                 │
└─────────────────┘    └─────────────────┘    └─────────────────┘
         │                       │                       │
         └───────────────────────┼───────────────────────┘
                                 │
                    ┌─────────────────┐
                    │   PostgreSQL    │
                    │   + Outbox      │
                    └─────────────────┘
                                 │
                    ┌─────────────────┐
                    │  Kafka/NATS     │
                    │  (Messaging)    │
                    └─────────────────┘
                                 │
                    ┌─────────────────┐
                    │  Microsoft Phi  │
                    │  (AI Service)   │
                    └─────────────────┘
```

### Core Components

- **`cmd/api`**: REST API server with Chi router and middleware
- **`cmd/outbox-worker`**: Processes outbox events and publishes to message broker
- **`cmd/orchestrator`**: Manages saga workflows and event coordination
- **`internal/domain`**: Business logic, aggregates, and value objects
- **`internal/application`**: Command handlers and application services
- **`internal/infra`**: Database, messaging, and external service adapters
- **`pkg/shared`**: Shared types, events, and DTOs

## 🎯 Use Case: Content Enrichment Workflow

1. **Content Submission**: User submits content (title + body)
2. **State Transition**: Content moves from `draft` → `enrichment_requested`
3. **AI Processing**: Microsoft Phi generates metadata (summaries, SEO titles, tone variations)
4. **Event Processing**: Asynchronous enrichment via event-driven architecture
5. **State Management**: Content transitions through `enriched` → `published` states
6. **Error Handling**: Robust failure handling with fallback and orchestration logic

### Content Lifecycle States

```go
draft → enrichment_requested → enriched → published
  ↓
failed (with retry logic)
```

## 🛠️ Technology Stack

- **Language**: Go 1.22+
- **Database**: PostgreSQL with `sqlx` for type-safe SQL
- **Messaging**: Kafka or NATS for event streaming
- **HTTP Router**: Chi for lightweight, idiomatic HTTP routing
- **Circuit Breaker**: `sony/gobreaker` for resilience
- **Database Access**: Manual SQL with domain-first approach (no ORM)
- **Containerization**: Docker with docker-compose for local development

## 📦 Project Structure

```
lexi/
├── cmd/                    # Application entry points
│   ├── api/               # REST API server
│   ├── orchestrator/      # Saga orchestration service
│   └── outbox-worker/     # Outbox pattern worker
├── internal/              # Private application code
│   ├── application/       # Application layer
│   │   ├── commands/      # Command handlers
│   │   └── events/        # Event handlers
│   ├── domain/            # Domain layer
│   │   ├── content/       # Content aggregate
│   │   └── enrichment/    # Enrichment domain
│   └── infra/             # Infrastructure layer
│       ├── messaging/     # Message broker adapters
│       ├── phi/           # Microsoft Phi integration
│       └── postgres/      # Database adapters
├── pkg/                   # Public packages
│   └── shared/            # Shared types and utilities
├── docs/                  # Architecture documentation
├── docker-compose.yaml    # Local development setup
└── go.mod                 # Go module definition
```

## 🚀 Quick Start

### Prerequisites

- Go 1.22 or later
- Docker and Docker Compose
- PostgreSQL (or use Docker)
- Kafka or NATS (or use Docker)

### Local Development

1. **Clone the repository**
   ```bash
   git clone https://github.com/tolgafiratoglu/lexi.git
   cd lexi
   ```

2. **Start dependencies**
   ```bash
   docker-compose up -d
   ```

3. **Install dependencies**
   ```bash
   go mod download
   ```

4. **Run the application**
   ```bash
   # Start the API server
   go run cmd/api/main.go
   
   # Start the outbox worker (in another terminal)
   go run cmd/outbox-worker/main.go
   
   # Start the orchestrator (in another terminal)
   go run cmd/orchestrator/main.go
   ```

5. **Test the API**
   ```bash
   curl -X POST http://localhost:8080/api/v1/content \
     -H "Content-Type: application/json" \
     -d '{
       "author_id": 1,
       "title": "Sample Article",
       "body": "This is a sample article for enrichment."
     }'
   ```

## 📚 API Documentation

### Content Management

#### Create Content
```http
POST /api/v1/content
Content-Type: application/json

{
  "author_id": 1,
  "title": "Your Article Title",
  "body": "Your article content here..."
}
```

#### Get Content
```http
GET /api/v1/content/{id}
```

#### List Content
```http
GET /api/v1/content?author_id=1&status=enriched
```

## 🧪 Testing

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run integration tests
go test -tags=integration ./...
```

## 🏗️ Development Guidelines

### Domain-Driven Design Principles

- **Aggregates**: Each aggregate is a consistency boundary
- **Value Objects**: Immutable objects representing domain concepts
- **Domain Events**: Capture business-relevant state changes
- **Repository Pattern**: Abstract data access behind domain interfaces

### Code Organization

- **Commands**: Represent user intentions and trigger state changes
- **Queries**: Read-only operations optimized for specific use cases
- **Events**: Represent things that have happened in the domain
- **Sagas**: Coordinate long-running business processes

### Error Handling

- Use domain-specific error types
- Implement circuit breakers for external services
- Provide meaningful error messages and status codes
- Log errors with appropriate context

## 🤝 Contributing

We welcome contributions! Please see our [Contributing Guidelines](CONTRIBUTING.md) for details.

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- [Microsoft Phi](https://github.com/microsoft/phi-3) for AI capabilities
- [Chi Router](https://github.com/go-chi/chi) for HTTP routing
- [Sony GoBreaker](https://github.com/sony/gobreaker) for circuit breaker pattern
- The Go community for excellent libraries and patterns

## 📞 Support

- 📧 Email: [your-email@example.com]
- 🐛 Issues: [GitHub Issues](https://github.com/tolgafiratoglu/lexi/issues)
- 💬 Discussions: [GitHub Discussions](https://github.com/tolgafiratoglu/lexi/discussions)

---

**Built with ❤️ using Go and Domain-Driven Design principles**
