# Local Time Go Server

A simple Go HTTP server that returns the current time in Vancouver (PST/PDT) timezone.

## Run Locally

```bash
go run main.go
```

The server will start on port 1599.

## Use

Open your browser or use curl:

```bash
curl http://localhost:1599
```

Output example:

```
Server Local Time: Monday, December 23, 2025 at 2:30:45 PM PST
```

## Run with Docker

Build:

```bash
docker build -t localtime .
```

Run:

```bash
docker run -p 1599:1599 localtime
```

