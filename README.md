# Explorer Tour Microservice
## Service-Oriented Architecture

### Overview
* This project represents an upgrade in software design, transitioning from a monolithic application to a service-oriented architecture. We migrated the "tour" module from the monolithic Explorer application into a separate microservice.
* The original monolithic application was divided into several microservices, each with its own repository, programming language, and database. These microservices are orchestrated using Docker to improve development efficiency and testing.
* The "tour" module was extracted from the monolithic application and refactored into a dedicated microservice written in Go (Golang). This microservice handles all logic related to tours, enabling better scalability and maintenance.
* Additional Modules: Besides the "tour" module, other parts of the monolithic application were also converted into microservices. Document-oriented databases, such as MongoDB and Neo4j, were used to effectively manage various data needs in these modules.

### Technologies
* Microservice Implementation: Developed in Go (Golang) for managing the "tour" functionality.
* Database Management: PostgreSQL 
* Front-End Interface: Built with Angular using TypeScript, HTML, and CSS, and communicates with the back-end through RESTful APIs.
* Back-End Integration: C# (ASP.NET) is used to interface with the Go-based microservice and handle requests.
* Service Orchestration: Managed and deployed with Docker to simplify operations and ensure smooth integration.

### Transition to gRPC
gRPC has been implemented using Protocol Buffers (Protobuf) and .proto files. This setup ensures efficient data exchange and clear API documentation, enhancing communication speed and system performance.

### Application Setup
To set up the project locally using Docker, follow these steps:
* Clone the repository
* Navigate to the project directory
* Ensure Docker and Docker Compose are installed and running on your machine
* Build and start all services using Docker Compose: docker-compose up --build
* To stop all services: docker-compose down

### Contributors
* Kristina Zelić
* Ana Radovanović
* Milica Petrović
* Petar Kovačević

