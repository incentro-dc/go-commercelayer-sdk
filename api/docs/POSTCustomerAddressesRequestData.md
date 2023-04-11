# POSTCustomerAddressesRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**POSTAdyenPaymentsRequestDataAttributes**](POSTAdyenPaymentsRequestDataAttributes.md) |  | 
**Relationships** | Pointer to [**POSTCustomerAddressesRequestDataRelationships**](POSTCustomerAddressesRequestDataRelationships.md) |  | [optional] 

## Methods

### NewPOSTCustomerAddressesRequestData

`func NewPOSTCustomerAddressesRequestData(type_ interface{}, attributes POSTAdyenPaymentsRequestDataAttributes, ) *POSTCustomerAddressesRequestData`

NewPOSTCustomerAddressesRequestData instantiates a new POSTCustomerAddressesRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTCustomerAddressesRequestDataWithDefaults

`func NewPOSTCustomerAddressesRequestDataWithDefaults() *POSTCustomerAddressesRequestData`

NewPOSTCustomerAddressesRequestDataWithDefaults instantiates a new POSTCustomerAddressesRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *POSTCustomerAddressesRequestData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *POSTCustomerAddressesRequestData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *POSTCustomerAddressesRequestData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *POSTCustomerAddressesRequestData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *POSTCustomerAddressesRequestData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *POSTCustomerAddressesRequestData) GetAttributes() POSTAdyenPaymentsRequestDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *POSTCustomerAddressesRequestData) GetAttributesOk() (*POSTAdyenPaymentsRequestDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *POSTCustomerAddressesRequestData) SetAttributes(v POSTAdyenPaymentsRequestDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *POSTCustomerAddressesRequestData) GetRelationships() POSTCustomerAddressesRequestDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *POSTCustomerAddressesRequestData) GetRelationshipsOk() (*POSTCustomerAddressesRequestDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *POSTCustomerAddressesRequestData) SetRelationships(v POSTCustomerAddressesRequestDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *POSTCustomerAddressesRequestData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

