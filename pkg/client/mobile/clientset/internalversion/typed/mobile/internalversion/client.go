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

package internalversion

import (
	mobile "github.com/aerogear/mobile-crd-client/pkg/apis/mobile"
	scheme "github.com/aerogear/mobile-crd-client/pkg/client/mobile/clientset/internalversion/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ClientsGetter has a method to return a ClientInterface.
// A group's client should implement this interface.
type ClientsGetter interface {
	Clients(namespace string) ClientInterface
}

// ClientInterface has methods to work with Client resources.
type ClientInterface interface {
	Create(*mobile.Client) (*mobile.Client, error)
	Update(*mobile.Client) (*mobile.Client, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*mobile.Client, error)
	List(opts v1.ListOptions) (*mobile.ClientList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *mobile.Client, err error)
	ClientExpansion
}

// clients implements ClientInterface
type clients struct {
	client rest.Interface
	ns     string
}

// newClients returns a Clients
func newClients(c *MobileClient, namespace string) *clients {
	return &clients{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the client, and returns the corresponding client object, and an error if there is any.
func (c *clients) Get(name string, options v1.GetOptions) (result *mobile.Client, err error) {
	result = &mobile.Client{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("clients").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Clients that match those selectors.
func (c *clients) List(opts v1.ListOptions) (result *mobile.ClientList, err error) {
	result = &mobile.ClientList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("clients").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested clients.
func (c *clients) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("clients").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a client and creates it.  Returns the server's representation of the client, and an error, if there is any.
func (c *clients) Create(client *mobile.Client) (result *mobile.Client, err error) {
	result = &mobile.Client{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("clients").
		Body(client).
		Do().
		Into(result)
	return
}

// Update takes the representation of a client and updates it. Returns the server's representation of the client, and an error, if there is any.
func (c *clients) Update(client *mobile.Client) (result *mobile.Client, err error) {
	result = &mobile.Client{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("clients").
		Name(client.Name).
		Body(client).
		Do().
		Into(result)
	return
}

// Delete takes name of the client and deletes it. Returns an error if one occurs.
func (c *clients) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("clients").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *clients) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("clients").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched client.
func (c *clients) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *mobile.Client, err error) {
	result = &mobile.Client{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("clients").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
