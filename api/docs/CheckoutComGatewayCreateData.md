# CheckoutComGatewayCreateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**POSTCheckoutComGateways201ResponseDataAttributes**](POSTCheckoutComGateways201ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**CheckoutComGatewayCreateDataRelationships**](CheckoutComGatewayCreateDataRelationships.md) |  | [optional] 

## Methods

### NewCheckoutComGatewayCreateData

`func NewCheckoutComGatewayCreateData(type_ interface{}, attributes POSTCheckoutComGateways201ResponseDataAttributes, ) *CheckoutComGatewayCreateData`

NewCheckoutComGatewayCreateData instantiates a new CheckoutComGatewayCreateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckoutComGatewayCreateDataWithDefaults

`func NewCheckoutComGatewayCreateDataWithDefaults() *CheckoutComGatewayCreateData`

NewCheckoutComGatewayCreateDataWithDefaults instantiates a new CheckoutComGatewayCreateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CheckoutComGatewayCreateData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CheckoutComGatewayCreateData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CheckoutComGatewayCreateData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *CheckoutComGatewayCreateData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *CheckoutComGatewayCreateData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *CheckoutComGatewayCreateData) GetAttributes() POSTCheckoutComGateways201ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CheckoutComGatewayCreateData) GetAttributesOk() (*POSTCheckoutComGateways201ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CheckoutComGatewayCreateData) SetAttributes(v POSTCheckoutComGateways201ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *CheckoutComGatewayCreateData) GetRelationships() CheckoutComGatewayCreateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *CheckoutComGatewayCreateData) GetRelationshipsOk() (*CheckoutComGatewayCreateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *CheckoutComGatewayCreateData) SetRelationships(v CheckoutComGatewayCreateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *CheckoutComGatewayCreateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


