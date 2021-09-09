# Upsee monitoring tool 📊

A system that records and visually displays the availability (uptime) of servers or services. In case of signal loss, the user is notified by mail or SMS.

## ✨ What we've got

- GO Stack: Chi, render and pq
- Postgres Database

## ⚡️ Requirements

- Golang (>= 1.16)
- Docker (>= 20.10.6)
- Docker-Compose (>= 1.27.4)
- Postgres
- make (>= GNU Make 4.2.1)

## 📦 Installation

1. Copy env example file `cp .env.example .env`
2. Update .env variables according to environment

## 🚀 Development

### Build Docker Container

```
make build
```

### Migrate database

```
make migrate
```
