# Configuration

## kubectl (in WSL)

- Pin a specific kubectl version matching w/ your cluster
```
vim ~/.bashrc

alias kubectl=/usr/local/Cellar/kubernetes-cli/1.17.0/bin/kubectl
```

- Open bash configuration file
```
vim ~/.bash_profile
# or
# vim ~/.bashrc
```

- Configure kubectl autocompletion
https://kubernetes.io/docs/reference/kubectl/cheatsheet/#kubectl-autocomplete
```
source <(kubectl completion bash)
```

- Configure some alias
```
alias k='kubectl'
complete -F __start_kubectl k
alias kn='k config set-context --current --namespace '
alias kcn='k config view --minify | grep namespace'
```

- Configure cluster/context in bash prompt
Download kube-ps1.sh
```
curl -Lo kube-ps1.sh https://raw.githubusercontent.com/jonmosco/kube-ps1/master/kube-ps1.sh
chmod +x kube-ps1.sh
mv kube-ps1.sh ~/kube-ps1.sh
```

Configure kube-ps1 adding the following lines in ```~/.bashrc```
```
source ~/kube-ps1.sh
PS1='[\w $(kube_ps1)]\$ '
```

Reload ~/.bashrc
```
source ~/.bashrc
```

## kubectl (Win10 only)

Mirror and use the kubeconfig from WSL to Win10 to connect to the kind cluster create in WSL.

- From WSL terminal
```
WIN_USER_NAME="your windoews user name"
cat ~/.kube/config >> /c/Users/$WIN_USER_NAME/.kube/config
```

- [install kubectl in Win10](https://kubernetes.io/docs/tasks/tools/install-kubectl/#install-kubectl-on-windows)

- From Win10 terminal
```
kubectl config set-context --current --namespace k8s-basic-demo
kubectl config view --minify
```
