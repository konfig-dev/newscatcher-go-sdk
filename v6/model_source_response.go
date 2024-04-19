/*
NewsCatcher-V3 Production API

<img src='https://uploads-ssl.webflow.com/6429857b17973b636c2195c5/646c6f1eb774ff2f2997bec5_newscatcher_.svg' width='286' height='35' /> <br>  <br>Visit our website  <a href='https://newscatcherapi.com'>https://newscatcherapi.com</a>

API version: 3.2.16
Contact: maksym@newscatcherapi.com
*/

// Code generated by Konfig (https://konfigthis.com); DO NOT EDIT.

package newscatcherapi

import (
	"encoding/json"
)

// SourceResponse SourceResponse DTO class.
type SourceResponse struct {
	Message string `json:"message"`
	Sources []SourcesPropertyInner `json:"sources"`
	UserInput map[string]interface{} `json:"user_input"`
}

// NewSourceResponse instantiates a new SourceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSourceResponse(message string, sources []SourcesPropertyInner, userInput map[string]interface{}) *SourceResponse {
	this := SourceResponse{}
	this.Message = message
	this.Sources = sources
	this.UserInput = userInput
	return &this
}

// NewSourceResponseWithDefaults instantiates a new SourceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSourceResponseWithDefaults() *SourceResponse {
	this := SourceResponse{}
	return &this
}

// GetMessage returns the Message field value
func (o *SourceResponse) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *SourceResponse) GetMessageOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *SourceResponse) SetMessage(v string) {
	o.Message = v
}

// GetSources returns the Sources field value
func (o *SourceResponse) GetSources() []SourcesPropertyInner {
	if o == nil {
		var ret []SourcesPropertyInner
		return ret
	}

	return o.Sources
}

// GetSourcesOk returns a tuple with the Sources field value
// and a boolean to check if the value has been set.
func (o *SourceResponse) GetSourcesOk() ([]SourcesPropertyInner, bool) {
	if o == nil {
    return nil, false
	}
	return o.Sources, true
}

// SetSources sets field value
func (o *SourceResponse) SetSources(v []SourcesPropertyInner) {
	o.Sources = v
}

// GetUserInput returns the UserInput field value
func (o *SourceResponse) GetUserInput() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.UserInput
}

// GetUserInputOk returns a tuple with the UserInput field value
// and a boolean to check if the value has been set.
func (o *SourceResponse) GetUserInputOk() (map[string]interface{}, bool) {
	if o == nil {
    return map[string]interface{}{}, false
	}
	return o.UserInput, true
}

// SetUserInput sets field value
func (o *SourceResponse) SetUserInput(v map[string]interface{}) {
	o.UserInput = v
}

func (o SourceResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["message"] = o.Message
	}
	if true {
		toSerialize["sources"] = o.Sources
	}
	if true {
		toSerialize["user_input"] = o.UserInput
	}
	return json.Marshal(toSerialize)
}

type NullableSourceResponse struct {
	value *SourceResponse
	isSet bool
}

func (v NullableSourceResponse) Get() *SourceResponse {
	return v.value
}

func (v *NullableSourceResponse) Set(val *SourceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSourceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSourceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourceResponse(val *SourceResponse) *NullableSourceResponse {
	return &NullableSourceResponse{value: val, isSet: true}
}

func (v NullableSourceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


