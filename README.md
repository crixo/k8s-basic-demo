# Kubernetes basic demo

## Dependencies

- go 1.13

- kind v0.6.0 go1.13.4 darwin/amd64

- skaffold v1.7.0

- kustomize 3.5.4 (due to skaffold)

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

- Clean all resource within a namespace
```
k delete all --all -n ${NS}
```

- Delete the cluster
```
kind delete cluster --name k8s-basic-demo 
```