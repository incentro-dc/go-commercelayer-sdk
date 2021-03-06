# OrderCreateDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Market** | Pointer to [**AvalaraAccountDataRelationshipsMarkets**](AvalaraAccountDataRelationshipsMarkets.md) |  | [optional] 
**Customer** | Pointer to [**CouponRecipientDataRelationshipsCustomer**](CouponRecipientDataRelationshipsCustomer.md) |  | [optional] 
**ShippingAddress** | Pointer to [**BingGeocoderDataRelationshipsAddresses**](BingGeocoderDataRelationshipsAddresses.md) |  | [optional] 
**BillingAddress** | Pointer to [**BingGeocoderDataRelationshipsAddresses**](BingGeocoderDataRelationshipsAddresses.md) |  | [optional] 
**PaymentMethod** | Pointer to [**AdyenGatewayDataRelationshipsPaymentMethods**](AdyenGatewayDataRelationshipsPaymentMethods.md) |  | [optional] 
**PaymentSource** | Pointer to [**CustomerPaymentSourceDataRelationshipsPaymentSource**](CustomerPaymentSourceDataRelationshipsPaymentSource.md) |  | [optional] 

## Methods

### NewOrderCreateDataRelationships

`func NewOrderCreateDataRelationships() *OrderCreateDataRelationships`

NewOrderCreateDataRelationships instantiates a new OrderCreateDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderCreateDataRelationshipsWithDefaults

`func NewOrderCreateDataRelationshipsWithDefaults() *OrderCreateDataRelationships`

NewOrderCreateDataRelationshipsWithDefaults instantiates a new OrderCreateDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMarket

`func (o *OrderCreateDataRelationships) GetMarket() AvalaraAccountDataRelationshipsMarkets`

GetMarket returns the Market field if non-nil, zero value otherwise.

### GetMarketOk

`func (o *OrderCreateDataRelationships) GetMarketOk() (*AvalaraAccountDataRelationshipsMarkets, bool)`

GetMarketOk returns a tuple with the Market field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarket

`func (o *OrderCreateDataRelationships) SetMarket(v AvalaraAccountDataRelationshipsMarkets)`

SetMarket sets Market field to given value.

### HasMarket

`func (o *OrderCreateDataRelationships) HasMarket() bool`

HasMarket returns a boolean if a field has been set.

### GetCustomer

`func (o *OrderCreateDataRelationships) GetCustomer() CouponRecipientDataRelationshipsCustomer`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *OrderCreateDataRelationships) GetCustomerOk() (*CouponRecipientDataRelationshipsCustomer, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *OrderCreateDataRelationships) SetCustomer(v CouponRecipientDataRelationshipsCustomer)`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *OrderCreateDataRelationships) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### GetShippingAddress

`func (o *OrderCreateDataRelationships) GetShippingAddress() BingGeocoderDataRelationshipsAddresses`

GetShippingAddress returns the ShippingAddress field if non-nil, zero value otherwise.

### GetShippingAddressOk

`func (o *OrderCreateDataRelationships) GetShippingAddressOk() (*BingGeocoderDataRelationshipsAddresses, bool)`

GetShippingAddressOk returns a tuple with the ShippingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingAddress

`func (o *OrderCreateDataRelationships) SetShippingAddress(v BingGeocoderDataRelationshipsAddresses)`

SetShippingAddress sets ShippingAddress field to given value.

### HasShippingAddress

`func (o *OrderCreateDataRelationships) HasShippingAddress() bool`

HasShippingAddress returns a boolean if a field has been set.

### GetBillingAddress

`func (o *OrderCreateDataRelationships) GetBillingAddress() BingGeocoderDataRelationshipsAddresses`

GetBillingAddress returns the BillingAddress field if non-nil, zero value otherwise.

### GetBillingAddressOk

`func (o *OrderCreateDataRelationships) GetBillingAddressOk() (*BingGeocoderDataRelationshipsAddresses, bool)`

GetBillingAddressOk returns a tuple with the BillingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddress

`func (o *OrderCreateDataRelationships) SetBillingAddress(v BingGeocoderDataRelationshipsAddresses)`

SetBillingAddress sets BillingAddress field to given value.

### HasBillingAddress

`func (o *OrderCreateDataRelationships) HasBillingAddress() bool`

HasBillingAddress returns a boolean if a field has been set.

### GetPaymentMethod

`func (o *OrderCreateDataRelationships) GetPaymentMethod() AdyenGatewayDataRelationshipsPaymentMethods`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *OrderCreateDataRelationships) GetPaymentMethodOk() (*AdyenGatewayDataRelationshipsPaymentMethods, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *OrderCreateDataRelationships) SetPaymentMethod(v AdyenGatewayDataRelationshipsPaymentMethods)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *OrderCreateDataRelationships) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.

### GetPaymentSource

`func (o *OrderCreateDataRelationships) GetPaymentSource() CustomerPaymentSourceDataRelationshipsPaymentSource`

GetPaymentSource returns the PaymentSource field if non-nil, zero value otherwise.

### GetPaymentSourceOk

`func (o *OrderCreateDataRelationships) GetPaymentSourceOk() (*CustomerPaymentSourceDataRelationshipsPaymentSource, bool)`

GetPaymentSourceOk returns a tuple with the PaymentSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentSource

`func (o *OrderCreateDataRelationships) SetPaymentSource(v CustomerPaymentSourceDataRelationshipsPaymentSource)`

SetPaymentSource sets PaymentSource field to given value.

### HasPaymentSource

`func (o *OrderCreateDataRelationships) HasPaymentSource() bool`

HasPaymentSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


