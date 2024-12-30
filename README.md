<div align=center>
<img 
  src="https://avatars.githubusercontent.com/u/91744653?s=200&v=4" 
  alt="Full Cycle Logo" 
  width="100" 
  style="vertical-align: middle; margin-right: 5px;" 
/>
  
# GoExpert Post-Graduation Full Cycle
</div>

This repository contains notes, exercises, and references for the **GoExpert** post-graduate program offered by [Full Cycle](https://goexpert.fullcycle.com.br/pos-goexpert/). The goal is to consolidate the advanced Go (Golang) knowledge acquired throughout the course and keep a record of practical examples for study purposes.

## Table of Contents

1. [About the Course](#about-the-course)  
2. [Repository Structure](#repository-structure)  
3. [Technologies and Topics Covered](#technologies-and-topics-covered)  
4. [How to Run](#how-to-run)  
5. [License](#license)  
6. [Contact](#contact)  

---

## About the Course

**GoExpert** is an advanced postgraduate program focused on Go (Golang), emphasizing software architecture, high scalability, and distributed systems development. During the course, you will learn:

- **Advanced Go Concepts:** Goroutines, Channels, Context, concurrency, and more.  
- **Microservices Architecture:** Building and communicating microservices with Go.  
- **DevOps & CI/CD:** Continuous integration and delivery pipelines, automation, and reliability practices.  
- **Design Patterns & Best Practices:** Software design patterns, clean code, and maintainability.  
- **Observability:** Distributed logging, metrics, and tracing for production environments.  
- **Tools & Frameworks:** Integration with popular Go libraries and frameworks.  

Learn more about the program at [Full Cycle](https://goexpert.fullcycle.com.br/pos-goexpert/).

---

## Repository Structure

Each module will be organized in its own folder, containing separate folders for lectures, which in turn have notes and exercises. An example structure could look like this:

```
.
├── module-01/
│   └── lectures/
│       ├── notes/
│       └── exercises/
├── module-02/
│   └── lectures/
│       ├── notes/
│       └── exercises/
├── module-03/
│   └── lectures/
│       ├── notes/
│       └── exercises/
├── docs/
│   └── references.md
└── README.md
```

- **module-XX**: Each folder represents a module from the GoExpert program.  
- **lectures**: Contains all material presented in the module’s lectures.  
  - **notes**: Your class notes, summaries, or additional explanations.  
  - **exercises**: Hands-on exercises to practice what you’ve learned.  
- **docs**: Additional documentation, reference materials, or helpful resources.  

---

## Technologies and Topics Covered

- **Go (Golang)**: Primary programming language, focusing on concurrency and performance.  
- **Containers & Orchestration**: Docker and Kubernetes for packaging and running applications.  
- **CI/CD Tools**: GitHub Actions, Jenkins, or other pipelines for continuous integration and deployment.  
- **Observability (Prometheus, Grafana, etc.)**: Monitoring services and gathering logs and metrics.  
- **Databases (SQL & NoSQL)**: Designing and querying relational and non-relational databases.  
- **Message Brokers**: Kafka, RabbitMQ, or similar tools for asynchronous communication.  

---

## How to Run

Depending on the lecture or module, you may have different instructions to run your notes and exercises. Generally, for a simple Go file:

```bash
go run main.go
```

To run all tests within a specific folder or module:

```bash
go test ./...
```

To build a binary:

```bash
go build -o my-binary .
```

Check the respective **notes** or **exercises** folder for any specific instructions related to each lecture or module.

---

## License

This repository is distributed under the [MIT License](LICENSE.md) (example). Feel free to use and modify the code according to the terms described. However, please do **not** share proprietary materials from the Full Cycle course that violate their rights.

---

## Contact

For questions or suggestions, feel free to reach out:

- **LinkedIn**: [Your LinkedIn Profile](#)
- **E-mail**: your.email@example.com  

---

*This README may evolve as the course progresses. New notes, exercises, and modules will be added as you advance through the GoExpert program.*
