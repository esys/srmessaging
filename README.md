# SR Messenging

## Objective

The purpose of this exercise is to implement a messaging service meeting the following
requirements :

- If client A sends a message M to client B and B is known by the service, the service
  should deliver M to B
- The client should be implemented in a web browser (very basic HTML + JS)
- The implementation technology is open: any language or middleware (databases,
  messaging system…) can be used (Golang and Python are preferred though). However,
  if a third-party system is used, the implementation must include a standalone version.
  To sum up, this problem can be reduced to a simple directed/targeted chat service

## Prerequisites

This project use :

- golang
- docker-compose

## Architecture

The message distribution is done using Kafka :

- User A connects through browser using websocket and is given a unique ID
- User A sent message to another user B
- Received message is sent to Kafka
- Message is consumed and routed to user B

## Building

```bash
make clean build
```

## Testing

```bash
make test
```

## Running

Run docker-compose

```bash
make run
```

## CLI usage

Instead of `make run`, you can build or install the go binary and then use the provided command line interface.

```bash
Run the server

Usage:
  srmessaging run [flags]

Flags:
  -b, --broker string     Message broker endpoint (default "localhost:9092")
  -e, --endpoint string   HTTP endpoint to listen to (default ":8080")
  -h, --help              help for run
```

# How to test the app

- run the backend as seen above
- open two tabs for user A & B with the test webpage (see project srmessenging-ui)
- in tab A : fill in the "Send To" field with the user B id
- in tab A : enter a message and press send
- in tab B : received message is displayed at the bottom
