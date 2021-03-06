# AddressData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | [default to "addresses"]
**Attributes** | [**AddressDataAttributes**](AddressDataAttributes.md) |  | 
**Relationships** | Pointer to [**AddressDataRelationships**](AddressDataRelationships.md) |  | [optional] 

## Methods

### NewAddressData

`func NewAddressData(type_ string, attributes AddressDataAttributes, ) *AddressData`

NewAddressData instantiates a new AddressData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressDataWithDefaults

`func NewAddressDataWithDefaults() *AddressData`

NewAddressDataWithDefaults instantiates a new AddressData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AddressData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AddressData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AddressData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *AddressData) GetAttributes() AddressDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *AddressData) GetAttributesOk() (*AddressDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *AddressData) SetAttributes(v AddressDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *AddressData) GetRelationships() AddressDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *AddressData) GetRelationshipsOk() (*AddressDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *AddressData) SetRelationships(v AddressDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *AddressData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


