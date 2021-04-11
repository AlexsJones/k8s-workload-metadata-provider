package subscriptions

import (
	"github.com/AlexsJones/k8s-workload-metadata-provider/apis/metadata/alphav1"
	"github.com/AlexsJones/k8s-workload-metadata-provider/lib/subscription"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/klog"
)

type MetaDataContextSubscriber struct{}

func (MetaDataContextSubscriber) WithElectedResource() interface{} {

	return &alphav1.MetaDataContextType{}
}

func (MetaDataContextSubscriber) WithEventType() []watch.EventType {

	return []watch.EventType {watch.Added, watch.Deleted, watch.Modified}
}

func (MetaDataContextSubscriber) OnEvent(msg subscription.Message) {

	kind := msg.Event.Object.GetObjectKind()

	klog.Infof("Found a new %s", kind.GroupVersionKind().Kind)

}
