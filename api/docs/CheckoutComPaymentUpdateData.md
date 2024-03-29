# CheckoutComPaymentUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Id** | **interface{}** | The resource&#39;s id | 
**Attributes** | [**PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes**](PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**AdyenPaymentUpdateDataRelationships**](AdyenPaymentUpdateDataRelationships.md) |  | [optional] 

## Methods

### NewCheckoutComPaymentUpdateData

`func NewCheckoutComPaymentUpdateData(type_ interface{}, id interface{}, attributes PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes, ) *CheckoutComPaymentUpdateData`

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

`func (o *CheckoutComPaymentUpdateData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CheckoutComPaymentUpdateData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CheckoutComPaymentUpdateData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *CheckoutComPaymentUpdateData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *CheckoutComPaymentUpdateData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetId

`func (o *CheckoutComPaymentUpdateData) GetId() interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CheckoutComPaymentUpdateData) GetIdOk() (*interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CheckoutComPaymentUpdateData) SetId(v interface{})`

SetId sets Id field to given value.


### SetIdNil

`func (o *CheckoutComPaymentUpdateData) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *CheckoutComPaymentUpdateData) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetAttributes

`func (o *CheckoutComPaymentUpdateData) GetAttributes() PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CheckoutComPaymentUpdateData) GetAttributesOk() (*PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CheckoutComPaymentUpdateData) SetAttributes(v PATCHCheckoutComPaymentsCheckoutComPaymentId200ResponseDataAttributes)`

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


