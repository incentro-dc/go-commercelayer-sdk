# ShippingZoneData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**GETShippingZonesShippingZoneId200ResponseDataAttributes**](GETShippingZonesShippingZoneId200ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**ShippingZoneDataRelationships**](ShippingZoneDataRelationships.md) |  | [optional] 

## Methods

### NewShippingZoneData

`func NewShippingZoneData(type_ interface{}, attributes GETShippingZonesShippingZoneId200ResponseDataAttributes, ) *ShippingZoneData`

NewShippingZoneData instantiates a new ShippingZoneData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShippingZoneDataWithDefaults

`func NewShippingZoneDataWithDefaults() *ShippingZoneData`

NewShippingZoneDataWithDefaults instantiates a new ShippingZoneData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ShippingZoneData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ShippingZoneData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ShippingZoneData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *ShippingZoneData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *ShippingZoneData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *ShippingZoneData) GetAttributes() GETShippingZonesShippingZoneId200ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ShippingZoneData) GetAttributesOk() (*GETShippingZonesShippingZoneId200ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ShippingZoneData) SetAttributes(v GETShippingZonesShippingZoneId200ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *ShippingZoneData) GetRelationships() ShippingZoneDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *ShippingZoneData) GetRelationshipsOk() (*ShippingZoneDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *ShippingZoneData) SetRelationships(v ShippingZoneDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *ShippingZoneData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


