# CouponCreateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**POSTCoupons201ResponseDataAttributes**](POSTCoupons201ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**CouponCreateDataRelationships**](CouponCreateDataRelationships.md) |  | [optional] 

## Methods

### NewCouponCreateData

`func NewCouponCreateData(type_ interface{}, attributes POSTCoupons201ResponseDataAttributes, ) *CouponCreateData`

NewCouponCreateData instantiates a new CouponCreateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCouponCreateDataWithDefaults

`func NewCouponCreateDataWithDefaults() *CouponCreateData`

NewCouponCreateDataWithDefaults instantiates a new CouponCreateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CouponCreateData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CouponCreateData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CouponCreateData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *CouponCreateData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *CouponCreateData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *CouponCreateData) GetAttributes() POSTCoupons201ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CouponCreateData) GetAttributesOk() (*POSTCoupons201ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CouponCreateData) SetAttributes(v POSTCoupons201ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *CouponCreateData) GetRelationships() CouponCreateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *CouponCreateData) GetRelationshipsOk() (*CouponCreateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *CouponCreateData) SetRelationships(v CouponCreateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *CouponCreateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


