## k8s-workload-metadata-provider

[![Docker](https://github.com/AlexsJones/k8s-workload-metadata-provider/actions/workflows/docker-publish.yml/badge.svg)](https://github.com/AlexsJones/k8s-workload-metadata-provider/actions/workflows/docker-publish.yml)

Built from [KubeOps](https://github.com/AlexsJones/KubeOps)

### Install the Custom Resource Definition on the cluster...

`kubectl apply -f apis/crd.yml`

### How it works

```go
	err = runtime.EventBuffer(ctx, kubeClient,
		&subscription.Registry{
			Subscriptions: []subscription.ISubscription{
				subscriptions.MetaDataContextSubscriber{},
			},
		}, []watcher.IObject{
			kubeClient.CoreV1().Pods(""),
			metaDataClient.MetaDataContextTypes(""), //Watch for our CRD
		})
```

#### Testing

You can then create an example of this CRD with `kubectl apply -f apis/example-resource.yaml`