# Kubernetes basic demo

## Dependencies

- go 1.13

- kind v0.6.0 go1.13.4 darwin/amd64

- skaffold v1.7.0

- kustomize 3.5.4 (due to skaffold)

## Run

- Create the cluster
```
kind create cluster --config ./kind/cluster-config-single-node.yaml --name k8s-basic-demo 
```

- create namespace for the demo
```
k create ns k8s-basic-demo
kn k8s-basic-demo
```

- Run skaffold once
```
skaffold run
```

- Clean all resource within a namespace
```
k delete all --all -n ${NS}
```

- Delete the cluster
```
kind delete cluster --name k8s-basic-demo 
```