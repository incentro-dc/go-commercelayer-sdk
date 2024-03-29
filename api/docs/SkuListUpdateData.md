# SkuListUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Id** | **interface{}** | The resource&#39;s id | 
**Attributes** | [**PATCHSkuListsSkuListId200ResponseDataAttributes**](PATCHSkuListsSkuListId200ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**CouponRecipientCreateDataRelationships**](CouponRecipientCreateDataRelationships.md) |  | [optional] 

## Methods

### NewSkuListUpdateData

`func NewSkuListUpdateData(type_ interface{}, id interface{}, attributes PATCHSkuListsSkuListId200ResponseDataAttributes, ) *SkuListUpdateData`

NewSkuListUpdateData instantiates a new SkuListUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSkuListUpdateDataWithDefaults

`func NewSkuListUpdateDataWithDefaults() *SkuListUpdateData`

NewSkuListUpdateDataWithDefaults instantiates a new SkuListUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SkuListUpdateData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SkuListUpdateData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SkuListUpdateData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *SkuListUpdateData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *SkuListUpdateData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetId

`func (o *SkuListUpdateData) GetId() interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SkuListUpdateData) GetIdOk() (*interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SkuListUpdateData) SetId(v interface{})`

SetId sets Id field to given value.


### SetIdNil

`func (o *SkuListUpdateData) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *SkuListUpdateData) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetAttributes

`func (o *SkuListUpdateData) GetAttributes() PATCHSkuListsSkuListId200ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *SkuListUpdateData) GetAttributesOk() (*PATCHSkuListsSkuListId200ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *SkuListUpdateData) SetAttributes(v PATCHSkuListsSkuListId200ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *SkuListUpdateData) GetRelationships() CouponRecipientCreateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *SkuListUpdateData) GetRelationshipsOk() (*CouponRecipientCreateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *SkuListUpdateData) SetRelationships(v CouponRecipientCreateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *SkuListUpdateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


