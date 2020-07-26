# kubectl-which-node

# Installation
```
# Get the binary
wget https://raw.githubusercontent.com/corneredrat/kubectl-which-node/master/bin/kubectl-which-node

# Change permissions
chmod +x kubectl-which-node

# Add to path
sudo mv ./kubectl-which-node /usr/bin/kubectl-which-node
```

# Build from source
```
# Get code
git clone https://github.com/corneredrat/kubectl-which-node.git

# Run build script
./build
```

# Working

This command display the kubernetes nodes in which a workload controller is deployed in. Workload controller can be `replicaset`, `daemonset` , `deployment` or any other standard controller.

All shortforms that kubernetes api server support , such as `rs` for `replicaset` , `ds` for `daemonset` is supported by the command.

### Examples:

To display which nodes a daemonset is deployed in:
