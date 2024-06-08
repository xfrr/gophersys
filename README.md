<p align="center">
  <a href="" rel="noopener">
 <img width=100px height=100px src="web/static/assets/logo.png" alt="Project logo"></a>
</p>

<h3 align="center">Gophersys</h3>

<div align="center">

[![Status](https://img.shields.io/badge/status-active-success.svg)]()
[![GitHub Issues](https://img.shields.io/github/issues/xfrr/gophersys.svg)]()
[![GitHub Pull Requests](https://img.shields.io/github/issues-pr/xfrr/gophersys.svg)]()
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](/LICENSE)

</div>

---

<p align="center"> A simple CRUD application built with Go, Chi, and MongoDB to manage Gophers.
    <br>
</p>


## 📝 Table of Contents

- [📝 Table of Contents](#-table-of-contents)
- [🏁 Getting Started ](#-getting-started-)
  - [Prerequisites](#prerequisites)
    - [Development Tools](#development-tools)
    - [Deployment Tools](#deployment-tools)
  - [Installing](#installing)
- [🎈 Usage ](#-usage-)
  - [Running the application locally](#running-the-application-locally)
  - [Running the application using Docker](#running-the-application-using-docker)
  - [Accessing the application](#accessing-the-application)
  - [Cleaning the application](#cleaning-the-application)
- [🚀 Deployment ](#-deployment-)
  - [Deploy using Docker](#deploy-using-docker)
  - [Deploy using Kubernetes and Helm](#deploy-using-kubernetes-and-helm)
- [📝 License ](#-license-)


## 🏁 Getting Started <a name = "getting_started"></a>

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.


### Prerequisites

- [Go >= 1.22](https://golang.org/dl/) - *Programming Language*
- [MongoDB >= 8.0](https://docs.mongodb.com/manual/installation/) - *Database for persistence*


#### Development Tools
- [Go Air](https://github.com/air-verse/air) - *Live Reloading*
- [Protoc >= 3.0](https://google.github.io/proto-lens/installing-protoc.html) - *Protocol Buffers Compiler*


#### Deployment Tools
- [Docker 24.0 or higher](https://www.docker.com/) - *Containerization*
- [Kubernetes 1.27 or higher](https://kubernetes.io/) - *Container Orchestration*
- [Helm 3.0 or higher](https://helm.sh/) - *Kubernetes Package Manager*


### Installing

1. Clone the repository

    ```bash
    git clone https://github.com/xfrr/gophersys.git
    ```

2. Change directory to the project folder and setup the environment

    ```bash
    cd gophersys && make setup
    ```
    > This will install all the dependencies required for development.


## 🎈 Usage <a name="usage"></a>


### Running the application locally

1. Start the application using one of the following commands:
    ```bash
    # using go run
    make run

    # using https://github.com/air-verse/air
    # for live reloading
    make run-air
    ```


### Running the application using Docker

Start the application using the following command:
```bash
  make run-docker
```


### Accessing the application
Navigate to [http://localhost:8080](http://localhost:8080) in your browser


### Cleaning the application
To clean up all development resources, run the following command:
```bash
make purge
```


## 🚀 Deployment <a name="deployment"></a>

### Deploy using Docker

1. Build the Docker image
    ```bash
    make build-docker
    ```
2. Run the Docker container using docker-compose
    ```bash
    make run-docker
    ```
3. Navigate to [http://localhost:8080](http://localhost:8080) in your browser

### Deploy using Kubernetes and Helm

1. Install the application using Helm
    ```bash
    make install-helm
    ```

2. Uninstall the application using Helm
    ```bash
    make uninstall-helm
    ```


## 📝 License <a name="license"></a>

This project is [MIT](LICENSE) licensed.