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

// checks if the POSTOrderCopiesRequestDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTOrderCopiesRequestDataAttributes{}

// POSTOrderCopiesRequestDataAttributes struct for POSTOrderCopiesRequestDataAttributes
type POSTOrderCopiesRequestDataAttributes struct {
	// Indicates if the target order must be placed upon copy.
	PlaceTargetOrder interface{} `json:"place_target_order,omitempty"`
	// Indicates if the payment source within the source order customer's wallet must be copied.
	ReuseWallet interface{} `json:"reuse_wallet,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference interface{} `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin interface{} `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata interface{} `json:"metadata,omitempty"`
	// Indicates if the source order must be cancelled upon copy.
	CancelSourceOrder interface{} `json:"cancel_source_order,omitempty"`
}

// NewPOSTOrderCopiesRequestDataAttributes instantiates a new POSTOrderCopiesRequestDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTOrderCopiesRequestDataAttributes() *POSTOrderCopiesRequestDataAttributes {
	this := POSTOrderCopiesRequestDataAttributes{}
	return &this
}

// NewPOSTOrderCopiesRequestDataAttributesWithDefaults instantiates a new POSTOrderCopiesRequestDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTOrderCopiesRequestDataAttributesWithDefaults() *POSTOrderCopiesRequestDataAttributes {
	this := POSTOrderCopiesRequestDataAttributes{}
	return &this
}

// GetPlaceTargetOrder returns the PlaceTargetOrder field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTOrderCopiesRequestDataAttributes) GetPlaceTargetOrder() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.PlaceTargetOrder
}

// GetPlaceTargetOrderOk returns a tuple with the PlaceTargetOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTOrderCopiesRequestDataAttributes) GetPlaceTargetOrderOk() (*interface{}, bool) {
	if o == nil || IsNil(o.PlaceTargetOrder) {
		return nil, false
	}
	return &o.PlaceTargetOrder, true
}

// HasPlaceTargetOrder returns a boolean if a field has been set.
func (o *POSTOrderCopiesRequestDataAttributes) HasPlaceTargetOrder() bool {
	if o != nil && IsNil(o.PlaceTargetOrder) {
		return true
	}

	return false
}

// SetPlaceTargetOrder gets a reference to the given interface{} and assigns it to the PlaceTargetOrder field.
func (o *POSTOrderCopiesRequestDataAttributes) SetPlaceTargetOrder(v interface{}) {
	o.PlaceTargetOrder = v
}

// GetReuseWallet returns the ReuseWallet field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTOrderCopiesRequestDataAttributes) GetReuseWallet() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReuseWallet
}

// GetReuseWalletOk returns a tuple with the ReuseWallet field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTOrderCopiesRequestDataAttributes) GetReuseWalletOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReuseWallet) {
		return nil, false
	}
	return &o.ReuseWallet, true
}

// HasReuseWallet returns a boolean if a field has been set.
func (o *POSTOrderCopiesRequestDataAttributes) HasReuseWallet() bool {
	if o != nil && IsNil(o.ReuseWallet) {
		return true
	}

	return false
}

// SetReuseWallet gets a reference to the given interface{} and assigns it to the ReuseWallet field.
func (o *POSTOrderCopiesRequestDataAttributes) SetReuseWallet(v interface{}) {
	o.ReuseWallet = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTOrderCopiesRequestDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTOrderCopiesRequestDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *POSTOrderCopiesRequestDataAttributes) HasReference() bool {
	if o != nil && IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *POSTOrderCopiesRequestDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTOrderCopiesRequestDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTOrderCopiesRequestDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReferenceOrigin) {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *POSTOrderCopiesRequestDataAttributes) HasReferenceOrigin() bool {
	if o != nil && IsNil(o.ReferenceOrigin) {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *POSTOrderCopiesRequestDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTOrderCopiesRequestDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTOrderCopiesRequestDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *POSTOrderCopiesRequestDataAttributes) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *POSTOrderCopiesRequestDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

// GetCancelSourceOrder returns the CancelSourceOrder field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTOrderCopiesRequestDataAttributes) GetCancelSourceOrder() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CancelSourceOrder
}

// GetCancelSourceOrderOk returns a tuple with the CancelSourceOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTOrderCopiesRequestDataAttributes) GetCancelSourceOrderOk() (*interface{}, bool) {
	if o == nil || IsNil(o.CancelSourceOrder) {
		return nil, false
	}
	return &o.CancelSourceOrder, true
}

// HasCancelSourceOrder returns a boolean if a field has been set.
func (o *POSTOrderCopiesRequestDataAttributes) HasCancelSourceOrder() bool {
	if o != nil && IsNil(o.CancelSourceOrder) {
		return true
	}

	return false
}

// SetCancelSourceOrder gets a reference to the given interface{} and assigns it to the CancelSourceOrder field.
func (o *POSTOrderCopiesRequestDataAttributes) SetCancelSourceOrder(v interface{}) {
	o.CancelSourceOrder = v
}

func (o POSTOrderCopiesRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTOrderCopiesRequestDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.PlaceTargetOrder != nil {
		toSerialize["place_target_order"] = o.PlaceTargetOrder
	}
	if o.ReuseWallet != nil {
		toSerialize["reuse_wallet"] = o.ReuseWallet
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
	if o.CancelSourceOrder != nil {
		toSerialize["cancel_source_order"] = o.CancelSourceOrder
	}
	return toSerialize, nil
}

type NullablePOSTOrderCopiesRequestDataAttributes struct {
	value *POSTOrderCopiesRequestDataAttributes
	isSet bool
}

func (v NullablePOSTOrderCopiesRequestDataAttributes) Get() *POSTOrderCopiesRequestDataAttributes {
	return v.value
}

func (v *NullablePOSTOrderCopiesRequestDataAttributes) Set(val *POSTOrderCopiesRequestDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTOrderCopiesRequestDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTOrderCopiesRequestDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTOrderCopiesRequestDataAttributes(val *POSTOrderCopiesRequestDataAttributes) *NullablePOSTOrderCopiesRequestDataAttributes {
	return &NullablePOSTOrderCopiesRequestDataAttributes{value: val, isSet: true}
}

func (v NullablePOSTOrderCopiesRequestDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTOrderCopiesRequestDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}