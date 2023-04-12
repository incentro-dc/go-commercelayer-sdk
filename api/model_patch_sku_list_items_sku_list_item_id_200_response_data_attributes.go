/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 4.1.3
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the PATCHSkuListItemsSkuListItemId200ResponseDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PATCHSkuListItemsSkuListItemId200ResponseDataAttributes{}

// PATCHSkuListItemsSkuListItemId200ResponseDataAttributes struct for PATCHSkuListItemsSkuListItemId200ResponseDataAttributes
type PATCHSkuListItemsSkuListItemId200ResponseDataAttributes struct {
	// The SKU list item's position.
	Position interface{} `json:"position,omitempty"`
	// The code of the associated SKU.
	SkuCode interface{} `json:"sku_code,omitempty"`
	// The SKU quantity for this SKU list item.
	Quantity interface{} `json:"quantity,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference interface{} `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin interface{} `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata interface{} `json:"metadata,omitempty"`
}

// NewPATCHSkuListItemsSkuListItemId200ResponseDataAttributes instantiates a new PATCHSkuListItemsSkuListItemId200ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHSkuListItemsSkuListItemId200ResponseDataAttributes() *PATCHSkuListItemsSkuListItemId200ResponseDataAttributes {
	this := PATCHSkuListItemsSkuListItemId200ResponseDataAttributes{}
	return &this
}

// NewPATCHSkuListItemsSkuListItemId200ResponseDataAttributesWithDefaults instantiates a new PATCHSkuListItemsSkuListItemId200ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHSkuListItemsSkuListItemId200ResponseDataAttributesWithDefaults() *PATCHSkuListItemsSkuListItemId200ResponseDataAttributes {
	this := PATCHSkuListItemsSkuListItemId200ResponseDataAttributes{}
	return &this
}

// GetPosition returns the Position field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHSkuListItemsSkuListItemId200ResponseDataAttributes) GetPosition() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Position
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHSkuListItemsSkuListItemId200ResponseDataAttributes) GetPositionOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Position) {
		return nil, false
	}
	return &o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *PATCHSkuListItemsSkuListItemId200ResponseDataAttributes) HasPosition() bool {
	if o != nil && IsNil(o.Position) {
		return true
	}

	return false
}

// SetPosition gets a reference to the given interface{} and assigns it to the Position field.
func (o *PATCHSkuListItemsSkuListItemId200ResponseDataAttributes) SetPosition(v interface{}) {
	o.Position = v
}

// GetSkuCode returns the SkuCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHSkuListItemsSkuListItemId200ResponseDataAttributes) GetSkuCode() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.SkuCode
}

// GetSkuCodeOk returns a tuple with the SkuCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHSkuListItemsSkuListItemId200ResponseDataAttributes) GetSkuCodeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.SkuCode) {
		return nil, false
	}
	return &o.SkuCode, true
}

// HasSkuCode returns a boolean if a field has been set.
func (o *PATCHSkuListItemsSkuListItemId200ResponseDataAttributes) HasSkuCode() bool {
	if o != nil && IsNil(o.SkuCode) {
		return true
	}

	return false
}

// SetSkuCode gets a reference to the given interface{} and assigns it to the SkuCode field.
func (o *PATCHSkuListItemsSkuListItemId200ResponseDataAttributes) SetSkuCode(v interface{}) {
	o.SkuCode = v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHSkuListItemsSkuListItemId200ResponseDataAttributes) GetQuantity() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHSkuListItemsSkuListItemId200ResponseDataAttributes) GetQuantityOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return &o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *PATCHSkuListItemsSkuListItemId200ResponseDataAttributes) HasQuantity() bool {
	if o != nil && IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given interface{} and assigns it to the Quantity field.
func (o *PATCHSkuListItemsSkuListItemId200ResponseDataAttributes) SetQuantity(v interface{}) {
	o.Quantity = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHSkuListItemsSkuListItemId200ResponseDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHSkuListItemsSkuListItemId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *PATCHSkuListItemsSkuListItemId200ResponseDataAttributes) HasReference() bool {
	if o != nil && IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *PATCHSkuListItemsSkuListItemId200ResponseDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHSkuListItemsSkuListItemId200ResponseDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHSkuListItemsSkuListItemId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReferenceOrigin) {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *PATCHSkuListItemsSkuListItemId200ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && IsNil(o.ReferenceOrigin) {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *PATCHSkuListItemsSkuListItemId200ResponseDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHSkuListItemsSkuListItemId200ResponseDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHSkuListItemsSkuListItemId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *PATCHSkuListItemsSkuListItemId200ResponseDataAttributes) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *PATCHSkuListItemsSkuListItemId200ResponseDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

func (o PATCHSkuListItemsSkuListItemId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PATCHSkuListItemsSkuListItemId200ResponseDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Position != nil {
		toSerialize["position"] = o.Position
	}
	if o.SkuCode != nil {
		toSerialize["sku_code"] = o.SkuCode
	}
	if o.Quantity != nil {
		toSerialize["quantity"] = o.Quantity
	}
	if o.Reference != nil {
		toSerialize["reference"] = o.Reference
	}
	if o.ReferenceOrigin != nil {
		toSerialize["reference_origin"] = o.ReferenceOrigin
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	return toSerialize, nil
}

type NullablePATCHSkuListItemsSkuListItemId200ResponseDataAttributes struct {
	value *PATCHSkuListItemsSkuListItemId200ResponseDataAttributes
	isSet bool
}

func (v NullablePATCHSkuListItemsSkuListItemId200ResponseDataAttributes) Get() *PATCHSkuListItemsSkuListItemId200ResponseDataAttributes {
	return v.value
}

func (v *NullablePATCHSkuListItemsSkuListItemId200ResponseDataAttributes) Set(val *PATCHSkuListItemsSkuListItemId200ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHSkuListItemsSkuListItemId200ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHSkuListItemsSkuListItemId200ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHSkuListItemsSkuListItemId200ResponseDataAttributes(val *PATCHSkuListItemsSkuListItemId200ResponseDataAttributes) *NullablePATCHSkuListItemsSkuListItemId200ResponseDataAttributes {
	return &NullablePATCHSkuListItemsSkuListItemId200ResponseDataAttributes{value: val, isSet: true}
}

func (v NullablePATCHSkuListItemsSkuListItemId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHSkuListItemsSkuListItemId200ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
