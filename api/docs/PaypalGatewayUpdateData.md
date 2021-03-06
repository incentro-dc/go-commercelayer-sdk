# PaypalGatewayUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | [default to "paypal_gateways"]
**Id** | **string** | The resource&#39;s id | 
**Attributes** | [**PaypalGatewayUpdateDataAttributes**](PaypalGatewayUpdateDataAttributes.md) |  | 
**Relationships** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPaypalGatewayUpdateData

`func NewPaypalGatewayUpdateData(type_ string, id string, attributes PaypalGatewayUpdateDataAttributes, ) *PaypalGatewayUpdateData`

NewPaypalGatewayUpdateData instantiates a new PaypalGatewayUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaypalGatewayUpdateDataWithDefaults

`func NewPaypalGatewayUpdateDataWithDefaults() *PaypalGatewayUpdateData`

NewPaypalGatewayUpdateDataWithDefaults instantiates a new PaypalGatewayUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PaypalGatewayUpdateData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PaypalGatewayUpdateData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PaypalGatewayUpdateData) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *PaypalGatewayUpdateData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaypalGatewayUpdateData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaypalGatewayUpdateData) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *PaypalGatewayUpdateData) GetAttributes() PaypalGatewayUpdateDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PaypalGatewayUpdateData) GetAttributesOk() (*PaypalGatewayUpdateDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *PaypalGatewayUpdateData) SetAttributes(v PaypalGatewayUpdateDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *PaypalGatewayUpdateData) GetRelationships() map[string]interface{}`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PaypalGatewayUpdateData) GetRelationshipsOk() (*map[string]interface{}, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PaypalGatewayUpdateData) SetRelationships(v map[string]interface{})`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PaypalGatewayUpdateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


