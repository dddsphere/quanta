A reference Todo List app that serves as a showcase for implementing lightweight microservices in Golang

Overview
----------
Quanta is a Todo List application that prioritizes simplicity and a streamlined approach to building services. It offers insights into constructing services that embrace CQRS, tailored to the needs of smaller, more focused systems. The codebase of Quanta is well-organized and designed to facilitate the development of CQRS-based solutions, while also allowing for the incorporation of traditional REST architecture when necessary, without dogmatically discarding it for specific use cases.

Key Features
----------
1. Lightweight Microservices: Quanta focuses on simplicity and an organized, smaller codebase compared to the more extensive approach of Martello. It provides a clear and concise structure for building microservices with a strong emphasis on CQRS principles without strongly advancing other DDD concepts.

2. Lightweight Microservices: Quanta emphasizes simplicity and an organized, compact codebase compared to the more extensive approach of Maetello. It provides a clear and concise structure for building microservices with a strong emphasis on CQRS principles, without aggressively pushing for other DDD concepts. 

3. HTTP and gRPC Communication: HTTP is primarily used for client-service communication, utilizing either CQRS (GET and POST) or REST when it proves to be more convenient. On the other hand, gRPC is employed for efficient and performant communication between services within the microservices architecture.

4. Optional NATS Integration: Quanta includes a simple event dispatcher as a reference implementation, utilizing NATS for hypothetical microservices within the orchestration. However, the use of NATS is not required for the core functionality of the app. It serves as a demonstration of event-driven architecture, where other hypothetical services, not implemented here, can consume these events. (*)

(*) Please note that Quanta serves as a reference implementation for a future generator. A flag in the generator will enable the inclusion of the necessary functionality to operate with NATS, allowing developers to customize and adapt the event bus integration as needed.

Notes
-----
Please note that the exploration and implementation showcased in Quanta does not indicate the discontinuation of efforts initiated in Maetello and TS. On the contrary, Quanta serves as a complementary reference, offering a different perspective and approach to building microservices. The development of the other two projects, [Martello](https://github.com/dddsphere/martello) and [TopSpin](https://github.com/dddsphere/topspin), continues concurrently, each addressing specific aspects and providing a comprehensive toolkit for developers. Quanta expands the options and insights available, contributing to a broader understanding and application of microservices architecture.

License
----------
Quanta is distributed under the MIT License. 

