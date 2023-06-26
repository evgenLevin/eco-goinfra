package sriov

import (
	"context"
	"fmt"

	"github.com/golang/glog"

	srIovV1 "github.com/k8snetworkplumbingwg/sriov-network-operator/api/v1"
	"github.com/openshift-kni/eco-goinfra/pkg/clients"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NetworkNodeStateListBuilder provides struct for SriovNetworkNodeStateList object
// which contains connection to cluster and SriovNetworkNodeStateList definitions.
type NetworkNodeStateListBuilder struct {
	// Dynamically discovered SriovNetworkNodeStateList object.
	Objects *srIovV1.SriovNetworkNodeStateList
	// apiClient opens api connection to the cluster.
	apiClient *clients.Settings
	// nsName defines SrIov operator namespace.
	nsName string
	// errorMsg used in discovery function before sending api request to cluster.
	errorMsg string
}

// NewNetworkNodeStateListBuilder creates new instance of NetworkNodeStateListBuilder.
func NewNetworkNodeStateListBuilder(apiClient *clients.Settings, nsname string) *NetworkNodeStateListBuilder {
	glog.V(100).Infof(
		"Initializing new NetworkNodeStateListBuilder structure with the following params: %s",
		nsname)

	builder := &NetworkNodeStateListBuilder{
		apiClient: apiClient,
		nsName:    nsname,
	}

	if nsname == "" {
		glog.V(100).Infof("The namespace of the SriovNetworkNodeStateList is empty")

		builder.errorMsg = "SriovNetworkNodeStateList 'nsname' is empty"
	}

	return builder
}

// Discover method gets the SriovNetworkNodeStateList items and stores them in the NetworkNodeStateListBuilder struct.
func (builder *NetworkNodeStateListBuilder) Discover() error {
	if builder.errorMsg != "" {
		return fmt.Errorf(builder.errorMsg)
	}

	glog.V(100).Infof("Getting the SriovNetworkNodeStateList object in namespace %s",
		builder.nsName)

	var err error
	builder.Objects, err = builder.apiClient.SriovNetworkNodeStates(builder.nsName).List(
		context.TODO(), v1.ListOptions{})

	return err
}
