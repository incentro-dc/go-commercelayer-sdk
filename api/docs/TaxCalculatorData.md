# TaxCalculatorData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**GETManualTaxCalculatorsManualTaxCalculatorId200ResponseDataAttributes**](GETManualTaxCalculatorsManualTaxCalculatorId200ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**ExternalTaxCalculatorDataRelationships**](ExternalTaxCalculatorDataRelationships.md) |  | [optional] 

## Methods

### NewTaxCalculatorData

`func NewTaxCalculatorData(type_ interface{}, attributes GETManualTaxCalculatorsManualTaxCalculatorId200ResponseDataAttributes, ) *TaxCalculatorData`

NewTaxCalculatorData instantiates a new TaxCalculatorData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxCalculatorDataWithDefaults

`func NewTaxCalculatorDataWithDefaults() *TaxCalculatorData`

NewTaxCalculatorDataWithDefaults instantiates a new TaxCalculatorData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TaxCalculatorData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TaxCalculatorData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TaxCalculatorData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *TaxCalculatorData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *TaxCalculatorData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *TaxCalculatorData) GetAttributes() GETManualTaxCalculatorsManualTaxCalculatorId200ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *TaxCalculatorData) GetAttributesOk() (*GETManualTaxCalculatorsManualTaxCalculatorId200ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *TaxCalculatorData) SetAttributes(v GETManualTaxCalculatorsManualTaxCalculatorId200ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *TaxCalculatorData) GetRelationships() ExternalTaxCalculatorDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *TaxCalculatorData) GetRelationshipsOk() (*ExternalTaxCalculatorDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *TaxCalculatorData) SetRelationships(v ExternalTaxCalculatorDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *TaxCalculatorData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


