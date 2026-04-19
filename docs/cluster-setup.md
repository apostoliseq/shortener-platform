# Cluster Setup

## Prerequisites (already installed)
- Ubuntu Server 24.04
- Docker (installed but not used by Kubernetes — containerd is the CRI)
- containerd
- kubeadm, kubelet, kubectl
- Helm v4.1.4

## Cluster Reset
If the cluster needs to be reset (e.g. IP change, certificate issues):

```bash
sudo kubeadm reset --cri-socket unix:///var/run/containerd/containerd.sock

# Clean up CNI and network rules after reset
sudo rm -rf /etc/cni/net.d
sudo iptables -F && sudo iptables -t nat -F && sudo iptables -t mangle -F && sudo iptables -X
```

## Cluster Initialization
Cluster config lives in `docs/kubeadm-config.yaml`.
Run the init script:

```bash
./scripts/init-cluster.sh
```

After running the script, verify the node is Ready:

```bash
kubectl get nodes
```
