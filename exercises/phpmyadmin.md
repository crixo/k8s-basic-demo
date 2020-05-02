# PhpMyadmin

Create deployment resource
```
k create deployment phpmyadmin --image=phpmyadmin/phpmyadmin:5.0.2-fpm-alpine --dry-run -o yaml > phpmyadmin-deployment.yaml
```

Set explicitly containers port where phpmyadmin runs
```
# get some help for yaml property
k explain deployment.spec.template.spec.containers.ports

vim pypmyadmin-deployment.yaml
```

Deploy the phpmyadmin deployment resource
```
k apply -f phpmyadmin-deployment.yaml
```

Expose phpmyadmin deployment via service
```
k expose deployment phpmyadmin --name phpmyadmin --port=5080 --target-port=80 --dry-run -o yaml > phpmyadmin-service.yaml
```

Deploy the phpmyadmin service resource
```
k apply -f phpmyadmin-service.yaml
```

## Use port-forwarding

Expose the service to the host using port-forwarding. Execute the following command into a new terminal and let it run
```
kubectl port-forward service/phpmyadmin 5080:5080
```

Open the local browser at http://localhost:5080
and use the following parameter to connect:

-  Server: mysql.k8s-basic-demo:3306
   SERVICE_NAME.NS_NAME:PORT namespace is not strictly required since service and deployment/pod are living into the same namespace.

- Username: root

- Password: root


## Use ingress

Create the [ingress resource](https://kubernetes.io/docs/concepts/services-networking/ingress/) for phpmyadmin

Open the local browser at http://phpmyadmin.192.168.1.15.xip.io 


