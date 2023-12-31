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

// SearchGetResponse struct for SearchGetResponse
type SearchGetResponse struct {
	ClusteringSearchResponse *ClusteringSearchResponse
	DtoResponsesSearchResponseFailedSearchResponse *DtoResponsesSearchResponseFailedSearchResponse
	DtoResponsesSearchResponseSearchResponse *DtoResponsesSearchResponseSearchResponse
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *SearchGetResponse) UnmarshalJSON(data []byte) error {
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

	// try to unmarshal JSON data into DtoResponsesSearchResponseFailedSearchResponse
	err = json.Unmarshal(data, &dst.DtoResponsesSearchResponseFailedSearchResponse);
	if err == nil {
		jsonDtoResponsesSearchResponseFailedSearchResponse, _ := json.Marshal(dst.DtoResponsesSearchResponseFailedSearchResponse)
		if string(jsonDtoResponsesSearchResponseFailedSearchResponse) == "{}" { // empty struct
			dst.DtoResponsesSearchResponseFailedSearchResponse = nil
		} else {
			return nil // data stored in dst.DtoResponsesSearchResponseFailedSearchResponse, return on the first match
		}
	} else {
		dst.DtoResponsesSearchResponseFailedSearchResponse = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(SearchGetResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *SearchGetResponse) MarshalJSON() ([]byte, error) {
	if src.ClusteringSearchResponse != nil {
		return json.Marshal(&src.ClusteringSearchResponse)
	}

	if src.DtoResponsesSearchResponseFailedSearchResponse != nil {
		return json.Marshal(&src.DtoResponsesSearchResponseFailedSearchResponse)
	}

	if src.DtoResponsesSearchResponseSearchResponse != nil {
		return json.Marshal(&src.DtoResponsesSearchResponseSearchResponse)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableSearchGetResponse struct {
	value *SearchGetResponse
	isSet bool
}

func (v NullableSearchGetResponse) Get() *SearchGetResponse {
	return v.value
}

func (v *NullableSearchGetResponse) Set(val *SearchGetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchGetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchGetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchGetResponse(val *SearchGetResponse) *NullableSearchGetResponse {
	return &NullableSearchGetResponse{value: val, isSet: true}
}

func (v NullableSearchGetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchGetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


