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
  We noticed potential issue w/ skaffold on WSL not able to reach docker host using the standard env variable ```DOCKER_HOST=localhost:2375``` use ```DOCKER_HOST=tcp://127.0.0.1:2375``` instead.

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
    docker images | grep k8s-basic-demo
    ```
    *latest* tag has been added in order to use the existing yaml directly with kustomize w/o required skaffold to deploy into the cluster

    Kind [does not allow to load and use *latest* image](https://kind.sigs.k8s.io/docs/user/quick-start/#loading-an-image-into-your-cluster). Actually it's k8s itself that always try to get images with tag latest even if they are already present in the node.  
    todo-app-xxxx is in status *ImagePullBackOff* due to missing image with tag *latest*
    So let's push it to docker hun so king we'll get from there
    ```
    docker push crixo/k8s-basic-demo-todo-app
    ```

    after pushing and kind downloading from there the pod todo-app-xxxx is up&running

    You could also tag the image with a non-latest tag (eg. kind)
    ```
    docker tag crixo/k8s-basic-demo-todo-app:latest crixo/k8s-basic-demo-todo-app:kind 
    ```
    and load that image in kind
    ```
    kind load docker-image crixo/k8s-basic-demo-todo-app:kind --name k8s-basic-demo
    ```
    now you can deploy using a dedicate kustomize overlays
    ```
    kubectl apply -k k8s/overlays/kind/
    ```

    CleanUp images created by skaffold CI/CD
    ```
    docker images | grep crixo/k8s-basic-demo | awk '{print $3}' | xargs docker rmi -f
    ```

  - Check everything is working as expected
    ```
    k get all
    ```

  -  [Browse the app](http://localhost:30001/home) and noticed the 2 records previously inserted are still there.  
     They are living within the PV that survived across the multiple deployments and cleanups