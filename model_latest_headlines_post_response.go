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

// LatestHeadlinesPostResponse struct for LatestHeadlinesPostResponse
type LatestHeadlinesPostResponse struct {
	ClusteringSearchResponse *ClusteringSearchResponse
	LatestHeadlinesResponse *LatestHeadlinesResponse
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *LatestHeadlinesPostResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into ClusteringSearchResponse
	err = json.Unmarshal(data, &dst.ClusteringSearchResponse);
	if err == nil {
		jsonClusteringSearchResponse, _ := json.Marshal(dst.ClusteringSearchResponse)
		if string(jsonClusteringSearchResponse) == "{}" { // empty struct
			dst.ClusteringSearchResponse = nil
		} else {
			return nil // data stored in dst.ClusteringSearchResponse, return on the first match
		}
	} else {
		dst.ClusteringSearchResponse = nil
	}

	// try to unmarshal JSON data into LatestHeadlinesResponse
	err = json.Unmarshal(data, &dst.LatestHeadlinesResponse);
	if err == nil {
		jsonLatestHeadlinesResponse, _ := json.Marshal(dst.LatestHeadlinesResponse)
		if string(jsonLatestHeadlinesResponse) == "{}" { // empty struct
			dst.LatestHeadlinesResponse = nil
		} else {
			return nil // data stored in dst.LatestHeadlinesResponse, return on the first match
		}
	} else {
		dst.LatestHeadlinesResponse = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(LatestHeadlinesPostResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *LatestHeadlinesPostResponse) MarshalJSON() ([]byte, error) {
	if src.ClusteringSearchResponse != nil {
		return json.Marshal(&src.ClusteringSearchResponse)
	}

	if src.LatestHeadlinesResponse != nil {
		return json.Marshal(&src.LatestHeadlinesResponse)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableLatestHeadlinesPostResponse struct {
	value *LatestHeadlinesPostResponse
	isSet bool
}

func (v NullableLatestHeadlinesPostResponse) Get() *LatestHeadlinesPostResponse {
	return v.value
}

func (v *NullableLatestHeadlinesPostResponse) Set(val *LatestHeadlinesPostResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableLatestHeadlinesPostResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableLatestHeadlinesPostResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLatestHeadlinesPostResponse(val *LatestHeadlinesPostResponse) *NullableLatestHeadlinesPostResponse {
	return &NullableLatestHeadlinesPostResponse{value: val, isSet: true}
}

func (v NullableLatestHeadlinesPostResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLatestHeadlinesPostResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


