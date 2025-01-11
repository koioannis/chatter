# Chatter

This project is a full-stack chat application built as a sandbox for exploring and learning various technologies and architectural patterns. The core stack includes Go, HTMX, and Apache Kafka as a message broker.

The application is designed as a single-binary where Kafka is utilized not out of necessity but for educational exploration.

## Overview

### Core Components

- Backend: [Echo](https://echo.labstack.com/) for routing, and [FX](https://github.com/uber-go/fx) for D.I.
- Frontend: [HTMX](https://htmx.org/), [Hyperscript](https://hyperscript.org/), [Tailwind](https://tailwindcss.com/) and [Templ](https://templ.guide/)
- Realtime Messanging: [Websockets](https://htmx.org/extensions/web-sockets/) and [Kafka](https://kafka.apache.org/).

### Getting Started

1. `git clone https://github.com/koioannis/chatter`
2. `cd chatter`
3. `docker compose up`

### Develop Locally

To launch the application locally you need:

- Compile the `styles.css` file using the Tailwind CLI
- Generate the `.go` files from the `.templ` files with Templ CLI.

For this process, you will require the following tools:

1. [Templ CLI](https://templ.guide/quick-start/installation)
2. [Tailwind CLI](https://tailwindcss.com/docs/installation) (I use the standalone binary)
3. You can _optionally_ use [Air](https://github.com/cosmtrek/air) for hot-reloading.

Afterwards, you can use the provided `Makefile` to either run the application, or generate the required files.

## Disclaimer

This project is a personal exploration into the technologies mentioned above. It is not intended to serve as a model of best practices, nor is it optimized for production environments. It's a snapshot of my learning process, shared in the spirit of open learning and development.
