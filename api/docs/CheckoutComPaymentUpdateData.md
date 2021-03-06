# CheckoutComPaymentUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | [default to "checkout_com_payments"]
**Id** | **string** | The resource&#39;s id | 
**Attributes** | [**CheckoutComPaymentUpdateDataAttributes**](CheckoutComPaymentUpdateDataAttributes.md) |  | 
**Relationships** | Pointer to [**AdyenPaymentUpdateDataRelationships**](AdyenPaymentUpdateDataRelationships.md) |  | [optional] 

## Methods

### NewCheckoutComPaymentUpdateData

`func NewCheckoutComPaymentUpdateData(type_ string, id string, attributes CheckoutComPaymentUpdateDataAttributes, ) *CheckoutComPaymentUpdateData`

NewCheckoutComPaymentUpdateData instantiates a new CheckoutComPaymentUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckoutComPaymentUpdateDataWithDefaults

`func NewCheckoutComPaymentUpdateDataWithDefaults() *CheckoutComPaymentUpdateData`

NewCheckoutComPaymentUpdateDataWithDefaults instantiates a new CheckoutComPaymentUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CheckoutComPaymentUpdateData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CheckoutComPaymentUpdateData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CheckoutComPaymentUpdateData) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *CheckoutComPaymentUpdateData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CheckoutComPaymentUpdateData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CheckoutComPaymentUpdateData) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *CheckoutComPaymentUpdateData) GetAttributes() CheckoutComPaymentUpdateDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CheckoutComPaymentUpdateData) GetAttributesOk() (*CheckoutComPaymentUpdateDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CheckoutComPaymentUpdateData) SetAttributes(v CheckoutComPaymentUpdateDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *CheckoutComPaymentUpdateData) GetRelationships() AdyenPaymentUpdateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *CheckoutComPaymentUpdateData) GetRelationshipsOk() (*AdyenPaymentUpdateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *CheckoutComPaymentUpdateData) SetRelationships(v AdyenPaymentUpdateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *CheckoutComPaymentUpdateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


