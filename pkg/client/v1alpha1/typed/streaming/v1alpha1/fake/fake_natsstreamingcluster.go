// Copyright 2018 The NATS Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/nats-io/nats-streaming-operator/pkg/apis/streaming/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeNatsStreamingClusters implements NatsStreamingClusterInterface
type FakeNatsStreamingClusters struct {
	Fake *FakeStreamingV1alpha1
	ns   string
}

var natsstreamingclustersResource = schema.GroupVersionResource{Group: "streaming.nats.io", Version: "v1alpha1", Resource: "natsstreamingclusters"}

var natsstreamingclustersKind = schema.GroupVersionKind{Group: "streaming.nats.io", Version: "v1alpha1", Kind: "NatsStreamingCluster"}

// Get takes name of the natsStreamingCluster, and returns the corresponding natsStreamingCluster object, and an error if there is any.
func (c *FakeNatsStreamingClusters) Get(name string, options v1.GetOptions) (result *v1alpha1.NatsStreamingCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(natsstreamingclustersResource, c.ns, name), &v1alpha1.NatsStreamingCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NatsStreamingCluster), err
}

// List takes label and field selectors, and returns the list of NatsStreamingClusters that match those selectors.
func (c *FakeNatsStreamingClusters) List(opts v1.ListOptions) (result *v1alpha1.NatsStreamingClusterList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(natsstreamingclustersResource, natsstreamingclustersKind, c.ns, opts), &v1alpha1.NatsStreamingClusterList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.NatsStreamingClusterList{ListMeta: obj.(*v1alpha1.NatsStreamingClusterList).ListMeta}
	for _, item := range obj.(*v1alpha1.NatsStreamingClusterList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested natsStreamingClusters.
func (c *FakeNatsStreamingClusters) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(natsstreamingclustersResource, c.ns, opts))

}

// Create takes the representation of a natsStreamingCluster and creates it.  Returns the server's representation of the natsStreamingCluster, and an error, if there is any.
func (c *FakeNatsStreamingClusters) Create(natsStreamingCluster *v1alpha1.NatsStreamingCluster) (result *v1alpha1.NatsStreamingCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(natsstreamingclustersResource, c.ns, natsStreamingCluster), &v1alpha1.NatsStreamingCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NatsStreamingCluster), err
}

// Update takes the representation of a natsStreamingCluster and updates it. Returns the server's representation of the natsStreamingCluster, and an error, if there is any.
func (c *FakeNatsStreamingClusters) Update(natsStreamingCluster *v1alpha1.NatsStreamingCluster) (result *v1alpha1.NatsStreamingCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(natsstreamingclustersResource, c.ns, natsStreamingCluster), &v1alpha1.NatsStreamingCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NatsStreamingCluster), err
}

// Delete takes name of the natsStreamingCluster and deletes it. Returns an error if one occurs.
func (c *FakeNatsStreamingClusters) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(natsstreamingclustersResource, c.ns, name), &v1alpha1.NatsStreamingCluster{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeNatsStreamingClusters) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(natsstreamingclustersResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.NatsStreamingClusterList{})
	return err
}

// Patch applies the patch and returns the patched natsStreamingCluster.
func (c *FakeNatsStreamingClusters) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.NatsStreamingCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(natsstreamingclustersResource, c.ns, name, data, subresources...), &v1alpha1.NatsStreamingCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NatsStreamingCluster), err
}