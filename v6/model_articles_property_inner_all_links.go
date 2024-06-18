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

// ArticlesPropertyInnerAllLinks - struct for ArticlesPropertyInnerAllLinks
type ArticlesPropertyInnerAllLinks struct {
	AllLinksProperty *AllLinksProperty
}

// AllLinksPropertyAsArticlesPropertyInnerAllLinks is a convenience function that returns AllLinksProperty wrapped in ArticlesPropertyInnerAllLinks
func AllLinksPropertyAsArticlesPropertyInnerAllLinks(v *AllLinksProperty) ArticlesPropertyInnerAllLinks {
	return ArticlesPropertyInnerAllLinks{
		AllLinksProperty: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ArticlesPropertyInnerAllLinks) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AllLinksProperty
	err = newStrictDecoder(data).Decode(&dst.AllLinksProperty)
	if err == nil {
		jsonAllLinksProperty, _ := json.Marshal(dst.AllLinksProperty)
		if string(jsonAllLinksProperty) == "{}" { // empty struct
			dst.AllLinksProperty = nil
		} else {
			match++
		}
	} else {
		dst.AllLinksProperty = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AllLinksProperty = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ArticlesPropertyInnerAllLinks)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ArticlesPropertyInnerAllLinks)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ArticlesPropertyInnerAllLinks) MarshalJSON() ([]byte, error) {
	if src.AllLinksProperty != nil {
		return json.Marshal(&src.AllLinksProperty)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ArticlesPropertyInnerAllLinks) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.AllLinksProperty != nil {
		return obj.AllLinksProperty
	}

	// all schemas are nil
	return nil
}

type NullableArticlesPropertyInnerAllLinks struct {
	value *ArticlesPropertyInnerAllLinks
	isSet bool
}

func (v NullableArticlesPropertyInnerAllLinks) Get() *ArticlesPropertyInnerAllLinks {
	return v.value
}

func (v *NullableArticlesPropertyInnerAllLinks) Set(val *ArticlesPropertyInnerAllLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableArticlesPropertyInnerAllLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableArticlesPropertyInnerAllLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArticlesPropertyInnerAllLinks(val *ArticlesPropertyInnerAllLinks) *NullableArticlesPropertyInnerAllLinks {
	return &NullableArticlesPropertyInnerAllLinks{value: val, isSet: true}
}

func (v NullableArticlesPropertyInnerAllLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArticlesPropertyInnerAllLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


