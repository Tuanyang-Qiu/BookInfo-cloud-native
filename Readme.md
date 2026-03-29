# BookInfo - Cloud-Native Hands-On Project

## Project Overview

This repository hosts a comprehensive, cloud-native hands-on project focused on Go web application development, microservices architecture, Docker, and Kubernetes. The project is structured into six progressive phases, guiding the evolution from a simple monolith to a fully automated microservices ecosystem:

### Project Phases

- ✅ **Phase 1: Monolithic Application Development** – Developed "BookInfo," a simple monolithic web application using Go.

- 🚧 **Phase 2: Containerization & Orchestration** – Containerized the application with Docker and deployed it to Kubernetes. Utilized declarative API orchestration tools, including Helm, Kustomize, and Custom Resource Definitions (CRDs).

- 🚧 **Phase 3: Microservices Decomposition** – Refactored the monolithic BookInfo application into a loosely coupled microservices architecture.

- 🚧 **Phase 4: Telemetry & Monitoring** – Implemented distributed tracing for the microservices using Jaeger, and integrated Prometheus Operator for metrics collection.

- 🚧 **Phase 5: Service Mesh Integration** – Introduced Istio to build a service mesh, focusing on advanced traffic management and enhanced system observability.

- 🚧 **Phase 6: CI/CD & Automated Operations** – Established a complete Continuous Integration and Continuous Deployment (CI/CD) pipeline to enable automated deployment and streamlined DevOps workflows.

---

## Startup

### Prerequisites

- Go 1.25.5 or higher

### Running the Application

Start the server
```bash
go run main.go serve
```
### Accessing the Application
**Home Page**: http://localhost:8080/
