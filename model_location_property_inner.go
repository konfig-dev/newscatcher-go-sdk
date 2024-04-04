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
	"fmt"
)

// LocationPropertyInner struct for LocationPropertyInner
type LocationPropertyInner struct {
	Int32 *int32
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *LocationPropertyInner) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into int32
	err = json.Unmarshal(data, &dst.Int32);
	if err == nil {
		jsonInt32, _ := json.Marshal(dst.Int32)
		if string(jsonInt32) == "{}" { // empty struct
			dst.Int32 = nil
		} else {
			return nil // data stored in dst.Int32, return on the first match
		}
	} else {
		dst.Int32 = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(LocationPropertyInner)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *LocationPropertyInner) MarshalJSON() ([]byte, error) {
	if src.Int32 != nil {
		return json.Marshal(&src.Int32)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableLocationPropertyInner struct {
	value *LocationPropertyInner
	isSet bool
}

func (v NullableLocationPropertyInner) Get() *LocationPropertyInner {
	return v.value
}

func (v *NullableLocationPropertyInner) Set(val *LocationPropertyInner) {
	v.value = val
	v.isSet = true
}

func (v NullableLocationPropertyInner) IsSet() bool {
	return v.isSet
}

func (v *NullableLocationPropertyInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocationPropertyInner(val *LocationPropertyInner) *NullableLocationPropertyInner {
	return &NullableLocationPropertyInner{value: val, isSet: true}
}

func (v NullableLocationPropertyInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocationPropertyInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


