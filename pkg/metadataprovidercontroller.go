package pkg

import (
	"context"
	metadatav1 "github.com/AlexsJones/k8s-workload-metadata-provider/apis/metadata.cloudskunkworks/v1"
	"github.com/orcaman/concurrent-map"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
	"k8s.io/klog"
)

type MetaDataProviderController struct {
	KubeClient *kubernetes.Clientset
	// TempMetaDataReference storage uses a concurrent map string // MetaDataContextType k/v
	TempMetaDataReferenceStorage cmap.ConcurrentMap
}

func (*MetaDataProviderController) OnPodEvent(event watch.EventType, pod *v1.Pod) {

}

func (m *MetaDataProviderController) OnMetaDataContextTypeEvent(event watch.Event,
	context *metadatav1.MetaDataContextType) {

	switch event.Type {
	case watch.Added:
		m.TempMetaDataReferenceStorage.Set(context.Name,context)
		break
	case watch.Modified:
		m.TempMetaDataReferenceStorage.Set(context.Name,context)
		break
	case watch.Deleted:
		m.TempMetaDataReferenceStorage.Remove(context.Name)
		break
	}
	klog.Info("Referential storage has %d MetaDataContextType objects being tracked", m.TempMetaDataReferenceStorage.Count())
}

func (m *MetaDataProviderController) ControlLoop(cancelContext context.Context) {

	m.TempMetaDataReferenceStorage = cmap.New()


	for {
		select {
		case <-cancelContext.Done():
			break
		}



	}
}