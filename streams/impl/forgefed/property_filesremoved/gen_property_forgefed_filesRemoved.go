// Code generated by astool. DO NOT EDIT.

package propertyfilesremoved

import (
	"fmt"
	string1 "github.com/go-fed/activity/streams/values/string"
	vocab "github.com/go-fed/activity/streams/vocab"
	"net/url"
)

// ForgeFedFilesRemovedPropertyIterator is an iterator for a property. It is
// permitted to be a single default-valued value type.
type ForgeFedFilesRemovedPropertyIterator struct {
	xmlschemaStringMember string
	hasStringMember       bool
	unknown               interface{}
	iri                   *url.URL
	alias                 string
	myIdx                 int
	parent                vocab.ForgeFedFilesRemovedProperty
}

// NewForgeFedFilesRemovedPropertyIterator creates a new ForgeFedFilesRemoved
// property.
func NewForgeFedFilesRemovedPropertyIterator() *ForgeFedFilesRemovedPropertyIterator {
	return &ForgeFedFilesRemovedPropertyIterator{alias: ""}
}

// deserializeForgeFedFilesRemovedPropertyIterator creates an iterator from an
// element that has been unmarshalled from a text or binary format.
func deserializeForgeFedFilesRemovedPropertyIterator(i interface{}, aliasMap map[string]string) (*ForgeFedFilesRemovedPropertyIterator, error) {
	alias := ""
	if a, ok := aliasMap["https://forgefed.peers.community/ns"]; ok {
		alias = a
	}
	if s, ok := i.(string); ok {
		u, err := url.Parse(s)
		// If error exists, don't error out -- skip this and treat as unknown string ([]byte) at worst
		// Also, if no scheme exists, don't treat it as a URL -- net/url is greedy
		if err == nil && len(u.Scheme) > 0 {
			this := &ForgeFedFilesRemovedPropertyIterator{
				alias: alias,
				iri:   u,
			}
			return this, nil
		}
	}
	if v, err := string1.DeserializeString(i); err == nil {
		this := &ForgeFedFilesRemovedPropertyIterator{
			alias:                 alias,
			hasStringMember:       true,
			xmlschemaStringMember: v,
		}
		return this, nil
	}
	this := &ForgeFedFilesRemovedPropertyIterator{
		alias:   alias,
		unknown: i,
	}
	return this, nil
}

// Get returns the value of this property. When IsXMLSchemaString returns false,
// Get will return any arbitrary value.
func (this ForgeFedFilesRemovedPropertyIterator) Get() string {
	return this.xmlschemaStringMember
}

// GetIRI returns the IRI of this property. When IsIRI returns false, GetIRI will
// return any arbitrary value.
func (this ForgeFedFilesRemovedPropertyIterator) GetIRI() *url.URL {
	return this.iri
}

// HasAny returns true if the value or IRI is set.
func (this ForgeFedFilesRemovedPropertyIterator) HasAny() bool {
	return this.IsXMLSchemaString() || this.iri != nil
}

// IsIRI returns true if this property is an IRI.
func (this ForgeFedFilesRemovedPropertyIterator) IsIRI() bool {
	return this.iri != nil
}

// IsXMLSchemaString returns true if this property is set and not an IRI.
func (this ForgeFedFilesRemovedPropertyIterator) IsXMLSchemaString() bool {
	return this.hasStringMember
}

// JSONLDContext returns the JSONLD URIs required in the context string for this
// property and the specific values that are set. The value in the map is the
// alias used to import the property's value or values.
func (this ForgeFedFilesRemovedPropertyIterator) JSONLDContext() map[string]string {
	m := map[string]string{"https://forgefed.peers.community/ns": this.alias}
	var child map[string]string

	/*
	   Since the literal maps in this function are determined at
	   code-generation time, this loop should not overwrite an existing key with a
	   new value.
	*/
	for k, v := range child {
		m[k] = v
	}
	return m
}

// KindIndex computes an arbitrary value for indexing this kind of value. This is
// a leaky API detail only for folks looking to replace the go-fed
// implementation. Applications should not use this method.
func (this ForgeFedFilesRemovedPropertyIterator) KindIndex() int {
	if this.IsXMLSchemaString() {
		return 0
	}
	if this.IsIRI() {
		return -2
	}
	return -1
}

