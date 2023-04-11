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

// POSTOrderSubscriptionItems201ResponseDataAttributes struct for POSTOrderSubscriptionItems201ResponseDataAttributes
type POSTOrderSubscriptionItems201ResponseDataAttributes struct {
	// The code of the associated SKU.
	SkuCode interface{} `json:"sku_code,omitempty"`
	// The code of the associated bundle.
	BundleCode interface{} `json:"bundle_code,omitempty"`
	// The subscription item quantity.
	Quantity interface{} `json:"quantity"`
	// The unit amount of the subscription item, in cents.
	UnitAmountCents interface{} `json:"unit_amount_cents,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference interface{} `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin interface{} `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata interface{} `json:"metadata,omitempty"`
}

// NewPOSTOrderSubscriptionItems201ResponseDataAttributes instantiates a new POSTOrderSubscriptionItems201ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTOrderSubscriptionItems201ResponseDataAttributes(quantity interface{}) *POSTOrderSubscriptionItems201ResponseDataAttributes {
	this := POSTOrderSubscriptionItems201ResponseDataAttributes{}
	this.Quantity = quantity
	return &this
}

// NewPOSTOrderSubscriptionItems201ResponseDataAttributesWithDefaults instantiates a new POSTOrderSubscriptionItems201ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTOrderSubscriptionItems201ResponseDataAttributesWithDefaults() *POSTOrderSubscriptionItems201ResponseDataAttributes {
	this := POSTOrderSubscriptionItems201ResponseDataAttributes{}
	return &this
}

// GetSkuCode returns the SkuCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTOrderSubscriptionItems201ResponseDataAttributes) GetSkuCode() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.SkuCode
}

// GetSkuCodeOk returns a tuple with the SkuCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTOrderSubscriptionItems201ResponseDataAttributes) GetSkuCodeOk() (*interface{}, bool) {
	if o == nil || o.SkuCode == nil {
		return nil, false
	}
	return &o.SkuCode, true
}

// HasSkuCode returns a boolean if a field has been set.
func (o *POSTOrderSubscriptionItems201ResponseDataAttributes) HasSkuCode() bool {
	if o != nil && o.SkuCode != nil {
		return true
	}

	return false
}

// SetSkuCode gets a reference to the given interface{} and assigns it to the SkuCode field.
func (o *POSTOrderSubscriptionItems201ResponseDataAttributes) SetSkuCode(v interface{}) {
	o.SkuCode = v
}

// GetBundleCode returns the BundleCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTOrderSubscriptionItems201ResponseDataAttributes) GetBundleCode() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.BundleCode
}

// GetBundleCodeOk returns a tuple with the BundleCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTOrderSubscriptionItems201ResponseDataAttributes) GetBundleCodeOk() (*interface{}, bool) {
	if o == nil || o.BundleCode == nil {
		return nil, false
	}
	return &o.BundleCode, true
}

// HasBundleCode returns a boolean if a field has been set.
func (o *POSTOrderSubscriptionItems201ResponseDataAttributes) HasBundleCode() bool {
	if o != nil && o.BundleCode != nil {
		return true
	}

	return false
}

// SetBundleCode gets a reference to the given interface{} and assigns it to the BundleCode field.
func (o *POSTOrderSubscriptionItems201ResponseDataAttributes) SetBundleCode(v interface{}) {
	o.BundleCode = v
}

// GetQuantity returns the Quantity field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *POSTOrderSubscriptionItems201ResponseDataAttributes) GetQuantity() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTOrderSubscriptionItems201ResponseDataAttributes) GetQuantityOk() (*interface{}, bool) {
	if o == nil || o.Quantity == nil {
		return nil, false
	}
	return &o.Quantity, true
}

// SetQuantity sets field value
func (o *POSTOrderSubscriptionItems201ResponseDataAttributes) SetQuantity(v interface{}) {
	o.Quantity = v
}

// GetUnitAmountCents returns the UnitAmountCents field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTOrderSubscriptionItems201ResponseDataAttributes) GetUnitAmountCents() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.UnitAmountCents
}

// GetUnitAmountCentsOk returns a tuple with the UnitAmountCents field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTOrderSubscriptionItems201ResponseDataAttributes) GetUnitAmountCentsOk() (*interface{}, bool) {
	if o == nil || o.UnitAmountCents == nil {
		return nil, false
	}
	return &o.UnitAmountCents, true
}

// HasUnitAmountCents returns a boolean if a field has been set.
func (o *POSTOrderSubscriptionItems201ResponseDataAttributes) HasUnitAmountCents() bool {
	if o != nil && o.UnitAmountCents != nil {
		return true
	}

	return false
}

// SetUnitAmountCents gets a reference to the given interface{} and assigns it to the UnitAmountCents field.
func (o *POSTOrderSubscriptionItems201ResponseDataAttributes) SetUnitAmountCents(v interface{}) {
	o.UnitAmountCents = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTOrderSubscriptionItems201ResponseDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTOrderSubscriptionItems201ResponseDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *POSTOrderSubscriptionItems201ResponseDataAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *POSTOrderSubscriptionItems201ResponseDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTOrderSubscriptionItems201ResponseDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTOrderSubscriptionItems201ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *POSTOrderSubscriptionItems201ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *POSTOrderSubscriptionItems201ResponseDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTOrderSubscriptionItems201ResponseDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTOrderSubscriptionItems201ResponseDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *POSTOrderSubscriptionItems201ResponseDataAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *POSTOrderSubscriptionItems201ResponseDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

func (o POSTOrderSubscriptionItems201ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SkuCode != nil {
		toSerialize["sku_code"] = o.SkuCode
	}
	if o.BundleCode != nil {
		toSerialize["bundle_code"] = o.BundleCode
	}
	if o.Quantity != nil {
		toSerialize["quantity"] = o.Quantity
	}
	if o.UnitAmountCents != nil {
		toSerialize["unit_amount_cents"] = o.UnitAmountCents
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
	return json.Marshal(toSerialize)
}

type NullablePOSTOrderSubscriptionItems201ResponseDataAttributes struct {
	value *POSTOrderSubscriptionItems201ResponseDataAttributes
	isSet bool
}

func (v NullablePOSTOrderSubscriptionItems201ResponseDataAttributes) Get() *POSTOrderSubscriptionItems201ResponseDataAttributes {
	return v.value
}

func (v *NullablePOSTOrderSubscriptionItems201ResponseDataAttributes) Set(val *POSTOrderSubscriptionItems201ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTOrderSubscriptionItems201ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTOrderSubscriptionItems201ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTOrderSubscriptionItems201ResponseDataAttributes(val *POSTOrderSubscriptionItems201ResponseDataAttributes) *NullablePOSTOrderSubscriptionItems201ResponseDataAttributes {
	return &NullablePOSTOrderSubscriptionItems201ResponseDataAttributes{value: val, isSet: true}
}

func (v NullablePOSTOrderSubscriptionItems201ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTOrderSubscriptionItems201ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
