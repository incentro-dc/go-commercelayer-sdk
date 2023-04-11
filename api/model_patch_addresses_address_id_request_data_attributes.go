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

// checks if the PATCHAddressesAddressIdRequestDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PATCHAddressesAddressIdRequestDataAttributes{}

// PATCHAddressesAddressIdRequestDataAttributes struct for PATCHAddressesAddressIdRequestDataAttributes
type PATCHAddressesAddressIdRequestDataAttributes struct {
	// Indicates if it's a business or a personal address
	Business interface{} `json:"business,omitempty"`
	// Address first name (personal)
	FirstName interface{} `json:"first_name,omitempty"`
	// Address last name (personal)
	LastName interface{} `json:"last_name,omitempty"`
	// Address company name (business)
	Company interface{} `json:"company,omitempty"`
	// Address line 1, i.e. Street address, PO Box
	Line1 interface{} `json:"line_1,omitempty"`
	// Address line 2, i.e. Apartment, Suite, Building
	Line2 interface{} `json:"line_2,omitempty"`
	// Address city
	City interface{} `json:"city,omitempty"`
	// ZIP or postal code
	ZipCode interface{} `json:"zip_code,omitempty"`
	// State, province or region code.
	StateCode interface{} `json:"state_code,omitempty"`
	// The international 2-letter country code as defined by the ISO 3166-1 standard
	CountryCode interface{} `json:"country_code,omitempty"`
	// Phone number (including extension).
	Phone interface{} `json:"phone,omitempty"`
	// Email address.
	Email interface{} `json:"email,omitempty"`
	// A free notes attached to the address. When used as a shipping address, this can be useful to let the customers add specific delivery instructions.
	Notes interface{} `json:"notes,omitempty"`
	// The address geocoded latitude. This is automatically generated when creating a shipping/billing address for an order and a valid geocoder is attached to the order's market.
	Lat interface{} `json:"lat,omitempty"`
	// The address geocoded longitude. This is automatically generated when creating a shipping/billing address for an order and a valid geocoder is attached to the order's market.
	Lng interface{} `json:"lng,omitempty"`
	// Customer's billing information (i.e. VAT number, codice fiscale)
	BillingInfo interface{} `json:"billing_info,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference interface{} `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin interface{} `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata interface{} `json:"metadata,omitempty"`
}

// NewPATCHAddressesAddressIdRequestDataAttributes instantiates a new PATCHAddressesAddressIdRequestDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHAddressesAddressIdRequestDataAttributes() *PATCHAddressesAddressIdRequestDataAttributes {
	this := PATCHAddressesAddressIdRequestDataAttributes{}
	return &this
}

// NewPATCHAddressesAddressIdRequestDataAttributesWithDefaults instantiates a new PATCHAddressesAddressIdRequestDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHAddressesAddressIdRequestDataAttributesWithDefaults() *PATCHAddressesAddressIdRequestDataAttributes {
	this := PATCHAddressesAddressIdRequestDataAttributes{}
	return &this
}

// GetBusiness returns the Business field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHAddressesAddressIdRequestDataAttributes) GetBusiness() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Business
}

// GetBusinessOk returns a tuple with the Business field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHAddressesAddressIdRequestDataAttributes) GetBusinessOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Business) {
		return nil, false
	}
	return &o.Business, true
}

// HasBusiness returns a boolean if a field has been set.
func (o *PATCHAddressesAddressIdRequestDataAttributes) HasBusiness() bool {
	if o != nil && IsNil(o.Business) {
		return true
	}

	return false
}

// SetBusiness gets a reference to the given interface{} and assigns it to the Business field.
func (o *PATCHAddressesAddressIdRequestDataAttributes) SetBusiness(v interface{}) {
	o.Business = v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHAddressesAddressIdRequestDataAttributes) GetFirstName() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHAddressesAddressIdRequestDataAttributes) GetFirstNameOk() (*interface{}, bool) {
	if o == nil || IsNil(o.FirstName) {
		return nil, false
	}
	return &o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *PATCHAddressesAddressIdRequestDataAttributes) HasFirstName() bool {
	if o != nil && IsNil(o.FirstName) {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given interface{} and assigns it to the FirstName field.
func (o *PATCHAddressesAddressIdRequestDataAttributes) SetFirstName(v interface{}) {
	o.FirstName = v
}

// GetLastName returns the LastName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHAddressesAddressIdRequestDataAttributes) GetLastName() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHAddressesAddressIdRequestDataAttributes) GetLastNameOk() (*interface{}, bool) {
	if o == nil || IsNil(o.LastName) {
		return nil, false
	}
	return &o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *PATCHAddressesAddressIdRequestDataAttributes) HasLastName() bool {
	if o != nil && IsNil(o.LastName) {
		return true
	}

	return false
}

// SetLastName gets a reference to the given interface{} and assigns it to the LastName field.
func (o *PATCHAddressesAddressIdRequestDataAttributes) SetLastName(v interface{}) {
	o.LastName = v
}

// GetCompany returns the Company field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHAddressesAddressIdRequestDataAttributes) GetCompany() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Company
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHAddressesAddressIdRequestDataAttributes) GetCompanyOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Company) {
		return nil, false
	}
	return &o.Company, true
}

// HasCompany returns a boolean if a field has been set.
func (o *PATCHAddressesAddressIdRequestDataAttributes) HasCompany() bool {
	if o != nil && IsNil(o.Company) {
		return true
	}

	return false
}

// SetCompany gets a reference to the given interface{} and assigns it to the Company field.
func (o *PATCHAddressesAddressIdRequestDataAttributes) SetCompany(v interface{}) {
	o.Company = v
}

// GetLine1 returns the Line1 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHAddressesAddressIdRequestDataAttributes) GetLine1() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Line1
}

// GetLine1Ok returns a tuple with the Line1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHAddressesAddressIdRequestDataAttributes) GetLine1Ok() (*interface{}, bool) {
	if o == nil || IsNil(o.Line1) {
		return nil, false
	}
	return &o.Line1, true
}

// HasLine1 returns a boolean if a field has been set.
func (o *PATCHAddressesAddressIdRequestDataAttributes) HasLine1() bool {
	if o != nil && IsNil(o.Line1) {
		return true
	}

	return false
}

// SetLine1 gets a reference to the given interface{} and assigns it to the Line1 field.
func (o *PATCHAddressesAddressIdRequestDataAttributes) SetLine1(v interface{}) {
	o.Line1 = v
}

// GetLine2 returns the Line2 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHAddressesAddressIdRequestDataAttributes) GetLine2() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Line2
}

// GetLine2Ok returns a tuple with the Line2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHAddressesAddressIdRequestDataAttributes) GetLine2Ok() (*interface{}, bool) {
	if o == nil || IsNil(o.Line2) {
		return nil, false
	}
	return &o.Line2, true
}

// HasLine2 returns a boolean if a field has been set.
func (o *PATCHAddressesAddressIdRequestDataAttributes) HasLine2() bool {
	if o != nil && IsNil(o.Line2) {
		return true
	}

	return false
}

// SetLine2 gets a reference to the given interface{} and assigns it to the Line2 field.
func (o *PATCHAddressesAddressIdRequestDataAttributes) SetLine2(v interface{}) {
	o.Line2 = v
}

// GetCity returns the City field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHAddressesAddressIdRequestDataAttributes) GetCity() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHAddressesAddressIdRequestDataAttributes) GetCityOk() (*interface{}, bool) {
	if o == nil || IsNil(o.City) {
		return nil, false
	}
	return &o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *PATCHAddressesAddressIdRequestDataAttributes) HasCity() bool {
	if o != nil && IsNil(o.City) {
		return true
	}

	return false
}

// SetCity gets a reference to the given interface{} and assigns it to the City field.
func (o *PATCHAddressesAddressIdRequestDataAttributes) SetCity(v interface{}) {
	o.City = v
}

// GetZipCode returns the ZipCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHAddressesAddressIdRequestDataAttributes) GetZipCode() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ZipCode
}

// GetZipCodeOk returns a tuple with the ZipCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHAddressesAddressIdRequestDataAttributes) GetZipCodeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ZipCode) {
		return nil, false
	}
	return &o.ZipCode, true
}

// HasZipCode returns a boolean if a field has been set.
func (o *PATCHAddressesAddressIdRequestDataAttributes) HasZipCode() bool {
	if o != nil && IsNil(o.ZipCode) {
		return true
	}

	return false
}

// SetZipCode gets a reference to the given interface{} and assigns it to the ZipCode field.
func (o *PATCHAddressesAddressIdRequestDataAttributes) SetZipCode(v interface{}) {
	o.ZipCode = v
}

// GetStateCode returns the StateCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHAddressesAddressIdRequestDataAttributes) GetStateCode() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.StateCode
}

// GetStateCodeOk returns a tuple with the StateCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHAddressesAddressIdRequestDataAttributes) GetStateCodeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.StateCode) {
		return nil, false
	}
	return &o.StateCode, true
}

// HasStateCode returns a boolean if a field has been set.
func (o *PATCHAddressesAddressIdRequestDataAttributes) HasStateCode() bool {
	if o != nil && IsNil(o.StateCode) {
		return true
	}

	return false
}

// SetStateCode gets a reference to the given interface{} and assigns it to the StateCode field.
func (o *PATCHAddressesAddressIdRequestDataAttributes) SetStateCode(v interface{}) {
	o.StateCode = v
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHAddressesAddressIdRequestDataAttributes) GetCountryCode() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHAddressesAddressIdRequestDataAttributes) GetCountryCodeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.CountryCode) {
		return nil, false
	}
	return &o.CountryCode, true
}

// HasCountryCode returns a boolean if a field has been set.
func (o *PATCHAddressesAddressIdRequestDataAttributes) HasCountryCode() bool {
	if o != nil && IsNil(o.CountryCode) {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given interface{} and assigns it to the CountryCode field.
func (o *PATCHAddressesAddressIdRequestDataAttributes) SetCountryCode(v interface{}) {
	o.CountryCode = v
}

// GetPhone returns the Phone field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHAddressesAddressIdRequestDataAttributes) GetPhone() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHAddressesAddressIdRequestDataAttributes) GetPhoneOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Phone) {
		return nil, false
	}
	return &o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *PATCHAddressesAddressIdRequestDataAttributes) HasPhone() bool {
	if o != nil && IsNil(o.Phone) {
		return true
	}

	return false
}

// SetPhone gets a reference to the given interface{} and assigns it to the Phone field.
func (o *PATCHAddressesAddressIdRequestDataAttributes) SetPhone(v interface{}) {
	o.Phone = v
}

// GetEmail returns the Email field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHAddressesAddressIdRequestDataAttributes) GetEmail() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHAddressesAddressIdRequestDataAttributes) GetEmailOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return &o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *PATCHAddressesAddressIdRequestDataAttributes) HasEmail() bool {
	if o != nil && IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given interface{} and assigns it to the Email field.
func (o *PATCHAddressesAddressIdRequestDataAttributes) SetEmail(v interface{}) {
	o.Email = v
}

// GetNotes returns the Notes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHAddressesAddressIdRequestDataAttributes) GetNotes() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHAddressesAddressIdRequestDataAttributes) GetNotesOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Notes) {
		return nil, false
	}
	return &o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *PATCHAddressesAddressIdRequestDataAttributes) HasNotes() bool {
	if o != nil && IsNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given interface{} and assigns it to the Notes field.
func (o *PATCHAddressesAddressIdRequestDataAttributes) SetNotes(v interface{}) {
	o.Notes = v
}

// GetLat returns the Lat field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHAddressesAddressIdRequestDataAttributes) GetLat() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Lat
}

// GetLatOk returns a tuple with the Lat field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHAddressesAddressIdRequestDataAttributes) GetLatOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Lat) {
		return nil, false
	}
	return &o.Lat, true
}

// HasLat returns a boolean if a field has been set.
func (o *PATCHAddressesAddressIdRequestDataAttributes) HasLat() bool {
	if o != nil && IsNil(o.Lat) {
		return true
	}

	return false
}

// SetLat gets a reference to the given interface{} and assigns it to the Lat field.
func (o *PATCHAddressesAddressIdRequestDataAttributes) SetLat(v interface{}) {
	o.Lat = v
}

// GetLng returns the Lng field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHAddressesAddressIdRequestDataAttributes) GetLng() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Lng
}

// GetLngOk returns a tuple with the Lng field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHAddressesAddressIdRequestDataAttributes) GetLngOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Lng) {
		return nil, false
	}
	return &o.Lng, true
}

// HasLng returns a boolean if a field has been set.
func (o *PATCHAddressesAddressIdRequestDataAttributes) HasLng() bool {
	if o != nil && IsNil(o.Lng) {
		return true
	}

	return false
}

// SetLng gets a reference to the given interface{} and assigns it to the Lng field.
func (o *PATCHAddressesAddressIdRequestDataAttributes) SetLng(v interface{}) {
	o.Lng = v
}

// GetBillingInfo returns the BillingInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHAddressesAddressIdRequestDataAttributes) GetBillingInfo() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.BillingInfo
}

// GetBillingInfoOk returns a tuple with the BillingInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHAddressesAddressIdRequestDataAttributes) GetBillingInfoOk() (*interface{}, bool) {
	if o == nil || IsNil(o.BillingInfo) {
		return nil, false
	}
	return &o.BillingInfo, true
}

// HasBillingInfo returns a boolean if a field has been set.
func (o *PATCHAddressesAddressIdRequestDataAttributes) HasBillingInfo() bool {
	if o != nil && IsNil(o.BillingInfo) {
		return true
	}

	return false
}

// SetBillingInfo gets a reference to the given interface{} and assigns it to the BillingInfo field.
func (o *PATCHAddressesAddressIdRequestDataAttributes) SetBillingInfo(v interface{}) {
	o.BillingInfo = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHAddressesAddressIdRequestDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHAddressesAddressIdRequestDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *PATCHAddressesAddressIdRequestDataAttributes) HasReference() bool {
	if o != nil && IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *PATCHAddressesAddressIdRequestDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHAddressesAddressIdRequestDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHAddressesAddressIdRequestDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReferenceOrigin) {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *PATCHAddressesAddressIdRequestDataAttributes) HasReferenceOrigin() bool {
	if o != nil && IsNil(o.ReferenceOrigin) {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *PATCHAddressesAddressIdRequestDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHAddressesAddressIdRequestDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHAddressesAddressIdRequestDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *PATCHAddressesAddressIdRequestDataAttributes) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *PATCHAddressesAddressIdRequestDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

func (o PATCHAddressesAddressIdRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PATCHAddressesAddressIdRequestDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Business != nil {
		toSerialize["business"] = o.Business
	}
	if o.FirstName != nil {
		toSerialize["first_name"] = o.FirstName
	}
	if o.LastName != nil {
		toSerialize["last_name"] = o.LastName
	}
	if o.Company != nil {
		toSerialize["company"] = o.Company
	}
	if o.Line1 != nil {
		toSerialize["line_1"] = o.Line1
	}
	if o.Line2 != nil {
		toSerialize["line_2"] = o.Line2
	}
	if o.City != nil {
		toSerialize["city"] = o.City
	}
	if o.ZipCode != nil {
		toSerialize["zip_code"] = o.ZipCode
	}
	if o.StateCode != nil {
		toSerialize["state_code"] = o.StateCode
	}
	if o.CountryCode != nil {
		toSerialize["country_code"] = o.CountryCode
	}
	if o.Phone != nil {
		toSerialize["phone"] = o.Phone
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.Notes != nil {
		toSerialize["notes"] = o.Notes
	}
	if o.Lat != nil {
		toSerialize["lat"] = o.Lat
	}
	if o.Lng != nil {
		toSerialize["lng"] = o.Lng
	}
	if o.BillingInfo != nil {
		toSerialize["billing_info"] = o.BillingInfo
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

type NullablePATCHAddressesAddressIdRequestDataAttributes struct {
	value *PATCHAddressesAddressIdRequestDataAttributes
	isSet bool
}

func (v NullablePATCHAddressesAddressIdRequestDataAttributes) Get() *PATCHAddressesAddressIdRequestDataAttributes {
	return v.value
}

func (v *NullablePATCHAddressesAddressIdRequestDataAttributes) Set(val *PATCHAddressesAddressIdRequestDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHAddressesAddressIdRequestDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHAddressesAddressIdRequestDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHAddressesAddressIdRequestDataAttributes(val *PATCHAddressesAddressIdRequestDataAttributes) *NullablePATCHAddressesAddressIdRequestDataAttributes {
	return &NullablePATCHAddressesAddressIdRequestDataAttributes{value: val, isSet: true}
}

func (v NullablePATCHAddressesAddressIdRequestDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHAddressesAddressIdRequestDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}