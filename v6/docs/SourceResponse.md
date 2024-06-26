# SourceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** |  | 
**Sources** | [**[]SourcesPropertyInner**](SourcesPropertyInner.md) |  | 
**UserInput** | **map[string]interface{}** |  | 

## Methods

### NewSourceResponse

`func NewSourceResponse(message string, sources []SourcesPropertyInner, userInput map[string]interface{}, ) *SourceResponse`

NewSourceResponse instantiates a new SourceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceResponseWithDefaults

`func NewSourceResponseWithDefaults() *SourceResponse`

NewSourceResponseWithDefaults instantiates a new SourceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *SourceResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SourceResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SourceResponse) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetSources

`func (o *SourceResponse) GetSources() []SourcesPropertyInner`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *SourceResponse) GetSourcesOk() (*[]SourcesPropertyInner, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *SourceResponse) SetSources(v []SourcesPropertyInner)`

SetSources sets Sources field to given value.


### GetUserInput

`func (o *SourceResponse) GetUserInput() map[string]interface{}`

GetUserInput returns the UserInput field if non-nil, zero value otherwise.

### GetUserInputOk

`func (o *SourceResponse) GetUserInputOk() (*map[string]interface{}, bool)`

GetUserInputOk returns a tuple with the UserInput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserInput

`func (o *SourceResponse) SetUserInput(v map[string]interface{})`

SetUserInput sets UserInput field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


