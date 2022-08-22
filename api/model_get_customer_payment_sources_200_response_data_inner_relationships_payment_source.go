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
	"fmt"
)

// GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource - struct for GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource
type GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource struct {
	AdyenPayment       *AdyenPayment
	BraintreePayment   *BraintreePayment
	CheckoutComPayment *CheckoutComPayment
	ExternalPayment    *ExternalPayment
	KlarnaPayment      *KlarnaPayment
	PaypalPayment      *PaypalPayment
	StripePayment      *StripePayment
	WireTransfer       *WireTransfer
}

// AdyenPaymentAsGETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource is a convenience function that returns AdyenPayment wrapped in GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource
func AdyenPaymentAsGETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource(v *AdyenPayment) GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource {
	return GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource{
		AdyenPayment: v,
	}
}

// BraintreePaymentAsGETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource is a convenience function that returns BraintreePayment wrapped in GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource
func BraintreePaymentAsGETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource(v *BraintreePayment) GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource {
	return GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource{
		BraintreePayment: v,
	}
}

// CheckoutComPaymentAsGETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource is a convenience function that returns CheckoutComPayment wrapped in GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource
func CheckoutComPaymentAsGETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource(v *CheckoutComPayment) GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource {
	return GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource{
		CheckoutComPayment: v,
	}
}

// ExternalPaymentAsGETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource is a convenience function that returns ExternalPayment wrapped in GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource
func ExternalPaymentAsGETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource(v *ExternalPayment) GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource {
	return GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource{
		ExternalPayment: v,
	}
}

// KlarnaPaymentAsGETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource is a convenience function that returns KlarnaPayment wrapped in GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource
func KlarnaPaymentAsGETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource(v *KlarnaPayment) GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource {
	return GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource{
		KlarnaPayment: v,
	}
}

// PaypalPaymentAsGETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource is a convenience function that returns PaypalPayment wrapped in GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource
func PaypalPaymentAsGETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource(v *PaypalPayment) GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource {
	return GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource{
		PaypalPayment: v,
	}
}

// StripePaymentAsGETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource is a convenience function that returns StripePayment wrapped in GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource
func StripePaymentAsGETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource(v *StripePayment) GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource {
	return GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource{
		StripePayment: v,
	}
}

// WireTransferAsGETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource is a convenience function that returns WireTransfer wrapped in GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource
func WireTransferAsGETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource(v *WireTransfer) GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource {
	return GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource{
		WireTransfer: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AdyenPayment
	err = newStrictDecoder(data).Decode(&dst.AdyenPayment)
	if err == nil {
		jsonAdyenPayment, _ := json.Marshal(dst.AdyenPayment)
		if string(jsonAdyenPayment) == "{}" { // empty struct
			dst.AdyenPayment = nil
		} else {
			match++
		}
	} else {
		dst.AdyenPayment = nil
	}

	// try to unmarshal data into BraintreePayment
	err = newStrictDecoder(data).Decode(&dst.BraintreePayment)
	if err == nil {
		jsonBraintreePayment, _ := json.Marshal(dst.BraintreePayment)
		if string(jsonBraintreePayment) == "{}" { // empty struct
			dst.BraintreePayment = nil
		} else {
			match++
		}
	} else {
		dst.BraintreePayment = nil
	}

	// try to unmarshal data into CheckoutComPayment
	err = newStrictDecoder(data).Decode(&dst.CheckoutComPayment)
	if err == nil {
		jsonCheckoutComPayment, _ := json.Marshal(dst.CheckoutComPayment)
		if string(jsonCheckoutComPayment) == "{}" { // empty struct
			dst.CheckoutComPayment = nil
		} else {
			match++
		}
	} else {
		dst.CheckoutComPayment = nil
	}

	// try to unmarshal data into ExternalPayment
	err = newStrictDecoder(data).Decode(&dst.ExternalPayment)
	if err == nil {
		jsonExternalPayment, _ := json.Marshal(dst.ExternalPayment)
		if string(jsonExternalPayment) == "{}" { // empty struct
			dst.ExternalPayment = nil
		} else {
			match++
		}
	} else {
		dst.ExternalPayment = nil
	}

	// try to unmarshal data into KlarnaPayment
	err = newStrictDecoder(data).Decode(&dst.KlarnaPayment)
	if err == nil {
		jsonKlarnaPayment, _ := json.Marshal(dst.KlarnaPayment)
		if string(jsonKlarnaPayment) == "{}" { // empty struct
			dst.KlarnaPayment = nil
		} else {
			match++
		}
	} else {
		dst.KlarnaPayment = nil
	}

	// try to unmarshal data into PaypalPayment
	err = newStrictDecoder(data).Decode(&dst.PaypalPayment)
	if err == nil {
		jsonPaypalPayment, _ := json.Marshal(dst.PaypalPayment)
		if string(jsonPaypalPayment) == "{}" { // empty struct
			dst.PaypalPayment = nil
		} else {
			match++
		}
	} else {
		dst.PaypalPayment = nil
	}

	// try to unmarshal data into StripePayment
	err = newStrictDecoder(data).Decode(&dst.StripePayment)
	if err == nil {
		jsonStripePayment, _ := json.Marshal(dst.StripePayment)
		if string(jsonStripePayment) == "{}" { // empty struct
			dst.StripePayment = nil
		} else {
			match++
		}
	} else {
		dst.StripePayment = nil
	}

	// try to unmarshal data into WireTransfer
	err = newStrictDecoder(data).Decode(&dst.WireTransfer)
	if err == nil {
		jsonWireTransfer, _ := json.Marshal(dst.WireTransfer)
		if string(jsonWireTransfer) == "{}" { // empty struct
			dst.WireTransfer = nil
		} else {
			match++
		}
	} else {
		dst.WireTransfer = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AdyenPayment = nil
		dst.BraintreePayment = nil
		dst.CheckoutComPayment = nil
		dst.ExternalPayment = nil
		dst.KlarnaPayment = nil
		dst.PaypalPayment = nil
		dst.StripePayment = nil
		dst.WireTransfer = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource) MarshalJSON() ([]byte, error) {
	if src.AdyenPayment != nil {
		return json.Marshal(&src.AdyenPayment)
	}

	if src.BraintreePayment != nil {
		return json.Marshal(&src.BraintreePayment)
	}

	if src.CheckoutComPayment != nil {
		return json.Marshal(&src.CheckoutComPayment)
	}

	if src.ExternalPayment != nil {
		return json.Marshal(&src.ExternalPayment)
	}

	if src.KlarnaPayment != nil {
		return json.Marshal(&src.KlarnaPayment)
	}

	if src.PaypalPayment != nil {
		return json.Marshal(&src.PaypalPayment)
	}

	if src.StripePayment != nil {
		return json.Marshal(&src.StripePayment)
	}

	if src.WireTransfer != nil {
		return json.Marshal(&src.WireTransfer)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.AdyenPayment != nil {
		return obj.AdyenPayment
	}

	if obj.BraintreePayment != nil {
		return obj.BraintreePayment
	}

	if obj.CheckoutComPayment != nil {
		return obj.CheckoutComPayment
	}

	if obj.ExternalPayment != nil {
		return obj.ExternalPayment
	}

	if obj.KlarnaPayment != nil {
		return obj.KlarnaPayment
	}

	if obj.PaypalPayment != nil {
		return obj.PaypalPayment
	}

	if obj.StripePayment != nil {
		return obj.StripePayment
	}

	if obj.WireTransfer != nil {
		return obj.WireTransfer
	}

	// all schemas are nil
	return nil
}

type NullableGETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource struct {
	value *GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource
	isSet bool
}

func (v NullableGETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource) Get() *GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource {
	return v.value
}

func (v *NullableGETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource) Set(val *GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource) {
	v.value = val
	v.isSet = true
}

func (v NullableGETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource) IsSet() bool {
	return v.isSet
}

func (v *NullableGETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource(val *GETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource) *NullableGETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource {
	return &NullableGETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource{value: val, isSet: true}
}

func (v NullableGETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETCustomerPaymentSources200ResponseDataInnerRelationshipsPaymentSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}