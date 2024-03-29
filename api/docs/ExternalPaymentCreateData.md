# ExternalPaymentCreateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**POSTExternalPayments201ResponseDataAttributes**](POSTExternalPayments201ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**AdyenPaymentCreateDataRelationships**](AdyenPaymentCreateDataRelationships.md) |  | [optional] 

## Methods

### NewExternalPaymentCreateData

`func NewExternalPaymentCreateData(type_ interface{}, attributes POSTExternalPayments201ResponseDataAttributes, ) *ExternalPaymentCreateData`

NewExternalPaymentCreateData instantiates a new ExternalPaymentCreateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalPaymentCreateDataWithDefaults

`func NewExternalPaymentCreateDataWithDefaults() *ExternalPaymentCreateData`

NewExternalPaymentCreateDataWithDefaults instantiates a new ExternalPaymentCreateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ExternalPaymentCreateData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ExternalPaymentCreateData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ExternalPaymentCreateData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *ExternalPaymentCreateData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *ExternalPaymentCreateData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *ExternalPaymentCreateData) GetAttributes() POSTExternalPayments201ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ExternalPaymentCreateData) GetAttributesOk() (*POSTExternalPayments201ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ExternalPaymentCreateData) SetAttributes(v POSTExternalPayments201ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *ExternalPaymentCreateData) GetRelationships() AdyenPaymentCreateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *ExternalPaymentCreateData) GetRelationshipsOk() (*AdyenPaymentCreateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *ExternalPaymentCreateData) SetRelationships(v AdyenPaymentCreateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *ExternalPaymentCreateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


