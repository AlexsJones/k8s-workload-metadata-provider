## k8s-workload-metadata-provider

[![Docker](https://github.com/AlexsJones/k8s-workload-metadata-provider/actions/workflows/docker-publish.yml/badge.svg)](https://github.com/AlexsJones/k8s-workload-metadata-provider/actions/workflows/docker-publish.yml)

Built from [KubeOps](https://github.com/AlexsJones/KubeOps)


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

