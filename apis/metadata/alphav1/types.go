package alphav1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MetaDataContextType is a top-level type
type MetaDataContextType struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +optional
	Status MetaDataContextTypeStatus `json:"status,omitempty"`
	// This is where you can define
	// your own custom spec
	Spec MetaDataContextSpec `json:"spec,omitempty"`
}

// custom spec
type MetaDataContextSpec struct {
	Message string `json:"message,omitempty"`
	DataMappingExample map[string]string`json:"datamapping,omitempty"`
}

// custom status
type MetaDataContextTypeStatus struct {
	Name string
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// no client needed for list as it's been created in above
type MetaDataContextTypeList struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ListMeta `son:"metadata,omitempty"`

	Items []MetaDataContextType `json:"items"`
}