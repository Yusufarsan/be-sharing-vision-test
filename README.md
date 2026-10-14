# be-sharing-vision-test

Backend repository for the Sharing Vision test.

## Setup

1. Copy `.env.example` to `.env` and fill in your MySQL credentials.
2. Start the API:

```bash
go run ./cmd
```

The server defaults to port `8080` and auto-runs GORM migration for the `posts` table.
