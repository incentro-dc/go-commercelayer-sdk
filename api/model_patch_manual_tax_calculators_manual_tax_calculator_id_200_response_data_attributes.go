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

// PATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseDataAttributes struct for PATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseDataAttributes
type PATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseDataAttributes struct {
	// The tax calculator's internal name.
	Name interface{} `json:"name,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference interface{} `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin interface{} `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata interface{} `json:"metadata,omitempty"`
}

// NewPATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseDataAttributes instantiates a new PATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseDataAttributes() *PATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseDataAttributes {
	this := PATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseDataAttributes{}
	return &this
}

// NewPATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseDataAttributesWithDefaults instantiates a new PATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseDataAttributesWithDefaults() *PATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseDataAttributes {
	this := PATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseDataAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseDataAttributes) GetName() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseDataAttributes) GetNameOk() (*interface{}, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return &o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseDataAttributes) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given interface{} and assigns it to the Name field.
func (o *PATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseDataAttributes) SetName(v interface{}) {
	o.Name = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *PATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseDataAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *PATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *PATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *PATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *PATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseDataAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *PATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

func (o PATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
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

type NullablePATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseDataAttributes struct {
	value *PATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseDataAttributes
	isSet bool
}

func (v NullablePATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseDataAttributes) Get() *PATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseDataAttributes {
	return v.value
}

func (v *NullablePATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseDataAttributes) Set(val *PATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseDataAttributes(val *PATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseDataAttributes) *NullablePATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseDataAttributes {
	return &NullablePATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseDataAttributes{value: val, isSet: true}
}

func (v NullablePATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHManualTaxCalculatorsManualTaxCalculatorId200ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
