## k8s-workload-metadata-provider

Built from [KubeOps](https://github.com/AlexsJones/KubeOps)

### CRD building

In another adjacent project create the CRD...

```
mkdir metadata-crd && cd metadata-crd
go mod init github.com/AlexsJones/metadata-crd
kubebuilder init --domain cloudskunkworks
kubebuilder create api --group metadata --version v1alpha1 \
 --kind MetaDataContext --resource=true --controller=true --namespaced=false
kubebuilder create api --group metadata --version v1alpha1 \
 --kind MetaDataAssociation --resource=true --controller=true --namespaced=false
```