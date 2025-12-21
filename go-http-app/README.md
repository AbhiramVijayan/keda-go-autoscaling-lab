# KEDA Go HTTP Autoscaling Lab

This lab demonstrates a simple Go HTTP service scaled by KEDA using HTTP-based autoscaling.

## 1. Go HTTP Service

The Go app lives in this directory and exposes two endpoints:

- `GET /health` – basic health check
- `GET /work` – simulates work (500ms delay) and returns the pod hostname

Main file: `main.go`.

## 2. Building the Docker Image

The app is containerized with a multi-stage Dockerfile in `Dockerfile`.

To build the image locally (from this directory):

```bash
# Build image tagged go-http-app:1.0
docker build -t go-http-app:1.0 .
```

## 3. Using Minikube Docker Daemon

To use the local Docker image directly from Minikube (so you don’t need to push to a registry), point your shell to Minikube’s Docker daemon:\n
```bash
eval $(minikube docker-env)
```

After this, rebuild the image (so it is available inside Minikube’s Docker daemon):

```bash
docker build -t go-http-app:1.0 .
```

## 4. Installing / Applying KEDA

First, install KEDA into your cluster using Helm:

```bash
helm repo add keda https://kedacore.github.io/charts
helm repo update
helm install keda keda/keda --namespace keda --create-namespace
```

Then apply the KEDA HTTP add-on configuration provided in `keda-2.18.2.yaml` in this directory:

```bash
kubectl apply --server-side -f keda-2.18.2.yaml
```

Wait until all KEDA components and HTTP add-on pods are running before proceeding.

## 5. Kubernetes Manifests

Kubernetes manifests for the Go app are in `k8s/`:

- `deployment.yaml` – Deployment for the Go service (image `go-http-app:1.0`, `imagePullPolicy: IfNotPresent`).
- `service.yaml` – Service exposing the deployment.
- `HTTPScaledObject.yaml` – KEDA HTTP ScaledObject for autoscaling based on HTTP traffic.

Apply them from the `k8s/` directory:

```bash
cd k8s
kubectl apply --server-side -f .
```

## 6. Port-Forwarding the KEDA HTTP Interceptor

To send HTTP traffic through the KEDA HTTP interceptor proxy, port-forward the Service in the `keda` namespace:

```bash
kubectl port-forward svc/keda-add-ons-http-interceptor-proxy 8081:8080 -n keda
```

This forwards local port `8081` to the interceptor proxy on port `8080`.

## 7. Load Testing the /work Endpoint

With port-forwarding active, you can generate load against the `/work` endpoint through the KEDA HTTP interceptor. One way is to run a simple loop:

```bash
for i in {1..1000}; do
  curl -s -H "Host: go-http.local" http://localhost:8081/work > /dev/null &
done
wait
```

- `Host: go-http.local` – host header that KEDA HTTP uses to route traffic.
- `http://localhost:8081/work` – hits the `/work` endpoint via the interceptor proxy.

Alternatively, you can use the helper script in this directory:

```bash
chmod +x send_requests.sh
./send_requests.sh
```

- `send_requests.sh` – runs the same 1000-request loop against the KEDA HTTP proxy.

As load increases, KEDA should scale the Go HTTP deployment based on the configured HTTPScaledObject.

## 8. Cleanup

To remove the resources created by this lab:

```bash
# From k8s
kubectl delete -f .

# From go-http-app (this directory)
kubectl delete -f ../go-http-app/keda-2.18.2.yaml
```

This returns the cluster to a clean state.
