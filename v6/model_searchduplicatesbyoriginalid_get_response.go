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

// SearchduplicatesbyoriginalidGetResponse struct for SearchduplicatesbyoriginalidGetResponse
type SearchduplicatesbyoriginalidGetResponse struct {
	ClusteringSearchResponse *ClusteringSearchResponse
	DtoResponsesSearchResponseSearchResponse *DtoResponsesSearchResponseSearchResponse
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *SearchduplicatesbyoriginalidGetResponse) UnmarshalJSON(data []byte) error {
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

	// try to unmarshal JSON data into DtoResponsesSearchResponseSearchResponse
	err = json.Unmarshal(data, &dst.DtoResponsesSearchResponseSearchResponse);
	if err == nil {
		jsonDtoResponsesSearchResponseSearchResponse, _ := json.Marshal(dst.DtoResponsesSearchResponseSearchResponse)
		if string(jsonDtoResponsesSearchResponseSearchResponse) == "{}" { // empty struct
			dst.DtoResponsesSearchResponseSearchResponse = nil
		} else {
			return nil // data stored in dst.DtoResponsesSearchResponseSearchResponse, return on the first match
		}
	} else {
		dst.DtoResponsesSearchResponseSearchResponse = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(SearchduplicatesbyoriginalidGetResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *SearchduplicatesbyoriginalidGetResponse) MarshalJSON() ([]byte, error) {
	if src.ClusteringSearchResponse != nil {
		return json.Marshal(&src.ClusteringSearchResponse)
	}

	if src.DtoResponsesSearchResponseSearchResponse != nil {
		return json.Marshal(&src.DtoResponsesSearchResponseSearchResponse)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableSearchduplicatesbyoriginalidGetResponse struct {
	value *SearchduplicatesbyoriginalidGetResponse
	isSet bool
}

func (v NullableSearchduplicatesbyoriginalidGetResponse) Get() *SearchduplicatesbyoriginalidGetResponse {
	return v.value
}

func (v *NullableSearchduplicatesbyoriginalidGetResponse) Set(val *SearchduplicatesbyoriginalidGetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchduplicatesbyoriginalidGetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchduplicatesbyoriginalidGetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchduplicatesbyoriginalidGetResponse(val *SearchduplicatesbyoriginalidGetResponse) *NullableSearchduplicatesbyoriginalidGetResponse {
	return &NullableSearchduplicatesbyoriginalidGetResponse{value: val, isSet: true}
}

func (v NullableSearchduplicatesbyoriginalidGetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchduplicatesbyoriginalidGetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


