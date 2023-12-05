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
	"time"
	"fmt"
)

// To struct for To
type To struct {
	String *string
	Time_Time *time.Time
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *To) UnmarshalJSON(data []byte) error {
	var err error
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

	// try to unmarshal JSON data into time.Time
	err = json.Unmarshal(data, &dst.Time_Time);
	if err == nil {
		jsonTime_Time, _ := json.Marshal(dst.Time_Time)
		if string(jsonTime_Time) == "{}" { // empty struct
			dst.Time_Time = nil
		} else {
			return nil // data stored in dst.Time_Time, return on the first match
		}
	} else {
		dst.Time_Time = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(To)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *To) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	if src.Time_Time != nil {
		return json.Marshal(&src.Time_Time)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableTo struct {
	value *To
	isSet bool
}

func (v NullableTo) Get() *To {
	return v.value
}

func (v *NullableTo) Set(val *To) {
	v.value = val
	v.isSet = true
}

func (v NullableTo) IsSet() bool {
	return v.isSet
}

func (v *NullableTo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTo(val *To) *NullableTo {
	return &NullableTo{value: val, isSet: true}
}

func (v NullableTo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


