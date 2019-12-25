// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "istio.io/client-go/pkg/apis/rbac/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeServiceRoleBindings implements ServiceRoleBindingInterface
type FakeServiceRoleBindings struct {
	Fake *FakeRbacV1alpha1
	ns   string
}

var servicerolebindingsResource = schema.GroupVersionResource{Group: "rbac", Version: "v1alpha1", Resource: "servicerolebindings"}

var servicerolebindingsKind = schema.GroupVersionKind{Group: "rbac", Version: "v1alpha1", Kind: "ServiceRoleBinding"}

// Get takes name of the serviceRoleBinding, and returns the corresponding serviceRoleBinding object, and an error if there is any.
func (c *FakeServiceRoleBindings) Get(name string, options v1.GetOptions) (result *v1alpha1.ServiceRoleBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(servicerolebindingsResource, c.ns, name), &v1alpha1.ServiceRoleBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServiceRoleBinding), err
}

// List takes label and field selectors, and returns the list of ServiceRoleBindings that match those selectors.
func (c *FakeServiceRoleBindings) List(opts v1.ListOptions) (result *v1alpha1.ServiceRoleBindingList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(servicerolebindingsResource, servicerolebindingsKind, c.ns, opts), &v1alpha1.ServiceRoleBindingList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ServiceRoleBindingList{ListMeta: obj.(*v1alpha1.ServiceRoleBindingList).ListMeta}
	for _, item := range obj.(*v1alpha1.ServiceRoleBindingList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested serviceRoleBindings.
func (c *FakeServiceRoleBindings) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(servicerolebindingsResource, c.ns, opts))

}

// Create takes the representation of a serviceRoleBinding and creates it.  Returns the server's representation of the serviceRoleBinding, and an error, if there is any.
func (c *FakeServiceRoleBindings) Create(serviceRoleBinding *v1alpha1.ServiceRoleBinding) (result *v1alpha1.ServiceRoleBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(servicerolebindingsResource, c.ns, serviceRoleBinding), &v1alpha1.ServiceRoleBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServiceRoleBinding), err
}

// Update takes the representation of a serviceRoleBinding and updates it. Returns the server's representation of the serviceRoleBinding, and an error, if there is any.
func (c *FakeServiceRoleBindings) Update(serviceRoleBinding *v1alpha1.ServiceRoleBinding) (result *v1alpha1.ServiceRoleBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(servicerolebindingsResource, c.ns, serviceRoleBinding), &v1alpha1.ServiceRoleBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServiceRoleBinding), err
}

// Delete takes name of the serviceRoleBinding and deletes it. Returns an error if one occurs.
func (c *FakeServiceRoleBindings) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(servicerolebindingsResource, c.ns, name), &v1alpha1.ServiceRoleBinding{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeServiceRoleBindings) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(servicerolebindingsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ServiceRoleBindingList{})
	return err
}

// Patch applies the patch and returns the patched serviceRoleBinding.
func (c *FakeServiceRoleBindings) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ServiceRoleBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(servicerolebindingsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ServiceRoleBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServiceRoleBinding), err
}