// LessThan compares two instances of this property with an arbitrary but stable
// comparison. Applications should not use this because it is only meant to
// help alternative implementations to go-fed to be able to normalize
// nonfunctional properties.
func (this ForgeFedFilesRemovedPropertyIterator) LessThan(o vocab.ForgeFedFilesRemovedPropertyIterator) bool {
	// LessThan comparison for if either or both are IRIs.
	if this.IsIRI() && o.IsIRI() {
		return this.iri.String() < o.GetIRI().String()
	} else if this.IsIRI() {
		// IRIs are always less than other values, none, or unknowns
		return true
	} else if o.IsIRI() {
		// This other, none, or unknown value is always greater than IRIs
		return false
	}
	// LessThan comparison for the single value or unknown value.
	if !this.IsXMLSchemaString() && !o.IsXMLSchemaString() {
		// Both are unknowns.
		return false
	} else if this.IsXMLSchemaString() && !o.IsXMLSchemaString() {
		// Values are always greater than unknown values.
		return false
	} else if !this.IsXMLSchemaString() && o.IsXMLSchemaString() {
		// Unknowns are always less than known values.
		return true
	} else {
		// Actual comparison.
		return string1.LessString(this.Get(), o.Get())
	}
}

// Name returns the name of this property: "ForgeFedFilesRemoved".
func (this ForgeFedFilesRemovedPropertyIterator) Name() string {
	if len(this.alias) > 0 {
		return this.alias + ":" + "ForgeFedFilesRemoved"
	} else {
		return "ForgeFedFilesRemoved"
	}
}

// Next returns the next iterator, or nil if there is no next iterator.
func (this ForgeFedFilesRemovedPropertyIterator) Next() vocab.ForgeFedFilesRemovedPropertyIterator {
	if this.myIdx+1 >= this.parent.Len() {
		return nil
	} else {
		return this.parent.At(this.myIdx + 1)
	}
}

// Prev returns the previous iterator, or nil if there is no previous iterator.
func (this ForgeFedFilesRemovedPropertyIterator) Prev() vocab.ForgeFedFilesRemovedPropertyIterator {
	if this.myIdx-1 < 0 {
		return nil
	} else {
		return this.parent.At(this.myIdx - 1)
	}
}

// Set sets the value of this property. Calling IsXMLSchemaString afterwards will
// return true.
func (this *ForgeFedFilesRemovedPropertyIterator) Set(v string) {
	this.clear()
	this.xmlschemaStringMember = v
	this.hasStringMember = true
}

// SetIRI sets the value of this property. Calling IsIRI afterwards will return
// true.
func (this *ForgeFedFilesRemovedPropertyIterator) SetIRI(v *url.URL) {
	this.clear()
	this.iri = v
}

// clear ensures no value of this property is set. Calling IsXMLSchemaString
// afterwards will return false.
func (this *ForgeFedFilesRemovedPropertyIterator) clear() {
	this.unknown = nil
	this.iri = nil
	this.hasStringMember = false
}

// serialize converts this into an interface representation suitable for
// marshalling into a text or binary format. Applications should not need this
// function as most typical use cases serialize types instead of individual
// properties. It is exposed for alternatives to go-fed implementations to use.
func (this ForgeFedFilesRemovedPropertyIterator) serialize() (interface{}, error) {
	if this.IsXMLSchemaString() {
		return string1.SerializeString(this.Get())
	} else if this.IsIRI() {
		return this.iri.String(), nil
	}
	return this.unknown, nil
}

// ForgeFedFilesRemovedProperty is the non-functional property "filesRemoved". It
// is permitted to have one or more values, and of different value types.
type ForgeFedFilesRemovedProperty struct {
	properties []*ForgeFedFilesRemovedPropertyIterator
	alias      string
}

// DeserializeFilesRemovedProperty creates a "filesRemoved" property from an
// interface representation that has been unmarshalled from a text or binary
// format.
func DeserializeFilesRemovedProperty(m map[string]interface{}, aliasMap map[string]string) (vocab.ForgeFedFilesRemovedProperty, error) {
	alias := ""
	if a, ok := aliasMap["https://forgefed.peers.community/ns"]; ok {
		alias = a
	}
	propName := "filesRemoved"
	if len(alias) > 0 {
		propName = fmt.Sprintf("%s:%s", alias, "filesRemoved")
	}
	i, ok := m[propName]

	if ok {
		this := &ForgeFedFilesRemovedProperty{
			alias:      alias,
			properties: []*ForgeFedFilesRemovedPropertyIterator{},
		}
		if list, ok := i.([]interface{}); ok {
			for _, iterator := range list {
				if p, err := deserializeForgeFedFilesRemovedPropertyIterator(iterator, aliasMap); err != nil {
					return this, err
				} else if p != nil {
					this.properties = append(this.properties, p)
				}
			}
		} else {
			if p, err := deserializeForgeFedFilesRemovedPropertyIterator(i, aliasMap); err != nil {
				return this, err
			} else if p != nil {
				this.properties = append(this.properties, p)
			}
		}
		// Set up the properties for iteration.
		for idx, ele := range this.properties {
			ele.parent = this
			ele.myIdx = idx
		}
		return this, nil
	}
	return nil, nil
}

