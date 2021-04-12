package subscriptions

import (
	"github.com/AlexsJones/k8s-workload-metadata-provider/apis/metadata.cloudskunkworks/v1"
	"github.com/AlexsJones/k8s-workload-metadata-provider/lib/subscription"
	"github.com/AlexsJones/k8s-workload-metadata-provider/pkg"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/klog"
)

type MetaDataContextSubscriber struct{
	MetaDataProvider pkg.MetaDataProviderController
}

func (MetaDataContextSubscriber) WithElectedResource() interface{} {

	return &v1.MetaDataContextType{}
}

func (MetaDataContextSubscriber) WithEventType() []watch.EventType {

	return []watch.EventType {watch.Added, watch.Deleted, watch.Modified}
}

func (m MetaDataContextSubscriber) OnEvent(msg subscription.Message) {

	context := msg.Event.Object.(*v1.MetaDataContextType)

	klog.V(7).Infof("MetaDataContextType %s",context.Name)

	for k, v := range context.Spec.DataMapping {
		klog.Infof("Key %s: Value: %v", k,v)
	}

	m.MetaDataProvider.OnMetaDataContextTypeEvent(msg.Event, context)
}
