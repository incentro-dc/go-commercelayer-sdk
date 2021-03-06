# CustomerPaymentSourceUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The resource&#39;s type | [default to "customer_payment_sources"]
**Id** | **string** | The resource&#39;s id | 
**Attributes** | [**AdyenPaymentCreateDataAttributes**](AdyenPaymentCreateDataAttributes.md) |  | 
**Relationships** | Pointer to [**CustomerPaymentSourceDataRelationships**](CustomerPaymentSourceDataRelationships.md) |  | [optional] 

## Methods

### NewCustomerPaymentSourceUpdateData

`func NewCustomerPaymentSourceUpdateData(type_ string, id string, attributes AdyenPaymentCreateDataAttributes, ) *CustomerPaymentSourceUpdateData`

NewCustomerPaymentSourceUpdateData instantiates a new CustomerPaymentSourceUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerPaymentSourceUpdateDataWithDefaults

`func NewCustomerPaymentSourceUpdateDataWithDefaults() *CustomerPaymentSourceUpdateData`

NewCustomerPaymentSourceUpdateDataWithDefaults instantiates a new CustomerPaymentSourceUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CustomerPaymentSourceUpdateData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CustomerPaymentSourceUpdateData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CustomerPaymentSourceUpdateData) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *CustomerPaymentSourceUpdateData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomerPaymentSourceUpdateData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomerPaymentSourceUpdateData) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *CustomerPaymentSourceUpdateData) GetAttributes() AdyenPaymentCreateDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CustomerPaymentSourceUpdateData) GetAttributesOk() (*AdyenPaymentCreateDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CustomerPaymentSourceUpdateData) SetAttributes(v AdyenPaymentCreateDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *CustomerPaymentSourceUpdateData) GetRelationships() CustomerPaymentSourceDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *CustomerPaymentSourceUpdateData) GetRelationshipsOk() (*CustomerPaymentSourceDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *CustomerPaymentSourceUpdateData) SetRelationships(v CustomerPaymentSourceDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *CustomerPaymentSourceUpdateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


