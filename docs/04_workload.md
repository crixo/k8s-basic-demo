# Workload Management

## Use Case description

Describe the use case to solve.  
The use case solution will be used to walk you through the kubernetes basic resources.

## Execution

- Create the cluster
  ```
  cd kind
  kind create cluster --config cluster-config-single-node.yaml --name k8s-basic-demo 
  cd ..
  ```

- Create namespace for the demo
  ```
  k create ns k8s-basic-demo
  kn k8s-basic-demo
  ```

- Deploy mysql
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

- Expose locally mysql
  ```
  kubectl port-forward service/mysql 3306:3306
  ```
  Browse mysql instance running within the cluster w/ your favorite mysql client.  
  Instead of install a client version on your machine, use docker to expose a web-browser version for a mysql client.  
  Open a new terminal and run the following command:
  ```
  cd docker
  docker-compose up
  ```

  If you decide to follow the docker-way, use the following port-forwarding instead the one above.  
  Open a new terminal and run the following command
  ```
  HOST_IP=$(ipconfig getifaddr en0)
  kubectl port-forward --address $HOST_IP service/mysql 3306:3306
  ```

  > Exercise: Try to deploy phpmyadmin into the k8s cluster connecting to mysql that runs in the cluster as well so you do not need port-forwarding anymore.


- Run skaffold once
  ```
  skaffold run --tag='latest'
  ```

  or keep it running. Skaffold will watch code changes and will kick/queue build&deploy according to your changes.
  ```
  skaffold run dev
  ```
  stopping the watcher will also clean the deployed workload.

  If you decide to use the container to build the golang application, remove the vendor folder locally otherwise if you add new packages won't be downloaded during the multi-stage build within the container.  
  To speed up the skaffold image rebuild, you should vendoring your dependency locally so these will be reuse/loaded during the container build.

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