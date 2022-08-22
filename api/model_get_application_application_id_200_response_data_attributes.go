/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 2.7.3
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// GETApplicationApplicationId200ResponseDataAttributes struct for GETApplicationApplicationId200ResponseDataAttributes
type GETApplicationApplicationId200ResponseDataAttributes struct {
	// The application's internal name.
	Name *string `json:"name,omitempty"`
	// The application's kind, can be one of: 'sales_channel', 'checkout', 'contentful', 'datocms', 'sanity', 'cli', 'integration', 'webapp', 'zapier', or 'channel'
	Kind *string `json:"kind,omitempty"`
	// Indicates if the application has public access.
	PublicAccess *bool `json:"public_access,omitempty"`
	// The application's redirect URI.
	RedirectUri *string `json:"redirect_uri,omitempty"`
	// The application's scopes.
	Scopes *string `json:"scopes,omitempty"`
	// Unique identifier for the resource (hash).
	Id *string `json:"id,omitempty"`
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

// NewGETApplicationApplicationId200ResponseDataAttributes instantiates a new GETApplicationApplicationId200ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETApplicationApplicationId200ResponseDataAttributes() *GETApplicationApplicationId200ResponseDataAttributes {
	this := GETApplicationApplicationId200ResponseDataAttributes{}
	return &this
}

// NewGETApplicationApplicationId200ResponseDataAttributesWithDefaults instantiates a new GETApplicationApplicationId200ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETApplicationApplicationId200ResponseDataAttributesWithDefaults() *GETApplicationApplicationId200ResponseDataAttributes {
	this := GETApplicationApplicationId200ResponseDataAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GETApplicationApplicationId200ResponseDataAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETApplicationApplicationId200ResponseDataAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GETApplicationApplicationId200ResponseDataAttributes) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GETApplicationApplicationId200ResponseDataAttributes) SetName(v string) {
	o.Name = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *GETApplicationApplicationId200ResponseDataAttributes) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETApplicationApplicationId200ResponseDataAttributes) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *GETApplicationApplicationId200ResponseDataAttributes) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *GETApplicationApplicationId200ResponseDataAttributes) SetKind(v string) {
	o.Kind = &v
}

// GetPublicAccess returns the PublicAccess field value if set, zero value otherwise.
func (o *GETApplicationApplicationId200ResponseDataAttributes) GetPublicAccess() bool {
	if o == nil || o.PublicAccess == nil {
		var ret bool
		return ret
	}
	return *o.PublicAccess
}

// GetPublicAccessOk returns a tuple with the PublicAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETApplicationApplicationId200ResponseDataAttributes) GetPublicAccessOk() (*bool, bool) {
	if o == nil || o.PublicAccess == nil {
		return nil, false
	}
	return o.PublicAccess, true
}

// HasPublicAccess returns a boolean if a field has been set.
func (o *GETApplicationApplicationId200ResponseDataAttributes) HasPublicAccess() bool {
	if o != nil && o.PublicAccess != nil {
		return true
	}

	return false
}

// SetPublicAccess gets a reference to the given bool and assigns it to the PublicAccess field.
func (o *GETApplicationApplicationId200ResponseDataAttributes) SetPublicAccess(v bool) {
	o.PublicAccess = &v
}

// GetRedirectUri returns the RedirectUri field value if set, zero value otherwise.
func (o *GETApplicationApplicationId200ResponseDataAttributes) GetRedirectUri() string {
	if o == nil || o.RedirectUri == nil {
		var ret string
		return ret
	}
	return *o.RedirectUri
}

// GetRedirectUriOk returns a tuple with the RedirectUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETApplicationApplicationId200ResponseDataAttributes) GetRedirectUriOk() (*string, bool) {
	if o == nil || o.RedirectUri == nil {
		return nil, false
	}
	return o.RedirectUri, true
}

// HasRedirectUri returns a boolean if a field has been set.
func (o *GETApplicationApplicationId200ResponseDataAttributes) HasRedirectUri() bool {
	if o != nil && o.RedirectUri != nil {
		return true
	}

	return false
}

