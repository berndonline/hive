// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	hivev1 "github.com/openshift/hive/pkg/apis/hive/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeSelectorSyncSets implements SelectorSyncSetInterface
type FakeSelectorSyncSets struct {
	Fake *FakeHiveV1
}

var selectorsyncsetsResource = schema.GroupVersionResource{Group: "hive.openshift.io", Version: "v1", Resource: "selectorsyncsets"}

var selectorsyncsetsKind = schema.GroupVersionKind{Group: "hive.openshift.io", Version: "v1", Kind: "SelectorSyncSet"}

// Get takes name of the selectorSyncSet, and returns the corresponding selectorSyncSet object, and an error if there is any.
func (c *FakeSelectorSyncSets) Get(ctx context.Context, name string, options v1.GetOptions) (result *hivev1.SelectorSyncSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(selectorsyncsetsResource, name), &hivev1.SelectorSyncSet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*hivev1.SelectorSyncSet), err
}

// List takes label and field selectors, and returns the list of SelectorSyncSets that match those selectors.
func (c *FakeSelectorSyncSets) List(ctx context.Context, opts v1.ListOptions) (result *hivev1.SelectorSyncSetList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(selectorsyncsetsResource, selectorsyncsetsKind, opts), &hivev1.SelectorSyncSetList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &hivev1.SelectorSyncSetList{ListMeta: obj.(*hivev1.SelectorSyncSetList).ListMeta}
	for _, item := range obj.(*hivev1.SelectorSyncSetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested selectorSyncSets.
func (c *FakeSelectorSyncSets) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(selectorsyncsetsResource, opts))
}

// Create takes the representation of a selectorSyncSet and creates it.  Returns the server's representation of the selectorSyncSet, and an error, if there is any.
func (c *FakeSelectorSyncSets) Create(ctx context.Context, selectorSyncSet *hivev1.SelectorSyncSet, opts v1.CreateOptions) (result *hivev1.SelectorSyncSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(selectorsyncsetsResource, selectorSyncSet), &hivev1.SelectorSyncSet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*hivev1.SelectorSyncSet), err
}

// Update takes the representation of a selectorSyncSet and updates it. Returns the server's representation of the selectorSyncSet, and an error, if there is any.
func (c *FakeSelectorSyncSets) Update(ctx context.Context, selectorSyncSet *hivev1.SelectorSyncSet, opts v1.UpdateOptions) (result *hivev1.SelectorSyncSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(selectorsyncsetsResource, selectorSyncSet), &hivev1.SelectorSyncSet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*hivev1.SelectorSyncSet), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSelectorSyncSets) UpdateStatus(ctx context.Context, selectorSyncSet *hivev1.SelectorSyncSet, opts v1.UpdateOptions) (*hivev1.SelectorSyncSet, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(selectorsyncsetsResource, "status", selectorSyncSet), &hivev1.SelectorSyncSet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*hivev1.SelectorSyncSet), err
}

// Delete takes name of the selectorSyncSet and deletes it. Returns an error if one occurs.
func (c *FakeSelectorSyncSets) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(selectorsyncsetsResource, name), &hivev1.SelectorSyncSet{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSelectorSyncSets) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(selectorsyncsetsResource, listOpts)

	_, err := c.Fake.Invokes(action, &hivev1.SelectorSyncSetList{})
	return err
}

// Patch applies the patch and returns the patched selectorSyncSet.
func (c *FakeSelectorSyncSets) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *hivev1.SelectorSyncSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(selectorsyncsetsResource, name, pt, data, subresources...), &hivev1.SelectorSyncSet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*hivev1.SelectorSyncSet), err
}
