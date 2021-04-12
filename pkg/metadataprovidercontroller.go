package pkg

import (
	"context"
	metadatav1 "github.com/AlexsJones/k8s-workload-metadata-provider/apis/metadata.cloudskunkworks/v1"
	v1 "k8s.io/api/core/v1"
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

func (*MetaDataProviderController) OnPodEvent(event watch.EventType, pod *v1.Pod) {

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
}

func (m *MetaDataProviderController) ControlLoop(cancelContext context.Context) {

	for {
		select {
		case <-cancelContext.Done():
			break
		}



	}
}