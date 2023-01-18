/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.4
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// GETStockLocations200ResponseDataInnerAttributes struct for GETStockLocations200ResponseDataInnerAttributes
type GETStockLocations200ResponseDataInnerAttributes struct {
	// Unique identifier for the stock location (numeric)
	Number *int32 `json:"number,omitempty"`
	// The stock location's internal name.
	Name *string `json:"name,omitempty"`
	// The shipping label format for this stock location. Can be one of 'PDF', 'ZPL', 'EPL2', or 'PNG'
	LabelFormat *string `json:"label_format,omitempty"`
	// Flag it if you want to skip the electronic invoice creation when generating the customs info for this stock location shipments.
	SuppressEtd *bool `json:"suppress_etd,omitempty"`
	// Time at which the resource was created.
	CreatedAt *string `json:"created_at,omitempty"`
	// Time at which the resource was last updated.
	UpdatedAt *string `json:"updated_at,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference *string `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin *string `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// NewGETStockLocations200ResponseDataInnerAttributes instantiates a new GETStockLocations200ResponseDataInnerAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETStockLocations200ResponseDataInnerAttributes() *GETStockLocations200ResponseDataInnerAttributes {
	this := GETStockLocations200ResponseDataInnerAttributes{}
	return &this
}

// NewGETStockLocations200ResponseDataInnerAttributesWithDefaults instantiates a new GETStockLocations200ResponseDataInnerAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETStockLocations200ResponseDataInnerAttributesWithDefaults() *GETStockLocations200ResponseDataInnerAttributes {
	this := GETStockLocations200ResponseDataInnerAttributes{}
	return &this
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *GETStockLocations200ResponseDataInnerAttributes) GetNumber() int32 {
	if o == nil || o.Number == nil {
		var ret int32
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETStockLocations200ResponseDataInnerAttributes) GetNumberOk() (*int32, bool) {
	if o == nil || o.Number == nil {
		return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *GETStockLocations200ResponseDataInnerAttributes) HasNumber() bool {
	if o != nil && o.Number != nil {
		return true
	}

	return false
}

// SetNumber gets a reference to the given int32 and assigns it to the Number field.
func (o *GETStockLocations200ResponseDataInnerAttributes) SetNumber(v int32) {
	o.Number = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GETStockLocations200ResponseDataInnerAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETStockLocations200ResponseDataInnerAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GETStockLocations200ResponseDataInnerAttributes) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GETStockLocations200ResponseDataInnerAttributes) SetName(v string) {
	o.Name = &v
}

// GetLabelFormat returns the LabelFormat field value if set, zero value otherwise.
func (o *GETStockLocations200ResponseDataInnerAttributes) GetLabelFormat() string {
	if o == nil || o.LabelFormat == nil {
		var ret string
		return ret
	}
	return *o.LabelFormat
}

// GetLabelFormatOk returns a tuple with the LabelFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETStockLocations200ResponseDataInnerAttributes) GetLabelFormatOk() (*string, bool) {
	if o == nil || o.LabelFormat == nil {
		return nil, false
	}
	return o.LabelFormat, true
}

// HasLabelFormat returns a boolean if a field has been set.
func (o *GETStockLocations200ResponseDataInnerAttributes) HasLabelFormat() bool {
	if o != nil && o.LabelFormat != nil {
		return true
	}

	return false
}

// SetLabelFormat gets a reference to the given string and assigns it to the LabelFormat field.
func (o *GETStockLocations200ResponseDataInnerAttributes) SetLabelFormat(v string) {
	o.LabelFormat = &v
}

// GetSuppressEtd returns the SuppressEtd field value if set, zero value otherwise.
func (o *GETStockLocations200ResponseDataInnerAttributes) GetSuppressEtd() bool {
	if o == nil || o.SuppressEtd == nil {
		var ret bool
		return ret
	}
	return *o.SuppressEtd
}

// GetSuppressEtdOk returns a tuple with the SuppressEtd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETStockLocations200ResponseDataInnerAttributes) GetSuppressEtdOk() (*bool, bool) {
	if o == nil || o.SuppressEtd == nil {
		return nil, false
	}
	return o.SuppressEtd, true
}

// HasSuppressEtd returns a boolean if a field has been set.
func (o *GETStockLocations200ResponseDataInnerAttributes) HasSuppressEtd() bool {
	if o != nil && o.SuppressEtd != nil {
		return true
	}

	return false
}

// SetSuppressEtd gets a reference to the given bool and assigns it to the SuppressEtd field.
func (o *GETStockLocations200ResponseDataInnerAttributes) SetSuppressEtd(v bool) {
	o.SuppressEtd = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *GETStockLocations200ResponseDataInnerAttributes) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETStockLocations200ResponseDataInnerAttributes) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GETStockLocations200ResponseDataInnerAttributes) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *GETStockLocations200ResponseDataInnerAttributes) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *GETStockLocations200ResponseDataInnerAttributes) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETStockLocations200ResponseDataInnerAttributes) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *GETStockLocations200ResponseDataInnerAttributes) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *GETStockLocations200ResponseDataInnerAttributes) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *GETStockLocations200ResponseDataInnerAttributes) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETStockLocations200ResponseDataInnerAttributes) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *GETStockLocations200ResponseDataInnerAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *GETStockLocations200ResponseDataInnerAttributes) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise.
func (o *GETStockLocations200ResponseDataInnerAttributes) GetReferenceOrigin() string {
	if o == nil || o.ReferenceOrigin == nil {
		var ret string
		return ret
	}
	return *o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETStockLocations200ResponseDataInnerAttributes) GetReferenceOriginOk() (*string, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *GETStockLocations200ResponseDataInnerAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given string and assigns it to the ReferenceOrigin field.
func (o *GETStockLocations200ResponseDataInnerAttributes) SetReferenceOrigin(v string) {
	o.ReferenceOrigin = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *GETStockLocations200ResponseDataInnerAttributes) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETStockLocations200ResponseDataInnerAttributes) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GETStockLocations200ResponseDataInnerAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *GETStockLocations200ResponseDataInnerAttributes) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o GETStockLocations200ResponseDataInnerAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Number != nil {
		toSerialize["number"] = o.Number
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.LabelFormat != nil {
		toSerialize["label_format"] = o.LabelFormat
	}
	if o.SuppressEtd != nil {
		toSerialize["suppress_etd"] = o.SuppressEtd
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

type NullableGETStockLocations200ResponseDataInnerAttributes struct {
	value *GETStockLocations200ResponseDataInnerAttributes
	isSet bool
}

func (v NullableGETStockLocations200ResponseDataInnerAttributes) Get() *GETStockLocations200ResponseDataInnerAttributes {
	return v.value
}

func (v *NullableGETStockLocations200ResponseDataInnerAttributes) Set(val *GETStockLocations200ResponseDataInnerAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableGETStockLocations200ResponseDataInnerAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableGETStockLocations200ResponseDataInnerAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETStockLocations200ResponseDataInnerAttributes(val *GETStockLocations200ResponseDataInnerAttributes) *NullableGETStockLocations200ResponseDataInnerAttributes {
	return &NullableGETStockLocations200ResponseDataInnerAttributes{value: val, isSet: true}
}

func (v NullableGETStockLocations200ResponseDataInnerAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETStockLocations200ResponseDataInnerAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
