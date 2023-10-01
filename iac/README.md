# Zot Developer IaC

## 🌟 Overview

This IaC project leverages the power of Pulumi and KinD (Kubernetes in Docker) to provision Kubernetes clusters swiftly and reliably. In a single command, you can go from zero to having a fully functional Kubernetes cluster right on your local machine or cloud environment, complete with configurable Docker volumes.

### 📊 What It Solves

- **Rapid Prototyping**: Ideal for developers who need to quickly spin up a Kubernetes cluster.
- **Local Development**: Perfect for testing Kubernetes deployments locally.
- **Automation**: Easily integrate into CI/CD pipelines.

## 🚀 Getting Started

These instructions will help you set up a KinD cluster using this IaC project.

### 📋 Prerequisites

- [Go](https://golang.org/dl/) - v1.16+
- [Pulumi](https://www.pulumi.com/docs/get-started/install/)
- [Docker](https://www.docker.com/products/docker-desktop)
- [kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/)

### 🖥️ Installation

1. **Clone the Repository**

    ```bash
    git clone https://github.com/emporous/uor-zot.git
    cd uor-zot/iac
    ```

2. **Download Go Deps**

    ```bash
    go get -u ./...
    go mod download
    go mod tidy
    ```

2. **Initialize Pulumi**

    ```bash
    pulumi stack init dev
    ```

3. **Run Pulumi**

    ```bash
    pulumi up
    ```

4. **Check Your Cluster**

    ```bash
    kubectl get nodes
    ```

## 🛠️ How It Works

### 📚 Components

1. **Pulumi**: Orchestrates the provisioning.
2. **KinD**: Creates a local Kubernetes cluster using Docker containers.
3. **Go**: The IaC logic is written in Go, leveraging Pulumi's Go SDK.

### 🔨 Code Structure

- `kind/kind.go`: Core logic for KinD cluster creation.
- `kind/volumes.go`: Manages Docker volumes for the KinD cluster.
- `helper/helper.go`: Utility functions and types.

### 📈 Workflow

1. **Initialize Configuration**: Sets up Pulumi context and configuration.
2. **Load Cluster Config**: Reads the desired Kind cluster setup.
3. **Manage Docker Volumes**: Checks, creates, or deletes Docker volumes as needed.
4. **Create/Delete KinD Cluster**: Uses Pulumi's `local.Command` to run the KinD CLI commands.

## 📜 Customizing the Cluster

For now, the code is set up to create a cluster with a default configuration. Customizations will be added in future releases.

## 🤝 Contributing

Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct and the process for submitting pull requests.

## 📄 License

This project is licensed under the Apache 2.0 License - see the [LICENSE.md](../LICENSE) file for details.

## 🙏 Acknowledgments

- The Pulumi Team for their fantastic IaC tool.
- Kubernetes for revolutionizing the way we think about container orchestration.
- [The ReadME Project](https://www.readme.com/the-readme-project) for README best practices.

-------------------------------------------------------

## Local Development

```bash
export GOFLAGS='-replace=github.com/emporous/uor-zot=.'
```