// NewForgeFedFilesRemovedProperty creates a new filesRemoved property.
func NewForgeFedFilesRemovedProperty() *ForgeFedFilesRemovedProperty {
	return &ForgeFedFilesRemovedProperty{alias: ""}
}

// AppendIRI appends an IRI value to the back of a list of the property
// "filesRemoved"
func (this *ForgeFedFilesRemovedProperty) AppendIRI(v *url.URL) {
	this.properties = append(this.properties, &ForgeFedFilesRemovedPropertyIterator{
		alias:  this.alias,
		iri:    v,
		myIdx:  this.Len(),
		parent: this,
	})
}

// AppendXMLSchemaString appends a string value to the back of a list of the
// property "filesRemoved". Invalidates iterators that are traversing using
// Prev.
func (this *ForgeFedFilesRemovedProperty) AppendXMLSchemaString(v string) {
	this.properties = append(this.properties, &ForgeFedFilesRemovedPropertyIterator{
		alias:                 this.alias,
		hasStringMember:       true,
		myIdx:                 this.Len(),
		parent:                this,
		xmlschemaStringMember: v,
	})
}

// At returns the property value for the specified index. Panics if the index is
// out of bounds.
func (this ForgeFedFilesRemovedProperty) At(index int) vocab.ForgeFedFilesRemovedPropertyIterator {
	return this.properties[index]
}

// Begin returns the first iterator, or nil if empty. Can be used with the
// iterator's Next method and this property's End method to iterate from front
// to back through all values.
func (this ForgeFedFilesRemovedProperty) Begin() vocab.ForgeFedFilesRemovedPropertyIterator {
	if this.Empty() {
		return nil
	} else {
		return this.properties[0]
	}
}

// Empty returns returns true if there are no elements.
func (this ForgeFedFilesRemovedProperty) Empty() bool {
	return this.Len() == 0
}

// End returns beyond-the-last iterator, which is nil. Can be used with the
// iterator's Next method and this property's Begin method to iterate from
// front to back through all values.
func (this ForgeFedFilesRemovedProperty) End() vocab.ForgeFedFilesRemovedPropertyIterator {
	return nil
}

