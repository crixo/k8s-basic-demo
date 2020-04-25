# Configuration

## kubectl

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
```
source ~/kube-ps1.sh
PS1='[\w $(kube_ps1)]\$ '
```
