# AddressUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Id** | **interface{}** | The resource&#39;s id | 
**Attributes** | [**PATCHAddressesAddressId200ResponseDataAttributes**](PATCHAddressesAddressId200ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**AddressCreateDataRelationships**](AddressCreateDataRelationships.md) |  | [optional] 

## Methods

### NewAddressUpdateData

`func NewAddressUpdateData(type_ interface{}, id interface{}, attributes PATCHAddressesAddressId200ResponseDataAttributes, ) *AddressUpdateData`

NewAddressUpdateData instantiates a new AddressUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressUpdateDataWithDefaults

`func NewAddressUpdateDataWithDefaults() *AddressUpdateData`

NewAddressUpdateDataWithDefaults instantiates a new AddressUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AddressUpdateData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AddressUpdateData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AddressUpdateData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *AddressUpdateData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *AddressUpdateData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetId

`func (o *AddressUpdateData) GetId() interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddressUpdateData) GetIdOk() (*interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddressUpdateData) SetId(v interface{})`

SetId sets Id field to given value.


### SetIdNil

`func (o *AddressUpdateData) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *AddressUpdateData) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetAttributes

`func (o *AddressUpdateData) GetAttributes() PATCHAddressesAddressId200ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *AddressUpdateData) GetAttributesOk() (*PATCHAddressesAddressId200ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *AddressUpdateData) SetAttributes(v PATCHAddressesAddressId200ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *AddressUpdateData) GetRelationships() AddressCreateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *AddressUpdateData) GetRelationshipsOk() (*AddressCreateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *AddressUpdateData) SetRelationships(v AddressCreateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *AddressUpdateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


