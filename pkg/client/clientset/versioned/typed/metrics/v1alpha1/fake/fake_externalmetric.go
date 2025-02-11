// Copyright 2019 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// A copy of the License is located at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/cronofy/k8s-cloudwatch-adapter/pkg/apis/metrics/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeExternalMetrics implements ExternalMetricInterface
type FakeExternalMetrics struct {
	Fake *FakeMetricsV1alpha1
	ns   string
}

var externalmetricsResource = schema.GroupVersionResource{Group: "metrics.aws", Version: "v1alpha1", Resource: "externalmetrics"}

var externalmetricsKind = schema.GroupVersionKind{Group: "metrics.aws", Version: "v1alpha1", Kind: "ExternalMetric"}

// Get takes name of the externalMetric, and returns the corresponding externalMetric object, and an error if there is any.
func (c *FakeExternalMetrics) Get(name string, options v1.GetOptions) (result *v1alpha1.ExternalMetric, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(externalmetricsResource, c.ns, name), &v1alpha1.ExternalMetric{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ExternalMetric), err
}

// List takes label and field selectors, and returns the list of ExternalMetrics that match those selectors.
func (c *FakeExternalMetrics) List(opts v1.ListOptions) (result *v1alpha1.ExternalMetricList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(externalmetricsResource, externalmetricsKind, c.ns, opts), &v1alpha1.ExternalMetricList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ExternalMetricList{ListMeta: obj.(*v1alpha1.ExternalMetricList).ListMeta}
	for _, item := range obj.(*v1alpha1.ExternalMetricList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested externalMetrics.
func (c *FakeExternalMetrics) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(externalmetricsResource, c.ns, opts))

}

// Create takes the representation of a externalMetric and creates it.  Returns the server's representation of the externalMetric, and an error, if there is any.
func (c *FakeExternalMetrics) Create(externalMetric *v1alpha1.ExternalMetric) (result *v1alpha1.ExternalMetric, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(externalmetricsResource, c.ns, externalMetric), &v1alpha1.ExternalMetric{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ExternalMetric), err
}

// Update takes the representation of a externalMetric and updates it. Returns the server's representation of the externalMetric, and an error, if there is any.
func (c *FakeExternalMetrics) Update(externalMetric *v1alpha1.ExternalMetric) (result *v1alpha1.ExternalMetric, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(externalmetricsResource, c.ns, externalMetric), &v1alpha1.ExternalMetric{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ExternalMetric), err
}

// Delete takes name of the externalMetric and deletes it. Returns an error if one occurs.
func (c *FakeExternalMetrics) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(externalmetricsResource, c.ns, name), &v1alpha1.ExternalMetric{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeExternalMetrics) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(externalmetricsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ExternalMetricList{})
	return err
}