// SetRedirectUri gets a reference to the given string and assigns it to the RedirectUri field.
func (o *GETApplicationApplicationId200ResponseDataAttributes) SetRedirectUri(v string) {
	o.RedirectUri = &v
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *GETApplicationApplicationId200ResponseDataAttributes) GetScopes() string {
	if o == nil || o.Scopes == nil {
		var ret string
		return ret
	}
	return *o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETApplicationApplicationId200ResponseDataAttributes) GetScopesOk() (*string, bool) {
	if o == nil || o.Scopes == nil {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *GETApplicationApplicationId200ResponseDataAttributes) HasScopes() bool {
	if o != nil && o.Scopes != nil {
		return true
	}

	return false
}

// SetScopes gets a reference to the given string and assigns it to the Scopes field.
func (o *GETApplicationApplicationId200ResponseDataAttributes) SetScopes(v string) {
	o.Scopes = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GETApplicationApplicationId200ResponseDataAttributes) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETApplicationApplicationId200ResponseDataAttributes) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GETApplicationApplicationId200ResponseDataAttributes) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GETApplicationApplicationId200ResponseDataAttributes) SetId(v string) {
	o.Id = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *GETApplicationApplicationId200ResponseDataAttributes) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETApplicationApplicationId200ResponseDataAttributes) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GETApplicationApplicationId200ResponseDataAttributes) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *GETApplicationApplicationId200ResponseDataAttributes) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *GETApplicationApplicationId200ResponseDataAttributes) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETApplicationApplicationId200ResponseDataAttributes) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *GETApplicationApplicationId200ResponseDataAttributes) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *GETApplicationApplicationId200ResponseDataAttributes) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *GETApplicationApplicationId200ResponseDataAttributes) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETApplicationApplicationId200ResponseDataAttributes) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *GETApplicationApplicationId200ResponseDataAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *GETApplicationApplicationId200ResponseDataAttributes) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise.
func (o *GETApplicationApplicationId200ResponseDataAttributes) GetReferenceOrigin() string {
	if o == nil || o.ReferenceOrigin == nil {
		var ret string
		return ret
	}
	return *o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETApplicationApplicationId200ResponseDataAttributes) GetReferenceOriginOk() (*string, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *GETApplicationApplicationId200ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given string and assigns it to the ReferenceOrigin field.
func (o *GETApplicationApplicationId200ResponseDataAttributes) SetReferenceOrigin(v string) {
	o.ReferenceOrigin = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *GETApplicationApplicationId200ResponseDataAttributes) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETApplicationApplicationId200ResponseDataAttributes) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GETApplicationApplicationId200ResponseDataAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *GETApplicationApplicationId200ResponseDataAttributes) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o GETApplicationApplicationId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Kind != nil {
		toSerialize["kind"] = o.Kind
	}
	if o.PublicAccess != nil {
		toSerialize["public_access"] = o.PublicAccess
	}
	if o.RedirectUri != nil {
		toSerialize["redirect_uri"] = o.RedirectUri
	}
	if o.Scopes != nil {
		toSerialize["scopes"] = o.Scopes
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
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

type NullableGETApplicationApplicationId200ResponseDataAttributes struct {
	value *GETApplicationApplicationId200ResponseDataAttributes
	isSet bool
}

func (v NullableGETApplicationApplicationId200ResponseDataAttributes) Get() *GETApplicationApplicationId200ResponseDataAttributes {
	return v.value
}

func (v *NullableGETApplicationApplicationId200ResponseDataAttributes) Set(val *GETApplicationApplicationId200ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableGETApplicationApplicationId200ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableGETApplicationApplicationId200ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETApplicationApplicationId200ResponseDataAttributes(val *GETApplicationApplicationId200ResponseDataAttributes) *NullableGETApplicationApplicationId200ResponseDataAttributes {
	return &NullableGETApplicationApplicationId200ResponseDataAttributes{value: val, isSet: true}
}

func (v NullableGETApplicationApplicationId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETApplicationApplicationId200ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
