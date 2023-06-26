package sriov

import (
	"context"
	"fmt"

	"github.com/golang/glog"

	srIovV1 "github.com/k8snetworkplumbingwg/sriov-network-operator/api/v1"
	"github.com/openshift-kni/eco-goinfra/pkg/clients"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// OperatorConfigBuilder provides struct for SriovOperatorConfigBuilder object which contains connection to cluster and
// SriovOperatorConfig definitions.
type OperatorConfigBuilder struct {
	// Dynamically discovered SriovOperatorConfig object.
	Objects *srIovV1.SriovOperatorConfig
	// apiClient opens api connection to the cluster.
	apiClient *clients.Settings
	// nodeName defines on what node SriovOperatorConfig resource should be queried.
	nodeName string
	// nsName defines SriovOperatorConfig operator namespace.
	nsName string
	// errorMsg used in discovery function before sending api request to cluster.
	errorMsg string
}

// NewOperatorConfigBuilder creates new instance of SriovOperatorConfig.
func NewOperatorConfigBuilder(apiClient *clients.Settings, nameConfig, nsname string) *OperatorConfigBuilder {
	glog.V(100).Infof(
		"Initializing new OperatorConfigBuilder structure with the following params: %s, %s",
		nameConfig, nsname)

	builder := &OperatorConfigBuilder{
		apiClient: apiClient,
		nodeName:  nameConfig,
		nsName:    nsname,
	}

	if nameConfig == "" {
		glog.V(100).Infof("The name of the nameConfig is empty")

		builder.errorMsg = "SriovOperatorConfig 'nameConfig' is empty"
	}

	if nsname == "" {
		glog.V(100).Infof("The namespace of the SriovOperatorConfig is empty")

		builder.errorMsg = "SriovOperatorConfig 'nsname' is empty"
	}

	return builder
}

// Discover method gets the SriovOperatorConfig items and stores them in the OperatorConfigBuilder struct.
func (builder *OperatorConfigBuilder) Discover() error {
	if builder.errorMsg != "" {
		return fmt.Errorf(builder.errorMsg)
	}

	glog.V(100).Infof("Getting the SriovOperatorConfig object in namespace %s for node %s",
		builder.nsName, builder.nodeName)

	var err error
	builder.Objects, err = builder.apiClient.SriovOperatorConfigs(builder.nsName).Get(
		context.TODO(), builder.nodeName, v1.GetOptions{})

	return err
}

// Update renovates the existing SriovOperatorConfig object with the SriovOperatorConfig definition in builder.
func (builder *OperatorConfigBuilder) Update() (*OperatorConfigBuilder, error) {
	glog.V(100).Infof("Updating the SriovOperatorConfig object %s in namespace %s",
		builder.Objects.Name, builder.Objects.Namespace)

	if builder.errorMsg != "" {
		return nil, fmt.Errorf(builder.errorMsg)
	}

	err := builder.apiClient.Update(context.TODO(), builder.Objects)
	if err != nil {
		return nil, fmt.Errorf("failed to update SriovOperatorConfig object %s in namespace %s",
			builder.Objects.Name, builder.Objects.Namespace)
	}

	return builder, err
}
