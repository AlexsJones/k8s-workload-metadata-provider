package pkg

import (
	"context"
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
	// TempMetaDataReference storage uses a concurrent map string // MetaDataContextType k/v
	TempMetaDataReferenceStorage *sync.Map
}

func NewMetaDataProviderController(kubeClient *kubernetes.Clientset) *MetaDataProviderController {

	return &MetaDataProviderController{
		KubeClient: kubeClient,
		TempMetaDataReferenceStorage: new(sync.Map),
	}
}

func (m *MetaDataProviderController) OnPodEvent(event watch.Event, pod *v1.Pod) {

	// Get the pods reference owner
	if pod.GetOwnerReferences()[0].Kind == "ReplicaSet" {
		rs, err := m.KubeClient.AppsV1().ReplicaSets(pod.GetNamespace()).
			Get(context.Background(),pod.GetOwnerReferences()[0].Name,metav1.GetOptions{})
		if err != nil {
			klog.Errorf(err.Error())
			return
		}
		if rs.GetOwnerReferences()[0].Kind == "Deployment" {
			deployment, err := m.KubeClient.AppsV1().ReplicaSets(pod.GetNamespace()).
				Get(context.Background(),pod.GetOwnerReferences()[0].Name,metav1.GetOptions{})
			if err != nil {
				klog.Errorf(err.Error())
				return
			}
			// Check for annotations
			if deployment.Annotations["metaDataContextAware"] != "" {
		 		klog.V(4).Infof("Deployment %s is metaDataContextAware aware",deployment.Name)
			}else {
				klog.V(4).Infof("Deployment %s is not metaDataContextAware aware",deployment.Name)
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