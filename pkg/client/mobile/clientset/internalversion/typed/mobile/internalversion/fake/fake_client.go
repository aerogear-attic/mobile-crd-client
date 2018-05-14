/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package fake

import (
	mobile "github.com/aerogear/mobile-crd-client/pkg/apis/mobile"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeClients implements ClientInterface
type FakeClients struct {
	Fake *FakeMobile
	ns   string
}

var clientsResource = schema.GroupVersionResource{Group: "mobile.k8s.io", Version: "", Resource: "clients"}

var clientsKind = schema.GroupVersionKind{Group: "mobile.k8s.io", Version: "", Kind: "Client"}

// Get takes name of the client, and returns the corresponding client object, and an error if there is any.
func (c *FakeClients) Get(name string, options v1.GetOptions) (result *mobile.Client, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(clientsResource, c.ns, name), &mobile.Client{})

	if obj == nil {
		return nil, err
	}
	return obj.(*mobile.Client), err
}

// List takes label and field selectors, and returns the list of Clients that match those selectors.
func (c *FakeClients) List(opts v1.ListOptions) (result *mobile.ClientList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(clientsResource, clientsKind, c.ns, opts), &mobile.ClientList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &mobile.ClientList{}
	for _, item := range obj.(*mobile.ClientList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clients.
func (c *FakeClients) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(clientsResource, c.ns, opts))

}

// Create takes the representation of a client and creates it.  Returns the server's representation of the client, and an error, if there is any.
func (c *FakeClients) Create(client *mobile.Client) (result *mobile.Client, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(clientsResource, c.ns, client), &mobile.Client{})

	if obj == nil {
		return nil, err
	}
	return obj.(*mobile.Client), err
}

// Update takes the representation of a client and updates it. Returns the server's representation of the client, and an error, if there is any.
func (c *FakeClients) Update(client *mobile.Client) (result *mobile.Client, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(clientsResource, c.ns, client), &mobile.Client{})

	if obj == nil {
		return nil, err
	}
	return obj.(*mobile.Client), err
}

// Delete takes name of the client and deletes it. Returns an error if one occurs.
func (c *FakeClients) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(clientsResource, c.ns, name), &mobile.Client{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClients) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(clientsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &mobile.ClientList{})
	return err
}

// Patch applies the patch and returns the patched client.
func (c *FakeClients) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *mobile.Client, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(clientsResource, c.ns, name, data, subresources...), &mobile.Client{})

	if obj == nil {
		return nil, err
	}
	return obj.(*mobile.Client), err
}
