# SearchJobMessages

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fields** | Pointer to [**[]SearchJobField**](SearchJobField.md) |  | [optional] 
**Messages** | Pointer to [**[]SearchJobMessagesMap**](SearchJobMessagesMap.md) |  | [optional] 

## Methods

### NewSearchJobMessages

`func NewSearchJobMessages() *SearchJobMessages`

NewSearchJobMessages instantiates a new SearchJobMessages object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchJobMessagesWithDefaults

`func NewSearchJobMessagesWithDefaults() *SearchJobMessages`

NewSearchJobMessagesWithDefaults instantiates a new SearchJobMessages object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFields

`func (o *SearchJobMessages) GetFields() []SearchJobField`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *SearchJobMessages) GetFieldsOk() (*[]SearchJobField, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *SearchJobMessages) SetFields(v []SearchJobField)`

SetFields sets Fields field to given value.

### HasFields

`func (o *SearchJobMessages) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetMessages

`func (o *SearchJobMessages) GetMessages() []SearchJobMessagesMap`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *SearchJobMessages) GetMessagesOk() (*[]SearchJobMessagesMap, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *SearchJobMessages) SetMessages(v []SearchJobMessagesMap)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *SearchJobMessages) HasMessages() bool`

HasMessages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


