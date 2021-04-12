package pkg

import (
	"context"
	m1client "github.com/AlexsJones/k8s-workload-metadata-provider/apis/client/clientset/versioned/typed/metadata.cloudskunkworks/v1"
	metadatav1 "github.com/AlexsJones/k8s-workload-metadata-provider/apis/metadata.cloudskunkworks/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
	"k8s.io/klog"
	"sync"
)

type MetaDataProviderController struct {
	KubeClient *kubernetes.Clientset
	MetaDataClient *m1client.MetadataV1Client
	// TempMetaDataReference storage uses a concurrent map string // MetaDataContextType k/v
	TempMetaDataReferenceStorage *sync.Map
}

func NewMetaDataProviderController(kubeClient *kubernetes.Clientset,metaDataClient *m1client.MetadataV1Client) *MetaDataProviderController {

	return &MetaDataProviderController{
		KubeClient: kubeClient,
		MetaDataClient: metaDataClient,
		TempMetaDataReferenceStorage: new(sync.Map),
	}
}

func (m *MetaDataProviderController) OnPodEvent(event watch.Event, pod *v1.Pod) {

	// Get the pods reference owner
	if len(pod.GetOwnerReferences()) == 0 {
		return
	}
	if pod.GetOwnerReferences()[0].Kind == "ReplicaSet" {
		rs, err := m.KubeClient.AppsV1().ReplicaSets(pod.GetNamespace()).
			Get(context.Background(),pod.GetOwnerReferences()[0].Name,metav1.GetOptions{})
		if err != nil {
			klog.Errorf(err.Error())
			return
		}
		if len(rs.GetOwnerReferences()) == 0 {
			return
		}
		if rs.GetOwnerReferences()[0].Kind == "Deployment"  {
			deployment, err := m.KubeClient.AppsV1().ReplicaSets(pod.GetNamespace()).
				Get(context.Background(),pod.GetOwnerReferences()[0].Name,metav1.GetOptions{})
			if err != nil {
				klog.Errorf(err.Error())
				return
			}
			// Check for annotations
			if deployment.Annotations["metaDataContext"] != "" {
		 		klog.V(4).Infof("Deployment %s is metaDataContext aware",deployment.Name)

		 		// Check whether the deployment is looking for an existing context
		 		// We can do this initially through the local cache then the remote
		 		if _, ok := m.TempMetaDataReferenceStorage.Load(deployment.Annotations["metaDataContext"]); ok {
		 			// Found the local cache reference
		 			// Check the active API
					mActiveType, err := m.MetaDataClient.MetaDataContextTypes("").Get(context.Background(),deployment.Annotations["metaDataContext"],metav1.GetOptions{})
					if err != nil {
						klog.Error(err.Error())
						return
					}
					// Do something with the CRD
					klog.V(4).Infof("Deployment %s has requested %s MetaDataContextType ",deployment.Name, mActiveType.Name)

		 		}else {
					klog.Warningln("Local cache of MetaDataContextType not found")
					return
				}

			}else {
				klog.V(4).Infof("Deployment %s is not metaDataContext aware",deployment.Name)
			}
		}
	}
}

func (m *MetaDataProviderController) OnMetaDataContextTypeEvent(event watch.Event,
	context *metadatav1.MetaDataContextType) {

	localObject := &context

	if context.Name == "" {
		klog.Errorf("Unable to retrieve name from incoming event")
		return
	}
	switch event.Type {
	case watch.Added:
		if _, ok := m.TempMetaDataReferenceStorage.Load(context.Name); !ok {
			m.TempMetaDataReferenceStorage.Store(context.Name,localObject)
		}
		break
	case watch.Modified:
			m.TempMetaDataReferenceStorage.Store(context.Name,localObject)
		break
	case watch.Deleted:
		m.TempMetaDataReferenceStorage.Delete(context.Name)
		break
	}

	count := 0
	m.TempMetaDataReferenceStorage.Range(func(key, value interface{}) bool {
		count++
		return true
	})
	klog.V(7).Infof("Current TempMetaDataReferenceStorage size is %d", count)
}

func (m *MetaDataProviderController) ControlLoop(cancelContext context.Context) {

	for {
		select {
		case <-cancelContext.Done():
			break
		}



	}
}