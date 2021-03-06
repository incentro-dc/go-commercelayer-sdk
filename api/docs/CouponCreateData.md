# CouponCreateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | [default to "coupons"]
**Attributes** | [**CouponCreateDataAttributes**](CouponCreateDataAttributes.md) |  | 
**Relationships** | Pointer to [**CouponCreateDataRelationships**](CouponCreateDataRelationships.md) |  | [optional] 

## Methods

### NewCouponCreateData

`func NewCouponCreateData(type_ string, attributes CouponCreateDataAttributes, ) *CouponCreateData`

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

`func (o *CouponCreateData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CouponCreateData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CouponCreateData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *CouponCreateData) GetAttributes() CouponCreateDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CouponCreateData) GetAttributesOk() (*CouponCreateDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CouponCreateData) SetAttributes(v CouponCreateDataAttributes)`

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


