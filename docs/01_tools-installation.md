# Tools Installation

If you have a WSL working fine, install all the rest of the tools on that, otherwise use the binaries/installer targeted to your OS. Al the tools should be available natively for the major OS also including windows 10 but I didn't test it all of them on that one, while on linux/wSL or MacOS I did it. I'm going to use MacOS during the class.

I use *grep* command a lot. That works only w/ bash, if you use powershell [search for an equivalent command](https://stackoverflow.com/questions/15199321/powershell-equivalent-to-grep-f).

- [Docker Desktop](https://www.docker.com/products/docker-desktop)  
You have to create a docker hub account.  **Do not** enable built-in kubernetes local cluster: we are going to use kind instead.

- [WSL](https://itnext.io/setting-up-the-kubernetes-tooling-on-windows-10-wsl-d852ddc6699c) - (Win user only) - Optional  
Ubuntu16.04 should be fine, but I suggest you to start with 18.04 instead.  
Someone had issue installing docker-compose via python. [Installing it via curl](https://www.digitalocean.com/community/tutorials/how-to-install-docker-compose-on-ubuntu-16-04) works just fine.

- [kind](https://kind.sigs.k8s.io/docs/user/quick-start).  
Version v0.7.0 should be fine but all the code has been tested against version v0.6.0 so install that one instead.  
After installing kind application, try to create and delete a cluster to ensure everything is working. Doing so your are going to download the required docker images containing the k8s version related to kind version you just installed. **Please make sure you download the kind docker image prior to the class due to the image size**.  
Launching the following command from terminal
```
docker images | grep kindest/node
```
you should get something like that
```
kindest/node    v1.16.3    066d19ae6707    5 months ago    1.22GB
```

- [kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/#install-kubectl-on-linux)  
WSL tutorial already included the kubectl installation step.  In principle you should have a kubectl version matching the k8s version running on the cluster you are targeting to. In our case should be 1.16 but any version above should be fine too.  
The k8s official documentation is the [best reference for kubectl](https://kubernetes.io/docs/reference/kubectl/cheatsheet/).  
In order to understand the version strategy around the kubernetes ecosystem you could use [this article](https://medium.com/@cristiano.deg/pinning-k8s-subcomponents-with-go-mod-1ad087731f83) and the links within it pointing to the github issue of kubernetes subcomponents such as kubectl or client-go.

- [kustomize](https://github.com/kubernetes-sigs/kustomize) 3.5.4  
kustomize lets you customize raw, template-free YAML files for multiple purposes, leaving the original YAML untouched and usable as is.  
Since v1.14 the kustomize build system has been included in kubectl. Installing kustomize explicitly is required only if you'd like to have a full class experience using skaffold.

- [skaffold](https://skaffold.dev/docs/install/) v1.7.0  
Skaffold is a command line tool that facilitates continuous development for Kubernetes-native applications. Skaffold handles the workflow for building, pushing, and deploying your application, and provides building blocks for creating CI/CD pipelines. This enables you to focus on iterating on your application locally while Skaffold continuously deploys to your local or remote Kubernetes cluster.

