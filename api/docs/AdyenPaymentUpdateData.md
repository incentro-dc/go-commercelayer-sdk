# AdyenPaymentUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | [default to "adyen_payments"]
**Id** | **string** | The resource&#39;s id | 
**Attributes** | [**AdyenPaymentUpdateDataAttributes**](AdyenPaymentUpdateDataAttributes.md) |  | 
**Relationships** | Pointer to [**AdyenPaymentUpdateDataRelationships**](AdyenPaymentUpdateDataRelationships.md) |  | [optional] 

## Methods

### NewAdyenPaymentUpdateData

`func NewAdyenPaymentUpdateData(type_ string, id string, attributes AdyenPaymentUpdateDataAttributes, ) *AdyenPaymentUpdateData`

NewAdyenPaymentUpdateData instantiates a new AdyenPaymentUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdyenPaymentUpdateDataWithDefaults

`func NewAdyenPaymentUpdateDataWithDefaults() *AdyenPaymentUpdateData`

NewAdyenPaymentUpdateDataWithDefaults instantiates a new AdyenPaymentUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AdyenPaymentUpdateData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AdyenPaymentUpdateData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AdyenPaymentUpdateData) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *AdyenPaymentUpdateData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AdyenPaymentUpdateData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AdyenPaymentUpdateData) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *AdyenPaymentUpdateData) GetAttributes() AdyenPaymentUpdateDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *AdyenPaymentUpdateData) GetAttributesOk() (*AdyenPaymentUpdateDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *AdyenPaymentUpdateData) SetAttributes(v AdyenPaymentUpdateDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *AdyenPaymentUpdateData) GetRelationships() AdyenPaymentUpdateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *AdyenPaymentUpdateData) GetRelationshipsOk() (*AdyenPaymentUpdateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *AdyenPaymentUpdateData) SetRelationships(v AdyenPaymentUpdateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *AdyenPaymentUpdateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