// Insert inserts an IRI value at the specified index for a property
// "filesRemoved". Existing elements at that index and higher are shifted back
// once. Invalidates all iterators.
func (this *ForgeFedFilesRemovedProperty) InsertIRI(idx int, v *url.URL) {
	this.properties = append(this.properties, nil)
	copy(this.properties[idx+1:], this.properties[idx:])
	this.properties[idx] = &ForgeFedFilesRemovedPropertyIterator{
		alias:  this.alias,
		iri:    v,
		myIdx:  idx,
		parent: this,
	}
	for i := idx; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// InsertXMLSchemaString inserts a string value at the specified index for a
// property "filesRemoved". Existing elements at that index and higher are
// shifted back once. Invalidates all iterators.
func (this *ForgeFedFilesRemovedProperty) InsertXMLSchemaString(idx int, v string) {
	this.properties = append(this.properties, nil)
	copy(this.properties[idx+1:], this.properties[idx:])
	this.properties[idx] = &ForgeFedFilesRemovedPropertyIterator{
		alias:                 this.alias,
		hasStringMember:       true,
		myIdx:                 idx,
		parent:                this,
		xmlschemaStringMember: v,
	}
	for i := idx; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// JSONLDContext returns the JSONLD URIs required in the context string for this
// property and the specific values that are set. The value in the map is the
// alias used to import the property's value or values.
func (this ForgeFedFilesRemovedProperty) JSONLDContext() map[string]string {
	m := map[string]string{"https://forgefed.peers.community/ns": this.alias}
	for _, elem := range this.properties {
		child := elem.JSONLDContext()
		/*
		   Since the literal maps in this function are determined at
		   code-generation time, this loop should not overwrite an existing key with a
		   new value.
		*/
		for k, v := range child {
			m[k] = v
		}
	}
	return m
}

// KindIndex computes an arbitrary value for indexing this kind of value. This is
// a leaky API method specifically needed only for alternate implementations
// for go-fed. Applications should not use this method. Panics if the index is
// out of bounds.
func (this ForgeFedFilesRemovedProperty) KindIndex(idx int) int {
	return this.properties[idx].KindIndex()
}

// Len returns the number of values that exist for the "filesRemoved" property.
func (this ForgeFedFilesRemovedProperty) Len() (length int) {
	return len(this.properties)
}

// Less computes whether another property is less than this one. Mixing types
// results in a consistent but arbitrary ordering
func (this ForgeFedFilesRemovedProperty) Less(i, j int) bool {
	idx1 := this.KindIndex(i)
	idx2 := this.KindIndex(j)
	if idx1 < idx2 {
		return true
	} else if idx1 == idx2 {
		if idx1 == 0 {
			lhs := this.properties[i].Get()
			rhs := this.properties[j].Get()
			return string1.LessString(lhs, rhs)
		} else if idx1 == -2 {
			lhs := this.properties[i].GetIRI()
			rhs := this.properties[j].GetIRI()
			return lhs.String() < rhs.String()
		}
	}
	return false
}

// LessThan compares two instances of this property with an arbitrary but stable
// comparison. Applications should not use this because it is only meant to
// help alternative implementations to go-fed to be able to normalize
// nonfunctional properties.
func (this ForgeFedFilesRemovedProperty) LessThan(o vocab.ForgeFedFilesRemovedProperty) bool {
	l1 := this.Len()
	l2 := o.Len()
	l := l1
	if l2 < l1 {
		l = l2
	}
	for i := 0; i < l; i++ {
		if this.properties[i].LessThan(o.At(i)) {
			return true
		} else if o.At(i).LessThan(this.properties[i]) {
			return false
		}
	}
	return l1 < l2
}

// Name returns the name of this property ("filesRemoved") with any alias.
func (this ForgeFedFilesRemovedProperty) Name() string {
	if len(this.alias) > 0 {
		return this.alias + ":" + "filesRemoved"
	} else {
		return "filesRemoved"
	}
}

// PrependIRI prepends an IRI value to the front of a list of the property
// "filesRemoved".
func (this *ForgeFedFilesRemovedProperty) PrependIRI(v *url.URL) {
	this.properties = append([]*ForgeFedFilesRemovedPropertyIterator{{
		alias:  this.alias,
		iri:    v,
		myIdx:  0,
		parent: this,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// PrependXMLSchemaString prepends a string value to the front of a list of the
// property "filesRemoved". Invalidates all iterators.
func (this *ForgeFedFilesRemovedProperty) PrependXMLSchemaString(v string) {
	this.properties = append([]*ForgeFedFilesRemovedPropertyIterator{{
		alias:                 this.alias,
		hasStringMember:       true,
		myIdx:                 0,
		parent:                this,
		xmlschemaStringMember: v,
	}}, this.properties...)
	for i := 1; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// Remove deletes an element at the specified index from a list of the property
// "filesRemoved", regardless of its type. Panics if the index is out of
// bounds. Invalidates all iterators.
func (this *ForgeFedFilesRemovedProperty) Remove(idx int) {
	(this.properties)[idx].parent = nil
	copy((this.properties)[idx:], (this.properties)[idx+1:])
	(this.properties)[len(this.properties)-1] = &ForgeFedFilesRemovedPropertyIterator{}
	this.properties = (this.properties)[:len(this.properties)-1]
	for i := idx; i < this.Len(); i++ {
		(this.properties)[i].myIdx = i
	}
}

// Serialize converts this into an interface representation suitable for
// marshalling into a text or binary format. Applications should not need this
// function as most typical use cases serialize types instead of individual
// properties. It is exposed for alternatives to go-fed implementations to use.
func (this ForgeFedFilesRemovedProperty) Serialize() (interface{}, error) {
	s := make([]interface{}, 0, len(this.properties))
	for _, iterator := range this.properties {
		if b, err := iterator.serialize(); err != nil {
			return s, err
		} else {
			s = append(s, b)
		}
	}
	// Shortcut: if serializing one value, don't return an array -- pretty sure other Fediverse software would choke on a "type" value with array, for example.
	if len(s) == 1 {
		return s[0], nil
	}
	return s, nil
}

// Set sets a string value to be at the specified index for the property
// "filesRemoved". Panics if the index is out of bounds. Invalidates all
// iterators.
func (this *ForgeFedFilesRemovedProperty) Set(idx int, v string) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &ForgeFedFilesRemovedPropertyIterator{
		alias:                 this.alias,
		hasStringMember:       true,
		myIdx:                 idx,
		parent:                this,
		xmlschemaStringMember: v,
	}
}

// SetIRI sets an IRI value to be at the specified index for the property
// "filesRemoved". Panics if the index is out of bounds.
func (this *ForgeFedFilesRemovedProperty) SetIRI(idx int, v *url.URL) {
	(this.properties)[idx].parent = nil
	(this.properties)[idx] = &ForgeFedFilesRemovedPropertyIterator{
		alias:  this.alias,
		iri:    v,
		myIdx:  idx,
		parent: this,
	}
}

// Swap swaps the location of values at two indices for the "filesRemoved"
// property.
func (this ForgeFedFilesRemovedProperty) Swap(i, j int) {
	this.properties[i], this.properties[j] = this.properties[j], this.properties[i]
}