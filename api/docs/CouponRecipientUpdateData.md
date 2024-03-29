# CouponRecipientUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Id** | **interface{}** | The resource&#39;s id | 
**Attributes** | [**PATCHCouponRecipientsCouponRecipientId200ResponseDataAttributes**](PATCHCouponRecipientsCouponRecipientId200ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**CouponRecipientCreateDataRelationships**](CouponRecipientCreateDataRelationships.md) |  | [optional] 

## Methods

### NewCouponRecipientUpdateData

`func NewCouponRecipientUpdateData(type_ interface{}, id interface{}, attributes PATCHCouponRecipientsCouponRecipientId200ResponseDataAttributes, ) *CouponRecipientUpdateData`

NewCouponRecipientUpdateData instantiates a new CouponRecipientUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCouponRecipientUpdateDataWithDefaults

`func NewCouponRecipientUpdateDataWithDefaults() *CouponRecipientUpdateData`

NewCouponRecipientUpdateDataWithDefaults instantiates a new CouponRecipientUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CouponRecipientUpdateData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CouponRecipientUpdateData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CouponRecipientUpdateData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *CouponRecipientUpdateData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *CouponRecipientUpdateData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetId

`func (o *CouponRecipientUpdateData) GetId() interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CouponRecipientUpdateData) GetIdOk() (*interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CouponRecipientUpdateData) SetId(v interface{})`

SetId sets Id field to given value.


### SetIdNil

`func (o *CouponRecipientUpdateData) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *CouponRecipientUpdateData) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetAttributes

`func (o *CouponRecipientUpdateData) GetAttributes() PATCHCouponRecipientsCouponRecipientId200ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CouponRecipientUpdateData) GetAttributesOk() (*PATCHCouponRecipientsCouponRecipientId200ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CouponRecipientUpdateData) SetAttributes(v PATCHCouponRecipientsCouponRecipientId200ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *CouponRecipientUpdateData) GetRelationships() CouponRecipientCreateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *CouponRecipientUpdateData) GetRelationshipsOk() (*CouponRecipientCreateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *CouponRecipientUpdateData) SetRelationships(v CouponRecipientCreateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *CouponRecipientUpdateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


