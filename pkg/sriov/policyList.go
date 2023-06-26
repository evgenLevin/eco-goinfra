package sriov

import (
	"context"
	"fmt"

	"github.com/golang/glog"

	srIovV1 "github.com/k8snetworkplumbingwg/sriov-network-operator/api/v1"
	"github.com/openshift-kni/eco-goinfra/pkg/clients"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// PolicyListBuilder provides struct for SriovNetworkNodePolicyList object which contains connection to cluster and
// SriovNetworkNodePolicyList definitions.
type PolicyListBuilder struct {
	// Dynamically discovered SriovNetworkNodePolicyList object.
	Objects *srIovV1.SriovNetworkNodePolicyList
	// apiClient opens api connection to the cluster.
	apiClient *clients.Settings
	// nsName defines SrIov operator namespace.
	nsName string
	// errorMsg used in discovery function before sending api request to cluster.
	errorMsg string
}

// NewPolicyListBuilder creates new instance of PolicyListBuilder.
func NewPolicyListBuilder(apiClient *clients.Settings, nsname string) *PolicyListBuilder {
	glog.V(100).Infof(
		"Initializing new PolicyListBuilder structure with the following params: %s",
		nsname)

	builder := &PolicyListBuilder{
		apiClient: apiClient,
		nsName:    nsname,
	}

	if nsname == "" {
		glog.V(100).Infof("The namespace of the SriovNetworkNodePolicyList is empty")

		builder.errorMsg = "SriovNetworkNodePolicyList 'nsname' is empty"
	}

	return builder
}

// Discover method gets the SriovNetworkNodePolicyList items and stores them in the PolicyListBuilder struct.
func (builder *PolicyListBuilder) Discover() error {
	if builder.errorMsg != "" {
		return fmt.Errorf(builder.errorMsg)
	}

	glog.V(100).Infof("Getting the SriovNetworkNodePolicyList object in namespace %s",
		builder.nsName)

	var err error
	builder.Objects, err = builder.apiClient.SriovNetworkNodePolicies(builder.nsName).List(
		context.TODO(), v1.ListOptions{})

	return err
}
