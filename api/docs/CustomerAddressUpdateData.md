# CustomerAddressUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | [default to "customer_addresses"]
**Id** | **string** | The resource&#39;s id | 
**Attributes** | [**AdyenPaymentCreateDataAttributes**](AdyenPaymentCreateDataAttributes.md) |  | 
**Relationships** | Pointer to [**CustomerAddressDataRelationships**](CustomerAddressDataRelationships.md) |  | [optional] 

## Methods

### NewCustomerAddressUpdateData

`func NewCustomerAddressUpdateData(type_ string, id string, attributes AdyenPaymentCreateDataAttributes, ) *CustomerAddressUpdateData`

NewCustomerAddressUpdateData instantiates a new CustomerAddressUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerAddressUpdateDataWithDefaults

`func NewCustomerAddressUpdateDataWithDefaults() *CustomerAddressUpdateData`

NewCustomerAddressUpdateDataWithDefaults instantiates a new CustomerAddressUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CustomerAddressUpdateData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CustomerAddressUpdateData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CustomerAddressUpdateData) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *CustomerAddressUpdateData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomerAddressUpdateData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomerAddressUpdateData) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *CustomerAddressUpdateData) GetAttributes() AdyenPaymentCreateDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CustomerAddressUpdateData) GetAttributesOk() (*AdyenPaymentCreateDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CustomerAddressUpdateData) SetAttributes(v AdyenPaymentCreateDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *CustomerAddressUpdateData) GetRelationships() CustomerAddressDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *CustomerAddressUpdateData) GetRelationshipsOk() (*CustomerAddressDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *CustomerAddressUpdateData) SetRelationships(v CustomerAddressDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *CustomerAddressUpdateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

