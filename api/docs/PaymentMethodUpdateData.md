# PaymentMethodUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Id** | **interface{}** | The resource&#39;s id | 
**Attributes** | [**PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes**](PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**PaymentMethodUpdateDataRelationships**](PaymentMethodUpdateDataRelationships.md) |  | [optional] 

## Methods

### NewPaymentMethodUpdateData

`func NewPaymentMethodUpdateData(type_ interface{}, id interface{}, attributes PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes, ) *PaymentMethodUpdateData`

NewPaymentMethodUpdateData instantiates a new PaymentMethodUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentMethodUpdateDataWithDefaults

`func NewPaymentMethodUpdateDataWithDefaults() *PaymentMethodUpdateData`

NewPaymentMethodUpdateDataWithDefaults instantiates a new PaymentMethodUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PaymentMethodUpdateData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PaymentMethodUpdateData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PaymentMethodUpdateData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *PaymentMethodUpdateData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *PaymentMethodUpdateData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetId

`func (o *PaymentMethodUpdateData) GetId() interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaymentMethodUpdateData) GetIdOk() (*interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaymentMethodUpdateData) SetId(v interface{})`

SetId sets Id field to given value.


### SetIdNil

`func (o *PaymentMethodUpdateData) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *PaymentMethodUpdateData) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetAttributes

`func (o *PaymentMethodUpdateData) GetAttributes() PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PaymentMethodUpdateData) GetAttributesOk() (*PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *PaymentMethodUpdateData) SetAttributes(v PATCHPaymentMethodsPaymentMethodId200ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *PaymentMethodUpdateData) GetRelationships() PaymentMethodUpdateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PaymentMethodUpdateData) GetRelationshipsOk() (*PaymentMethodUpdateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PaymentMethodUpdateData) SetRelationships(v PaymentMethodUpdateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PaymentMethodUpdateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


