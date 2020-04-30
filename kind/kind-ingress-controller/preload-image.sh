#!/bin/sh
read -r -p "cluster name(k8s-basic-demo-ingress-ctrl): " CLUSTER_NAME
if [ -z "$CLUSTER_NAME" ]; then 
    # echo "CLUSTER_NAME is mandatory"
    # exit
    CLUSTER_NAME="k8s-basic-demo-ingress-ctrl"
fi

declare -a arr=("crixo/k8s-basic-demo-ingress-ctrl-todo-app:v0.0.0" 
                "crixo/k8s-basic-demo-ingress-ctrl-informer:v0.0.0"
                "crixo/k8s-basic-demo-ingress-ctrl-webhook-server:v0.0.0"
                "bitnami/kubectl:1.16"
                )

for image in "${arr[@]}"
do
   echo "IMAGE: $image"
   docker pull $image
   #command=(kind load docker-image $image --name $CLUSTER_NAME --nodes='k8s-basic-demo-ingress-ctrl-worker,k8s-basic-demo-ingress-ctrl-worker2')
   #"${command[@]}"
   kind load docker-image $image --name $CLUSTER_NAME --nodes='k8s-basic-demo-ingress-ctrl-worker,k8s-basic-demo-ingress-ctrl-worker2'
done
