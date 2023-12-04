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

// AuthorsProperty struct for AuthorsProperty
type AuthorsProperty struct {
	[]string *[]string
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AuthorsProperty) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into []string
	err = json.Unmarshal(data, &dst.[]string);
	if err == nil {
		json[]string, _ := json.Marshal(dst.[]string)
		if string(json[]string) == "{}" { // empty struct
			dst.[]string = nil
		} else {
			return nil // data stored in dst.[]string, return on the first match
		}
	} else {
		dst.[]string = nil
	}

	// try to unmarshal JSON data into string
	err = json.Unmarshal(data, &dst.string);
	if err == nil {
		jsonstring, _ := json.Marshal(dst.string)
		if string(jsonstring) == "{}" { // empty struct
			dst.string = nil
		} else {
			return nil // data stored in dst.string, return on the first match
		}
	} else {
		dst.string = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(AuthorsProperty)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AuthorsProperty) MarshalJSON() ([]byte, error) {
	if src.[]string != nil {
		return json.Marshal(&src.[]string)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAuthorsProperty struct {
	value *AuthorsProperty
	isSet bool
}

func (v NullableAuthorsProperty) Get() *AuthorsProperty {
	return v.value
}

func (v *NullableAuthorsProperty) Set(val *AuthorsProperty) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorsProperty) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorsProperty) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorsProperty(val *AuthorsProperty) *NullableAuthorsProperty {
	return &NullableAuthorsProperty{value: val, isSet: true}
}

func (v NullableAuthorsProperty) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorsProperty) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


