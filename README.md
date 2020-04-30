# Kubernetes basic demo

## Dependencies

-  [Docker Desktop](https://www.docker.com/products/docker-desktop)  

- [kubectl](https://github.com/kubernetes/kubectl) v1.16.3

- [kind](https://kind.sigs.k8s.io/) v0.6.0 go1.13.4 darwin/amd64 -  kindest/node v1.16.3 

- [kustomize](https://github.com/kubernetes-sigs/kustomize) 3.5.4 (required by skaffold)

## DevDependencies

- [skaffold](https://skaffold.dev/docs/install/) v1.7.0

- [kube-ps1](https://github.com/jonmosco/kube-ps1)

- [golang](https://golang.org/doc/install/source) 1.13 (optional)

## Run

- Create the cluster
```
cd kind
kind create cluster --config cluster-config-single-node.yaml --name k8s-basic-demo 
cd ..
```

- create namespace for the demo
```
k create ns k8s-basic-demo
kn k8s-basic-demo
```

- deploy mysql
```
kubectl apply -f k8s/_manual/mysql-statefulset.yaml
```

- expose locally mysql
```
kubectl port-forward service/mysql 3306:3306
```

- Run skaffold once
```
skaffold run
```

- Check deployment status
```
k get all
```

- Delete the cluster
```
kind delete cluster --name k8s-basic-demo 
```

## Topics

- [Tools installation](docs/01_tools-installation.md)

- [Kubernetes Architecture Overview](https://kubernetes.io/docs/concepts/overview/components/)

- [Configuration](docs/02_config.md)

- [Workload management](docs/03_workload.md)

- [Ingress Controller](https://kind.sigs.k8s.io/docs/user/ingress/)