module github.com/AlexsJones/k8s-workload-metadata-provider

go 1.15

require (
	github.com/AlexsJones/KubeOps v0.0.0-20210407143708-1f4fbe87f2be
	github.com/orcaman/concurrent-map v0.0.0-20210106121528-16402b402231
	github.com/prometheus/client_golang v1.8.0
	github.com/sirupsen/logrus v1.6.0
	golang.org/x/net v0.0.0-20201021035429-f5854403a974 // indirect
	golang.org/x/sys v0.0.0-20210403161142-5e06dd20ab57 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	k8s.io/api v0.18.4
	k8s.io/apimachinery v0.18.4
	k8s.io/client-go v0.18.4
	k8s.io/klog v1.0.0
	k8s.io/sample-controller v0.18.4
)
