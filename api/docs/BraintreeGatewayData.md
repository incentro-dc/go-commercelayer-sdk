# BraintreeGatewayData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes**](GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**BraintreeGatewayDataRelationships**](BraintreeGatewayDataRelationships.md) |  | [optional] 

## Methods

### NewBraintreeGatewayData

`func NewBraintreeGatewayData(type_ interface{}, attributes GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes, ) *BraintreeGatewayData`

NewBraintreeGatewayData instantiates a new BraintreeGatewayData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBraintreeGatewayDataWithDefaults

`func NewBraintreeGatewayDataWithDefaults() *BraintreeGatewayData`

NewBraintreeGatewayDataWithDefaults instantiates a new BraintreeGatewayData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BraintreeGatewayData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BraintreeGatewayData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BraintreeGatewayData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *BraintreeGatewayData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *BraintreeGatewayData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *BraintreeGatewayData) GetAttributes() GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *BraintreeGatewayData) GetAttributesOk() (*GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *BraintreeGatewayData) SetAttributes(v GETBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *BraintreeGatewayData) GetRelationships() BraintreeGatewayDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *BraintreeGatewayData) GetRelationshipsOk() (*BraintreeGatewayDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *BraintreeGatewayData) SetRelationships(v BraintreeGatewayDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *BraintreeGatewayData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


