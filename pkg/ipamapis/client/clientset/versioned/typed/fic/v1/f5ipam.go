/*
Copyright The Kubernetes Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"time"

	v1 "github.com/F5Networks/f5-ipam-controller/pkg/ipamapis/apis/fic/v1"
	scheme "github.com/F5Networks/f5-ipam-controller/pkg/ipamapis/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// F5IPAMsGetter has a method to return a F5IPAMInterface.
// A group's client should implement this interface.
type F5IPAMsGetter interface {
	F5IPAMs(namespace string) F5IPAMInterface
}

// F5IPAMInterface has methods to work with F5IPAM resources.
type F5IPAMInterface interface {
	Create(*v1.F5IPAM) (*v1.F5IPAM, error)
	Update(*v1.F5IPAM) (*v1.F5IPAM, error)
	UpdateStatus(*v1.F5IPAM) (*v1.F5IPAM, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.F5IPAM, error)
	List(opts metav1.ListOptions) (*v1.F5IPAMList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.F5IPAM, err error)
	F5IPAMExpansion
}

// f5IPAMs implements F5IPAMInterface
type f5IPAMs struct {
	client rest.Interface
	ns     string
}

// newF5IPAMs returns a F5IPAMs
func newF5IPAMs(c *K8sV1Client, namespace string) *f5IPAMs {
	return &f5IPAMs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the f5IPAM, and returns the corresponding f5IPAM object, and an error if there is any.
func (c *f5IPAMs) Get(name string, options metav1.GetOptions) (result *v1.F5IPAM, err error) {
	result = &v1.F5IPAM{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("f5ipams").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of F5IPAMs that match those selectors.
func (c *f5IPAMs) List(opts metav1.ListOptions) (result *v1.F5IPAMList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.F5IPAMList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("f5ipams").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested f5IPAMs.
func (c *f5IPAMs) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("f5ipams").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a f5IPAM and creates it.  Returns the server's representation of the f5IPAM, and an error, if there is any.
func (c *f5IPAMs) Create(f5IPAM *v1.F5IPAM) (result *v1.F5IPAM, err error) {
	result = &v1.F5IPAM{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("f5ipams").
		Body(f5IPAM).
		Do().
		Into(result)
	return
}

// Update takes the representation of a f5IPAM and updates it. Returns the server's representation of the f5IPAM, and an error, if there is any.
func (c *f5IPAMs) Update(f5IPAM *v1.F5IPAM) (result *v1.F5IPAM, err error) {
	result = &v1.F5IPAM{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("f5ipams").
		Name(f5IPAM.Name).
		Body(f5IPAM).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *f5IPAMs) UpdateStatus(f5IPAM *v1.F5IPAM) (result *v1.F5IPAM, err error) {
	result = &v1.F5IPAM{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("f5ipams").
		Name(f5IPAM.Name).
		SubResource("status").
		Body(f5IPAM).
		Do().
		Into(result)
	return
}

// Delete takes name of the f5IPAM and deletes it. Returns an error if one occurs.
func (c *f5IPAMs) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("f5ipams").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *f5IPAMs) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("f5ipams").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched f5IPAM.
func (c *f5IPAMs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.F5IPAM, err error) {
	result = &v1.F5IPAM{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("f5ipams").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
