// Code generated by astool. DO NOT EDIT.

package vocab

import "net/url"

// GoToSocialCanAnnouncePropertyIterator represents a single value for the
// "canAnnounce" property.
type GoToSocialCanAnnouncePropertyIterator interface {
	// Get returns the value of this property. When IsGoToSocialCanAnnounce
	// returns false, Get will return any arbitrary value.
	Get() GoToSocialCanAnnounce
	// GetIRI returns the IRI of this property. When IsIRI returns false,
	// GetIRI will return any arbitrary value.
	GetIRI() *url.URL
	// GetType returns the value in this property as a Type. Returns nil if
	// the value is not an ActivityStreams type, such as an IRI or another
	// value.
	GetType() Type
	// HasAny returns true if the value or IRI is set.
	HasAny() bool
	// IsGoToSocialCanAnnounce returns true if this property is set and not an
	// IRI.
	IsGoToSocialCanAnnounce() bool
	// IsIRI returns true if this property is an IRI.
	IsIRI() bool
	// JSONLDContext returns the JSONLD URIs required in the context string
	// for this property and the specific values that are set. The value
	// in the map is the alias used to import the property's value or
	// values.
	JSONLDContext() map[string]string
	// KindIndex computes an arbitrary value for indexing this kind of value.
	// This is a leaky API detail only for folks looking to replace the
	// go-fed implementation. Applications should not use this method.
	KindIndex() int
	// LessThan compares two instances of this property with an arbitrary but
	// stable comparison. Applications should not use this because it is
	// only meant to help alternative implementations to go-fed to be able
	// to normalize nonfunctional properties.
	LessThan(o GoToSocialCanAnnouncePropertyIterator) bool
	// Name returns the name of this property: "GoToSocialCanAnnounce".
	Name() string
	// Next returns the next iterator, or nil if there is no next iterator.
	Next() GoToSocialCanAnnouncePropertyIterator
	// Prev returns the previous iterator, or nil if there is no previous
	// iterator.
	Prev() GoToSocialCanAnnouncePropertyIterator
	// Set sets the value of this property. Calling IsGoToSocialCanAnnounce
	// afterwards will return true.
	Set(v GoToSocialCanAnnounce)
	// SetIRI sets the value of this property. Calling IsIRI afterwards will
	// return true.
	SetIRI(v *url.URL)
	// SetType attempts to set the property for the arbitrary type. Returns an
	// error if it is not a valid type to set on this property.
	SetType(t Type) error
}

// Defines who can Announce with an object property set to the URI/ID of the
// Object to which this interactionPolicy is attached.
type GoToSocialCanAnnounceProperty interface {
	// AppendGoToSocialCanAnnounce appends a CanAnnounce value to the back of
	// a list of the property "canAnnounce". Invalidates iterators that
	// are traversing using Prev.
	AppendGoToSocialCanAnnounce(v GoToSocialCanAnnounce)
	// AppendIRI appends an IRI value to the back of a list of the property
	// "canAnnounce"
	AppendIRI(v *url.URL)
	// PrependType prepends an arbitrary type value to the front of a list of
	// the property "canAnnounce". Invalidates iterators that are
	// traversing using Prev. Returns an error if the type is not a valid
	// one to set for this property.
	AppendType(t Type) error
	// At returns the property value for the specified index. Panics if the
	// index is out of bounds.
	At(index int) GoToSocialCanAnnouncePropertyIterator
	// Begin returns the first iterator, or nil if empty. Can be used with the
	// iterator's Next method and this property's End method to iterate
	// from front to back through all values.
	Begin() GoToSocialCanAnnouncePropertyIterator
	// Empty returns returns true if there are no elements.
	Empty() bool
	// End returns beyond-the-last iterator, which is nil. Can be used with
	// the iterator's Next method and this property's Begin method to
	// iterate from front to back through all values.
	End() GoToSocialCanAnnouncePropertyIterator
	// InsertGoToSocialCanAnnounce inserts a CanAnnounce value at the
	// specified index for a property "canAnnounce". Existing elements at
	// that index and higher are shifted back once. Invalidates all
	// iterators.
	InsertGoToSocialCanAnnounce(idx int, v GoToSocialCanAnnounce)
	// Insert inserts an IRI value at the specified index for a property
	// "canAnnounce". Existing elements at that index and higher are
	// shifted back once. Invalidates all iterators.
	InsertIRI(idx int, v *url.URL)
	// PrependType prepends an arbitrary type value to the front of a list of
	// the property "canAnnounce". Invalidates all iterators. Returns an
	// error if the type is not a valid one to set for this property.
	InsertType(idx int, t Type) error
	// JSONLDContext returns the JSONLD URIs required in the context string
	// for this property and the specific values that are set. The value
	// in the map is the alias used to import the property's value or
	// values.
	JSONLDContext() map[string]string
	// KindIndex computes an arbitrary value for indexing this kind of value.
	// This is a leaky API method specifically needed only for alternate
	// implementations for go-fed. Applications should not use this
	// method. Panics if the index is out of bounds.
	KindIndex(idx int) int
	// Len returns the number of values that exist for the "canAnnounce"
	// property.
	Len() (length int)
	// Less computes whether another property is less than this one. Mixing
	// types results in a consistent but arbitrary ordering
	Less(i, j int) bool
	// LessThan compares two instances of this property with an arbitrary but
	// stable comparison. Applications should not use this because it is
	// only meant to help alternative implementations to go-fed to be able
	// to normalize nonfunctional properties.
	LessThan(o GoToSocialCanAnnounceProperty) bool
	// Name returns the name of this property ("canAnnounce") with any alias.
	Name() string
	// PrependGoToSocialCanAnnounce prepends a CanAnnounce value to the front
	// of a list of the property "canAnnounce". Invalidates all iterators.
	PrependGoToSocialCanAnnounce(v GoToSocialCanAnnounce)
	// PrependIRI prepends an IRI value to the front of a list of the property
	// "canAnnounce".
	PrependIRI(v *url.URL)
	// PrependType prepends an arbitrary type value to the front of a list of
	// the property "canAnnounce". Invalidates all iterators. Returns an
	// error if the type is not a valid one to set for this property.
	PrependType(t Type) error
	// Remove deletes an element at the specified index from a list of the
	// property "canAnnounce", regardless of its type. Panics if the index
	// is out of bounds. Invalidates all iterators.
	Remove(idx int)
	// Serialize converts this into an interface representation suitable for
	// marshalling into a text or binary format. Applications should not
	// need this function as most typical use cases serialize types
	// instead of individual properties. It is exposed for alternatives to
	// go-fed implementations to use.
	Serialize() (interface{}, error)
	// Set sets a CanAnnounce value to be at the specified index for the
	// property "canAnnounce". Panics if the index is out of bounds.
	// Invalidates all iterators.
	Set(idx int, v GoToSocialCanAnnounce)
	// SetIRI sets an IRI value to be at the specified index for the property
	// "canAnnounce". Panics if the index is out of bounds.
	SetIRI(idx int, v *url.URL)
	// SetType sets an arbitrary type value to the specified index of the
	// property "canAnnounce". Invalidates all iterators. Returns an error
	// if the type is not a valid one to set for this property. Panics if
	// the index is out of bounds.
	SetType(idx int, t Type) error
	// Swap swaps the location of values at two indices for the "canAnnounce"
	// property.
	Swap(i, j int)
}
