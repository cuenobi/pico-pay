# 🏦 PicoPay - Fintech Microservice System (Backend Only)

A simulated Fintech backend system designed to demonstrate production-level backend engineering skills in a microservices environment using Go, gRPC, and Kubernetes.

---

## 📌 Objectives

- Showcase real-world backend engineering skills (Go, GORM, gRPC, microservices)
- Build a system with clean architecture and observability in mind
- Practice infrastructure tools: Docker, Kubernetes
- Apply design patterns such as Domain-driven design, hexagonal architecture, and message-driven async flows
- Serve as a self-evaluation checklist for backend career growth

---

## ⚙️ Tech Stack

| Layer            | Technology                        |
|------------------|-----------------------------------|
| Language         | Golang                            |
| Transport        | gRPC over HTTP/2                  |
| Database         | PostgreSQL + GORM ORM             |
| Messaging        | NATS / Kafka (depending on setup) |
| Containerization | Docker                            |
| Orchestration    | Kubernetes (with Helm optional)   |
| Observability    | Zap Logger, context tracing       |
| CI/CD            | (optional: GitHub Actions / ArgoCD)|

---

## 🧱 System Overview

The system simulates core components of a digital payment infrastructure:

- **User Service** – manages user registration and KYC (mock)
- **Wallet Service** – handles wallet creation, balance management
- **Transaction Service** – processes money transfers between users
- **Notification Service** – asynchronously sends email/SMS notifications via messaging
- **Common Utils** – shared libraries for logging, database connection, message consumption

All services communicate through gRPC, with message queues used for background processing.

---

## 🚀 Getting Started

### 1. Prerequisites

- Go 1.21+
- Docker & Docker Compose
- [Optional] kubectl + Minikube / Kind (for K8s mode)
- [Optional] Protobuf compiler (`protoc`)

---

### 2. Running Locally (Docker Compose)

```bash
    make up         # or docker compose up --build
```
Services:

user-service → gRPC at :50051

wallet-service → gRPC at :50052

DB: PostgreSQL at :5432

---


### 3. Running on Kubernetes

```bash
    kubectl apply -f k8s/
```
Includes:

Deployments per service

Services (ClusterIP or LoadBalancer)

Secrets/configMaps for env vars

(Optional) Helm charts available in /charts/

---
