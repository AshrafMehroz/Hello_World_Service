# Virtual USB Hello World

This is a simple **gRPC Hello World service** designed for testing **NGINX ingress**. Built in just one day!

## Project Overview

### What This Project Provides:

- **Hello World gRPC Service**: A minimal gRPC service that responds with a greeting message and server info.
- **Kubernetes Setup**: Deploys 2+ pods across your Kubernetes cluster (supports auto-scaling and load balancing).
- **Scalable Infrastructure**: Uses **Horizontal Pod Autoscaler** (HPA) to automatically scale pods based on CPU utilization.

### Key Components:

1. **Dockerized gRPC Server**:  
   - The service is built using Go, with the `main.go` file handling the gRPC server and `proto/vusb.proto` defining the service.
   
2. **Deployment Configuration**:  
   - Kubernetes YAML files (`namespace.yaml`, `deployment.yaml`, `service.yaml`, `hpa.yaml`) for deploying the service with auto-scaling and load balancing across the cluster.

3. **Harbor Image**:  
   - The service is built into a Docker image that can be pushed to Harbor.

---

### Service Behavior:

- **gRPC Server**: The `hello-world` service responds to incoming requests with a simple greeting message that includes the device name and server info.
- **Auto-scaling**: The Kubernetes deployment will scale the number of pods between **2 and 10**, based on the CPU utilization (scaling up when CPU usage exceeds 70% and scaling down when it drops below 40%).
- **Round-robin Load Balancing**: The Kubernetes service exposes the app using **round-robin load balancing**, ensuring traffic is evenly distributed across the pods.

---

### **Important Notes**:

- **Replace Node Names**: In the `deployment.yaml`, you'll need to **replace the node names** in the affinity section for `DMSNV` and `DNV` nodes:
  
  ```yaml
  affinity:
    nodeAffinity:
      requiredDuringSchedulingIgnoredDuringExecution:
        nodeSelectorTerms:
          - matchExpressions:
              - key: kubernetes.io/hostname
                operator: In
                values:
                  - "dmsnv-node-1"  # Replace with actual DMSNV node names
                  - "dmsnv-node-2"
                  - "dnv-node-1"    # Replace with actual DNV node names
                  - "dnv-node-2"


