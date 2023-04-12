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

// checks if the POSTParcelLineItems201ResponseDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTParcelLineItems201ResponseDataAttributes{}

// POSTParcelLineItems201ResponseDataAttributes struct for POSTParcelLineItems201ResponseDataAttributes
type POSTParcelLineItems201ResponseDataAttributes struct {
	// The code of the SKU of the associated shipment_line_item.
	SkuCode interface{} `json:"sku_code,omitempty"`
	// The parcel line item quantity.
	Quantity interface{} `json:"quantity"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference interface{} `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin interface{} `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata interface{} `json:"metadata,omitempty"`
}

// NewPOSTParcelLineItems201ResponseDataAttributes instantiates a new POSTParcelLineItems201ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTParcelLineItems201ResponseDataAttributes(quantity interface{}) *POSTParcelLineItems201ResponseDataAttributes {
	this := POSTParcelLineItems201ResponseDataAttributes{}
	this.Quantity = quantity
	return &this
}

// NewPOSTParcelLineItems201ResponseDataAttributesWithDefaults instantiates a new POSTParcelLineItems201ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTParcelLineItems201ResponseDataAttributesWithDefaults() *POSTParcelLineItems201ResponseDataAttributes {
	this := POSTParcelLineItems201ResponseDataAttributes{}
	return &this
}

// GetSkuCode returns the SkuCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTParcelLineItems201ResponseDataAttributes) GetSkuCode() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.SkuCode
}

// GetSkuCodeOk returns a tuple with the SkuCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTParcelLineItems201ResponseDataAttributes) GetSkuCodeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.SkuCode) {
		return nil, false
	}
	return &o.SkuCode, true
}

// HasSkuCode returns a boolean if a field has been set.
func (o *POSTParcelLineItems201ResponseDataAttributes) HasSkuCode() bool {
	if o != nil && IsNil(o.SkuCode) {
		return true
	}

	return false
}

// SetSkuCode gets a reference to the given interface{} and assigns it to the SkuCode field.
func (o *POSTParcelLineItems201ResponseDataAttributes) SetSkuCode(v interface{}) {
	o.SkuCode = v
}

// GetQuantity returns the Quantity field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *POSTParcelLineItems201ResponseDataAttributes) GetQuantity() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTParcelLineItems201ResponseDataAttributes) GetQuantityOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return &o.Quantity, true
}

// SetQuantity sets field value
func (o *POSTParcelLineItems201ResponseDataAttributes) SetQuantity(v interface{}) {
	o.Quantity = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTParcelLineItems201ResponseDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTParcelLineItems201ResponseDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *POSTParcelLineItems201ResponseDataAttributes) HasReference() bool {
	if o != nil && IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *POSTParcelLineItems201ResponseDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTParcelLineItems201ResponseDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTParcelLineItems201ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReferenceOrigin) {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *POSTParcelLineItems201ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && IsNil(o.ReferenceOrigin) {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *POSTParcelLineItems201ResponseDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTParcelLineItems201ResponseDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTParcelLineItems201ResponseDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *POSTParcelLineItems201ResponseDataAttributes) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *POSTParcelLineItems201ResponseDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

func (o POSTParcelLineItems201ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTParcelLineItems201ResponseDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
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

type NullablePOSTParcelLineItems201ResponseDataAttributes struct {
	value *POSTParcelLineItems201ResponseDataAttributes
	isSet bool
}

func (v NullablePOSTParcelLineItems201ResponseDataAttributes) Get() *POSTParcelLineItems201ResponseDataAttributes {
	return v.value
}

func (v *NullablePOSTParcelLineItems201ResponseDataAttributes) Set(val *POSTParcelLineItems201ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTParcelLineItems201ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTParcelLineItems201ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTParcelLineItems201ResponseDataAttributes(val *POSTParcelLineItems201ResponseDataAttributes) *NullablePOSTParcelLineItems201ResponseDataAttributes {
	return &NullablePOSTParcelLineItems201ResponseDataAttributes{value: val, isSet: true}
}

func (v NullablePOSTParcelLineItems201ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTParcelLineItems201ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
