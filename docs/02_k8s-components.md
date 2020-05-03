# Kubernetes components

https://kubernetes.io/docs/concepts/overview/components/

## Show k8s components through kind

Connect to kind k8s master node
```
CLUSTER_CONTROL_PLANE_NODE_NAME="k8s-basic-demo-control-plane"
docker exec -it $CLUSTER_CONTROL_PLANE_NODE_NAME> /bin/bash
```

From the container bash:
```
ps -ef | grep kube-apiserver
ps -ef | grep /usr/bin/kubelet
ps -ef | grep kube-scheduler
ps -f | grep kube-controller-manager #you can see the cni settings --cluster-cidr=192.168.0.0/12236
```
