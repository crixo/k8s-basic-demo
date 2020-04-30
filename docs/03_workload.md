# Workload Management

## Use Case description

Describe the use case to solve.  
The use case solution will be used to walk you through the kubernetes basic resources.

## Execution

- create namespace for the demo
```
k create ns k8s-basic-demo
kn k8s-basic-demo
```

- deploy mysql
```
kubectl apply -f k8s/_manual/mysql-statefulset.yaml
```
  - StatefulSet, PesistentVolume and PesistentVolumeClaim
  ```
  k get statefulsets
  ```

  - Verify volume bindings
  ```
  k get pv,pvc
  ```
  - PV and PVC are sharing the same *fake* storage class that is used for linking the 2 resources.  
    An explicit label has been used to enforce the link.

  - With no PV explicitly defined, the default storage class will be used for PV provisioning.
  ```
  k get storageclasses standard -o yaml
  ```

- expose locally mysql
```
kubectl port-forward service/mysql 3306:3306
```
Browse mysql instance running within the cluster w/ your favorite mysql client

- Run skaffold once
```
skaffold run --tag='latest'
```

- Interact with the webapp

  - [Browse the app](http://localhost:30001/home) and add some records.

  - Verify records through the mysql local client


- Destroy everything but the PV
```
k delete deployments.apps todo-app 
k delete svc todo-app
k delete statefulsets.apps mysql 
k delete svc mysql 
k delete pvc mysql-mysql-0
```

- Check the PV and its Status
```
k get pv mysql-pv-volume
```

- Re-Deploy everything

  - First the StatefulSet
  ```
  kubectl apply -f k8s/_manual/mysql-statefulset.yaml
  ```

  - Notice the Status of the PVC and PV
  ```
  k get pv,pvc
  ```

  - Make the PV available to be re-claimed by the PVC
  ```
  kubectl patch pv mysql-pv-volume -p '{"spec":{"claimRef": null}}'
  ```

  - Check again the the Status of the PVC and PV
  ```
  k get pv,pvc
  ```

  - Now re-deploy the webapp using Kustomize directly via kubectl w/ the *-k* flag
  ```
  kubectl apply -k k8s/overlays/dev/
  ```
  Noticed the docker images built by skaffold
  ```
  docker images | grep kk8s-basic-demo
  ```
  *latest* tag has been added in order to use the existing yaml directly with kustomize w/o required skaffold to deploy into the cluster

  - Check everything is working as expected
  ```
  k get all
  ```

  -  [Browse the app](http://localhost:30001/home) and noticed the 2 records previously inserted are still there.  
     They are living within the PV that survived across the multiple deployments and cleanups