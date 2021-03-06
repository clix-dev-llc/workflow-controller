// Generated file, do not modify manually!
package versioned

import (
	glog "github.com/golang/glog"
	workflowv1 "github.com/sdminonne/workflow-controller/pkg/client/versioned/typed/workflow/v1"
	discovery "k8s.io/client-go/discovery"
	rest "k8s.io/client-go/rest"
	flowcontrol "k8s.io/client-go/util/flowcontrol"
)

type Interface interface {
	Discovery() discovery.DiscoveryInterface
	WorkflowV1() workflowv1.WorkflowV1Interface
	// Deprecated: please explicitly pick a version if possible.
	Workflow() workflowv1.WorkflowV1Interface
}

// Clientset contains the clients for groups. Each group has exactly one
// version included in a Clientset.
type Clientset struct {
	*discovery.DiscoveryClient
	workflowV1 *workflowv1.WorkflowV1Client
}

// WorkflowV1 retrieves the WorkflowV1Client
func (c *Clientset) WorkflowV1() workflowv1.WorkflowV1Interface {
	return c.workflowV1
}

// Deprecated: Workflow retrieves the default version of WorkflowClient.
// Please explicitly pick a version.
func (c *Clientset) Workflow() workflowv1.WorkflowV1Interface {
	return c.workflowV1
}

// Discovery retrieves the DiscoveryClient
func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	if c == nil {
		return nil
	}
	return c.DiscoveryClient
}

// NewForConfig creates a new Clientset for the given config.
func NewForConfig(c *rest.Config) (*Clientset, error) {
	configShallowCopy := *c
	if configShallowCopy.RateLimiter == nil && configShallowCopy.QPS > 0 {
		configShallowCopy.RateLimiter = flowcontrol.NewTokenBucketRateLimiter(configShallowCopy.QPS, configShallowCopy.Burst)
	}
	var cs Clientset
	var err error
	cs.workflowV1, err = workflowv1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}

	cs.DiscoveryClient, err = discovery.NewDiscoveryClientForConfig(&configShallowCopy)
	if err != nil {
		glog.Errorf("failed to create the DiscoveryClient: %v", err)
		return nil, err
	}
	return &cs, nil
}

// NewForConfigOrDie creates a new Clientset for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *Clientset {
	var cs Clientset
	cs.workflowV1 = workflowv1.NewForConfigOrDie(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClientForConfigOrDie(c)
	return &cs
}

// New creates a new Clientset for the given RESTClient.
func New(c rest.Interface) *Clientset {
	var cs Clientset
	cs.workflowV1 = workflowv1.New(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClient(c)
	return &cs
}
