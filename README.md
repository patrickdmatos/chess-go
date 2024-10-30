# Chess Game

A multiplayer chess game built with Go, featuring a backend structure using the Fiber framework and GORM for database interactions.

## Table of Contents
- Overview
- Features
- Project Structure
- Getting Started
- Prerequisites
- Installation
- Usage
- Contributing
- Thanks

## Overview
This chess game allows players to compete in a multiplayer environment. The backend is designed for performance and scalability, making use of the Fiber web framework and GORM for ORM support with SQLite.

## Features
| Feature           | Summary                                                       |
|-------------------|---------------------------------------------------------------|
| ⚙️ Architecture    | Built using Go with Fiber for web handling and GORM for ORM. |
| 🔌 Integrations    | SQLite for database management and environment configurations. |
| 🧩 Modularity      | Structured code with separation of concerns across modules.   |
| ⚡️ Performance     | Efficient request handling with Fiber and optimized database queries. |

## Project Structure
chess_game/
├── config/
│   └── config.go              # Configuration settings and environment variables
├── controllers/
│   └── player.controller.go     # Handle player-related requests
│   └── match.controller.go      # Handle game-related requests
├── database/
│   └── database.go              # Database configuration file
├── models/
│   └── player.go               # Player model definition and methods
│   └── match.go               # Friend model definition and methods
├── repositories/
│   └── playerRepository.go      # Database interactions for players
│   └── gameRepository.go        # Database interactions for games
├── routes/
│   └── routes.go                # Route definitions for the application
├── services/
│   └── player.service.go         Business logic for players
│   └── match.service.go         # Business logic for games
├── main.go                      # Entry point for the application
├── go.mod                       # Go module definition
├── go.sum                       # Go module dependencies
└── README.md                    # Project documentation


## Getting Started
### Prerequisites
Before getting started, ensure you have the following:
- Go 1.19 or later
- SQLite installed

### Installation
1. Clone the repository:
   ```bash
   git clone https://github.com/patrickdmatos/chess-go.git
   cd chess-go
   
go mod tidy

go run main.go

Contributing
💬 Join the Discussions: Share your insights, provide feedback, or ask questions.

🐛 Report Issues: Submit bugs found or log feature requests.

💡 Submit Pull Requests: Review open PRs, and submit your own PRs.

Thanks
