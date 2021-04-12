package subscriptions

import (
	"github.com/AlexsJones/k8s-workload-metadata-provider/lib/subscription"
	"github.com/AlexsJones/k8s-workload-metadata-provider/pkg"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/klog"
)

type PodSubscriber struct{
	MetaDataProvider *pkg.MetaDataProviderController
}

func (PodSubscriber) WithElectedResource() interface{} {

	return &v1.Pod{}
}

func (PodSubscriber) WithEventType() []watch.EventType {

	return []watch.EventType {watch.Added}
}

func (PodSubscriber) OnEvent(msg subscription.Message) {

	pod := msg.Event.Object.(*v1.Pod)

	klog.V(7).Infof("Pod %s",pod.Name)

}
