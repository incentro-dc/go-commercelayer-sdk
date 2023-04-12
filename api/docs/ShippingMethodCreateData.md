# ShippingMethodCreateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**POSTShippingMethods201ResponseDataAttributes**](POSTShippingMethods201ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**ShippingMethodCreateDataRelationships**](ShippingMethodCreateDataRelationships.md) |  | [optional] 

## Methods

### NewShippingMethodCreateData

`func NewShippingMethodCreateData(type_ interface{}, attributes POSTShippingMethods201ResponseDataAttributes, ) *ShippingMethodCreateData`

NewShippingMethodCreateData instantiates a new ShippingMethodCreateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShippingMethodCreateDataWithDefaults

`func NewShippingMethodCreateDataWithDefaults() *ShippingMethodCreateData`

NewShippingMethodCreateDataWithDefaults instantiates a new ShippingMethodCreateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ShippingMethodCreateData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ShippingMethodCreateData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ShippingMethodCreateData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *ShippingMethodCreateData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *ShippingMethodCreateData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *ShippingMethodCreateData) GetAttributes() POSTShippingMethods201ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ShippingMethodCreateData) GetAttributesOk() (*POSTShippingMethods201ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ShippingMethodCreateData) SetAttributes(v POSTShippingMethods201ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *ShippingMethodCreateData) GetRelationships() ShippingMethodCreateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *ShippingMethodCreateData) GetRelationshipsOk() (*ShippingMethodCreateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *ShippingMethodCreateData) SetRelationships(v ShippingMethodCreateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *ShippingMethodCreateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

