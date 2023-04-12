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

// checks if the POSTSubscriptionModels201ResponseDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTSubscriptionModels201ResponseDataAttributes{}

// POSTSubscriptionModels201ResponseDataAttributes struct for POSTSubscriptionModels201ResponseDataAttributes
type POSTSubscriptionModels201ResponseDataAttributes struct {
	// The subscription model's internal name.
	Name interface{} `json:"name"`
	// The subscription model's strategy used to generate order subscriptions: one between 'by_frequency' (default) and 'by_line_items'.
	Strategy interface{} `json:"strategy,omitempty"`
	// An object that contains the frequencies available for this subscription model. Supported ones are 'hourly', 'daily', 'weekly', 'monthly', 'two-month', 'three-month', 'four-month', 'six-month', 'yearly', or your custom crontab expression (min unit is hour).
	Frequencies interface{} `json:"frequencies"`
	// Indicates if the created subscriptions will be activated considering the placed source order as its first run, default to true.
	AutoActivate interface{} `json:"auto_activate,omitempty"`
	// Indicates if the created subscriptions will be cancelled in case the source order is cancelled, default to false.
	AutoCancel interface{} `json:"auto_cancel,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference interface{} `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin interface{} `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata interface{} `json:"metadata,omitempty"`
}

// NewPOSTSubscriptionModels201ResponseDataAttributes instantiates a new POSTSubscriptionModels201ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTSubscriptionModels201ResponseDataAttributes(name interface{}, frequencies interface{}) *POSTSubscriptionModels201ResponseDataAttributes {
	this := POSTSubscriptionModels201ResponseDataAttributes{}
	this.Name = name
	this.Frequencies = frequencies
	return &this
}

// NewPOSTSubscriptionModels201ResponseDataAttributesWithDefaults instantiates a new POSTSubscriptionModels201ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTSubscriptionModels201ResponseDataAttributesWithDefaults() *POSTSubscriptionModels201ResponseDataAttributes {
	this := POSTSubscriptionModels201ResponseDataAttributes{}
	return &this
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *POSTSubscriptionModels201ResponseDataAttributes) GetName() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTSubscriptionModels201ResponseDataAttributes) GetNameOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *POSTSubscriptionModels201ResponseDataAttributes) SetName(v interface{}) {
	o.Name = v
}

// GetStrategy returns the Strategy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTSubscriptionModels201ResponseDataAttributes) GetStrategy() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Strategy
}

// GetStrategyOk returns a tuple with the Strategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTSubscriptionModels201ResponseDataAttributes) GetStrategyOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Strategy) {
		return nil, false
	}
	return &o.Strategy, true
}

// HasStrategy returns a boolean if a field has been set.
func (o *POSTSubscriptionModels201ResponseDataAttributes) HasStrategy() bool {
	if o != nil && IsNil(o.Strategy) {
		return true
	}

	return false
}

// SetStrategy gets a reference to the given interface{} and assigns it to the Strategy field.
func (o *POSTSubscriptionModels201ResponseDataAttributes) SetStrategy(v interface{}) {
	o.Strategy = v
}

// GetFrequencies returns the Frequencies field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *POSTSubscriptionModels201ResponseDataAttributes) GetFrequencies() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Frequencies
}

// GetFrequenciesOk returns a tuple with the Frequencies field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTSubscriptionModels201ResponseDataAttributes) GetFrequenciesOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Frequencies) {
		return nil, false
	}
	return &o.Frequencies, true
}

// SetFrequencies sets field value
func (o *POSTSubscriptionModels201ResponseDataAttributes) SetFrequencies(v interface{}) {
	o.Frequencies = v
}

// GetAutoActivate returns the AutoActivate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTSubscriptionModels201ResponseDataAttributes) GetAutoActivate() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.AutoActivate
}

// GetAutoActivateOk returns a tuple with the AutoActivate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTSubscriptionModels201ResponseDataAttributes) GetAutoActivateOk() (*interface{}, bool) {
	if o == nil || IsNil(o.AutoActivate) {
		return nil, false
	}
	return &o.AutoActivate, true
}

// HasAutoActivate returns a boolean if a field has been set.
func (o *POSTSubscriptionModels201ResponseDataAttributes) HasAutoActivate() bool {
	if o != nil && IsNil(o.AutoActivate) {
		return true
	}

	return false
}

// SetAutoActivate gets a reference to the given interface{} and assigns it to the AutoActivate field.
func (o *POSTSubscriptionModels201ResponseDataAttributes) SetAutoActivate(v interface{}) {
	o.AutoActivate = v
}

// GetAutoCancel returns the AutoCancel field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTSubscriptionModels201ResponseDataAttributes) GetAutoCancel() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.AutoCancel
}

// GetAutoCancelOk returns a tuple with the AutoCancel field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTSubscriptionModels201ResponseDataAttributes) GetAutoCancelOk() (*interface{}, bool) {
	if o == nil || IsNil(o.AutoCancel) {
		return nil, false
	}
	return &o.AutoCancel, true
}

// HasAutoCancel returns a boolean if a field has been set.
func (o *POSTSubscriptionModels201ResponseDataAttributes) HasAutoCancel() bool {
	if o != nil && IsNil(o.AutoCancel) {
		return true
	}

	return false
}

// SetAutoCancel gets a reference to the given interface{} and assigns it to the AutoCancel field.
func (o *POSTSubscriptionModels201ResponseDataAttributes) SetAutoCancel(v interface{}) {
	o.AutoCancel = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTSubscriptionModels201ResponseDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTSubscriptionModels201ResponseDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *POSTSubscriptionModels201ResponseDataAttributes) HasReference() bool {
	if o != nil && IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *POSTSubscriptionModels201ResponseDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTSubscriptionModels201ResponseDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTSubscriptionModels201ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReferenceOrigin) {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *POSTSubscriptionModels201ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && IsNil(o.ReferenceOrigin) {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *POSTSubscriptionModels201ResponseDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTSubscriptionModels201ResponseDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTSubscriptionModels201ResponseDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *POSTSubscriptionModels201ResponseDataAttributes) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *POSTSubscriptionModels201ResponseDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

func (o POSTSubscriptionModels201ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTSubscriptionModels201ResponseDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Strategy != nil {
		toSerialize["strategy"] = o.Strategy
	}
	if o.Frequencies != nil {
		toSerialize["frequencies"] = o.Frequencies
	}
	if o.AutoActivate != nil {
		toSerialize["auto_activate"] = o.AutoActivate
	}
	if o.AutoCancel != nil {
		toSerialize["auto_cancel"] = o.AutoCancel
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

type NullablePOSTSubscriptionModels201ResponseDataAttributes struct {
	value *POSTSubscriptionModels201ResponseDataAttributes
	isSet bool
}

func (v NullablePOSTSubscriptionModels201ResponseDataAttributes) Get() *POSTSubscriptionModels201ResponseDataAttributes {
	return v.value
}

func (v *NullablePOSTSubscriptionModels201ResponseDataAttributes) Set(val *POSTSubscriptionModels201ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTSubscriptionModels201ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTSubscriptionModels201ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTSubscriptionModels201ResponseDataAttributes(val *POSTSubscriptionModels201ResponseDataAttributes) *NullablePOSTSubscriptionModels201ResponseDataAttributes {
	return &NullablePOSTSubscriptionModels201ResponseDataAttributes{value: val, isSet: true}
}

func (v NullablePOSTSubscriptionModels201ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTSubscriptionModels201ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}