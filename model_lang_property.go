/*
NewsCatcher-V3 Production API

<img src='https://uploads-ssl.webflow.com/6429857b17973b636c2195c5/646c6f1eb774ff2f2997bec5_newscatcher_.svg' width='286' height='35' /> <br>  <br>Visit our website  <a href='https://newscatcherapi.com'>https://newscatcherapi.com</a>

API version: Beta-3.0.0
Contact: maksym@newscatcherapi.com
*/

// Code generated by Konfig (https://konfigthis.com); DO NOT EDIT.

package newscatcherapi

import (
	"encoding/json"
	"fmt"
)

// LangProperty struct for LangProperty
type LangProperty struct {
	StringArray *[]string
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *LangProperty) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into []string
	err = json.Unmarshal(data, &dst.StringArray);
	if err == nil {
		jsonStringArray, _ := json.Marshal(dst.StringArray)
		if string(jsonStringArray) == "{}" { // empty struct
			dst.StringArray = nil
		} else {
			return nil // data stored in dst.StringArray, return on the first match
		}
	} else {
		dst.StringArray = nil
	}

	// try to unmarshal JSON data into string
	err = json.Unmarshal(data, &dst.String);
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			return nil // data stored in dst.String, return on the first match
		}
	} else {
		dst.String = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(LangProperty)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *LangProperty) MarshalJSON() ([]byte, error) {
	if src.StringArray != nil {
		return json.Marshal(&src.StringArray)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableLangProperty struct {
	value *LangProperty
	isSet bool
}

func (v NullableLangProperty) Get() *LangProperty {
	return v.value
}

func (v *NullableLangProperty) Set(val *LangProperty) {
	v.value = val
	v.isSet = true
}

func (v NullableLangProperty) IsSet() bool {
	return v.isSet
}

func (v *NullableLangProperty) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLangProperty(val *LangProperty) *NullableLangProperty {
	return &NullableLangProperty{value: val, isSet: true}
}

func (v NullableLangProperty) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLangProperty) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


