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

// GETInventoryStockLocations200ResponseDataInnerAttributes struct for GETInventoryStockLocations200ResponseDataInnerAttributes
type GETInventoryStockLocations200ResponseDataInnerAttributes struct {
	// The stock location priority within the associated invetory model.
	Priority interface{} `json:"priority,omitempty"`
	// Indicates if the shipment should be put on hold if fulfilled from the associated stock location. This is useful to manage use cases like back-orders, pre-orders or personalized orders that need to be customized before being fulfilled.
	OnHold interface{} `json:"on_hold,omitempty"`
	// Time at which the resource was created.
	CreatedAt interface{} `json:"created_at,omitempty"`
	// Time at which the resource was last updated.
	UpdatedAt interface{} `json:"updated_at,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference interface{} `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin interface{} `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata interface{} `json:"metadata,omitempty"`
}

// NewGETInventoryStockLocations200ResponseDataInnerAttributes instantiates a new GETInventoryStockLocations200ResponseDataInnerAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETInventoryStockLocations200ResponseDataInnerAttributes() *GETInventoryStockLocations200ResponseDataInnerAttributes {
	this := GETInventoryStockLocations200ResponseDataInnerAttributes{}
	return &this
}

// NewGETInventoryStockLocations200ResponseDataInnerAttributesWithDefaults instantiates a new GETInventoryStockLocations200ResponseDataInnerAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETInventoryStockLocations200ResponseDataInnerAttributesWithDefaults() *GETInventoryStockLocations200ResponseDataInnerAttributes {
	this := GETInventoryStockLocations200ResponseDataInnerAttributes{}
	return &this
}

// GetPriority returns the Priority field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETInventoryStockLocations200ResponseDataInnerAttributes) GetPriority() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETInventoryStockLocations200ResponseDataInnerAttributes) GetPriorityOk() (*interface{}, bool) {
	if o == nil || o.Priority == nil {
		return nil, false
	}
	return &o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *GETInventoryStockLocations200ResponseDataInnerAttributes) HasPriority() bool {
	if o != nil && o.Priority != nil {
		return true
	}

	return false
}

// SetPriority gets a reference to the given interface{} and assigns it to the Priority field.
func (o *GETInventoryStockLocations200ResponseDataInnerAttributes) SetPriority(v interface{}) {
	o.Priority = v
}

// GetOnHold returns the OnHold field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETInventoryStockLocations200ResponseDataInnerAttributes) GetOnHold() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.OnHold
}

// GetOnHoldOk returns a tuple with the OnHold field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETInventoryStockLocations200ResponseDataInnerAttributes) GetOnHoldOk() (*interface{}, bool) {
	if o == nil || o.OnHold == nil {
		return nil, false
	}
	return &o.OnHold, true
}

// HasOnHold returns a boolean if a field has been set.
func (o *GETInventoryStockLocations200ResponseDataInnerAttributes) HasOnHold() bool {
	if o != nil && o.OnHold != nil {
		return true
	}

	return false
}

// SetOnHold gets a reference to the given interface{} and assigns it to the OnHold field.
func (o *GETInventoryStockLocations200ResponseDataInnerAttributes) SetOnHold(v interface{}) {
	o.OnHold = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETInventoryStockLocations200ResponseDataInnerAttributes) GetCreatedAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETInventoryStockLocations200ResponseDataInnerAttributes) GetCreatedAtOk() (*interface{}, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GETInventoryStockLocations200ResponseDataInnerAttributes) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given interface{} and assigns it to the CreatedAt field.
func (o *GETInventoryStockLocations200ResponseDataInnerAttributes) SetCreatedAt(v interface{}) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETInventoryStockLocations200ResponseDataInnerAttributes) GetUpdatedAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETInventoryStockLocations200ResponseDataInnerAttributes) GetUpdatedAtOk() (*interface{}, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *GETInventoryStockLocations200ResponseDataInnerAttributes) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given interface{} and assigns it to the UpdatedAt field.
func (o *GETInventoryStockLocations200ResponseDataInnerAttributes) SetUpdatedAt(v interface{}) {
	o.UpdatedAt = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETInventoryStockLocations200ResponseDataInnerAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETInventoryStockLocations200ResponseDataInnerAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *GETInventoryStockLocations200ResponseDataInnerAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *GETInventoryStockLocations200ResponseDataInnerAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETInventoryStockLocations200ResponseDataInnerAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETInventoryStockLocations200ResponseDataInnerAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *GETInventoryStockLocations200ResponseDataInnerAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *GETInventoryStockLocations200ResponseDataInnerAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETInventoryStockLocations200ResponseDataInnerAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETInventoryStockLocations200ResponseDataInnerAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GETInventoryStockLocations200ResponseDataInnerAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *GETInventoryStockLocations200ResponseDataInnerAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

func (o GETInventoryStockLocations200ResponseDataInnerAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Priority != nil {
		toSerialize["priority"] = o.Priority
	}
	if o.OnHold != nil {
		toSerialize["on_hold"] = o.OnHold
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
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

type NullableGETInventoryStockLocations200ResponseDataInnerAttributes struct {
	value *GETInventoryStockLocations200ResponseDataInnerAttributes
	isSet bool
}

func (v NullableGETInventoryStockLocations200ResponseDataInnerAttributes) Get() *GETInventoryStockLocations200ResponseDataInnerAttributes {
	return v.value
}

func (v *NullableGETInventoryStockLocations200ResponseDataInnerAttributes) Set(val *GETInventoryStockLocations200ResponseDataInnerAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableGETInventoryStockLocations200ResponseDataInnerAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableGETInventoryStockLocations200ResponseDataInnerAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETInventoryStockLocations200ResponseDataInnerAttributes(val *GETInventoryStockLocations200ResponseDataInnerAttributes) *NullableGETInventoryStockLocations200ResponseDataInnerAttributes {
	return &NullableGETInventoryStockLocations200ResponseDataInnerAttributes{value: val, isSet: true}
}

func (v NullableGETInventoryStockLocations200ResponseDataInnerAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETInventoryStockLocations200ResponseDataInnerAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
