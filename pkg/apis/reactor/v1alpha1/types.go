package v1alpha1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// These const variables are used in our custom controller.
const (
	GroupName string = "huyage.dev"
	Kind      string = "Reactor"
	Version   string = "v1alpha1"
	Plural    string = "reactors"
	Singlular string = "reactor"
	ShortName string = "rt"
	Name      string = Plural + "." + GroupName
)

// ReactorSpec specifies the 'spec' of Reactor CRD.
type ReactorSpec struct {
	ReactTo string `json:"reactTo"`
	Product struct {
		Name  string `json:"name"`
		Image string `json:"image"`
	} `json:"product"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Reactor describes a Reactor custom resource.
type Reactor struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ReactorSpec `json:"spec"`
	Status string      `json:"status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ReactorList is a list of Reactor resources.
type ReactorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Reactor `json:"items"`
}
