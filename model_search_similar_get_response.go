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

// SearchSimilarGetResponse struct for SearchSimilarGetResponse
type SearchSimilarGetResponse struct {
	DtoResponsesMoreLikeThisResponseFailedSearchResponse *DtoResponsesMoreLikeThisResponseFailedSearchResponse
	DtoResponsesMoreLikeThisResponseSearchResponse *DtoResponsesMoreLikeThisResponseSearchResponse
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *SearchSimilarGetResponse) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into DtoResponsesMoreLikeThisResponseFailedSearchResponse
	err = json.Unmarshal(data, &dst.DtoResponsesMoreLikeThisResponseFailedSearchResponse);
	if err == nil {
		jsonDtoResponsesMoreLikeThisResponseFailedSearchResponse, _ := json.Marshal(dst.DtoResponsesMoreLikeThisResponseFailedSearchResponse)
		if string(jsonDtoResponsesMoreLikeThisResponseFailedSearchResponse) == "{}" { // empty struct
			dst.DtoResponsesMoreLikeThisResponseFailedSearchResponse = nil
		} else {
			return nil // data stored in dst.DtoResponsesMoreLikeThisResponseFailedSearchResponse, return on the first match
		}
	} else {
		dst.DtoResponsesMoreLikeThisResponseFailedSearchResponse = nil
	}

	// try to unmarshal JSON data into DtoResponsesMoreLikeThisResponseSearchResponse
	err = json.Unmarshal(data, &dst.DtoResponsesMoreLikeThisResponseSearchResponse);
	if err == nil {
		jsonDtoResponsesMoreLikeThisResponseSearchResponse, _ := json.Marshal(dst.DtoResponsesMoreLikeThisResponseSearchResponse)
		if string(jsonDtoResponsesMoreLikeThisResponseSearchResponse) == "{}" { // empty struct
			dst.DtoResponsesMoreLikeThisResponseSearchResponse = nil
		} else {
			return nil // data stored in dst.DtoResponsesMoreLikeThisResponseSearchResponse, return on the first match
		}
	} else {
		dst.DtoResponsesMoreLikeThisResponseSearchResponse = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(SearchSimilarGetResponse)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *SearchSimilarGetResponse) MarshalJSON() ([]byte, error) {
	if src.DtoResponsesMoreLikeThisResponseFailedSearchResponse != nil {
		return json.Marshal(&src.DtoResponsesMoreLikeThisResponseFailedSearchResponse)
	}

	if src.DtoResponsesMoreLikeThisResponseSearchResponse != nil {
		return json.Marshal(&src.DtoResponsesMoreLikeThisResponseSearchResponse)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableSearchSimilarGetResponse struct {
	value *SearchSimilarGetResponse
	isSet bool
}

func (v NullableSearchSimilarGetResponse) Get() *SearchSimilarGetResponse {
	return v.value
}

func (v *NullableSearchSimilarGetResponse) Set(val *SearchSimilarGetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchSimilarGetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchSimilarGetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchSimilarGetResponse(val *SearchSimilarGetResponse) *NullableSearchSimilarGetResponse {
	return &NullableSearchSimilarGetResponse{value: val, isSet: true}
}

func (v NullableSearchSimilarGetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchSimilarGetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


