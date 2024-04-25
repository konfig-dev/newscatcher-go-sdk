/*
NewsCatcher-V3 Production API

Testing SearchApiService

*/

// Code generated by Konfig (https://konfigthis.com);

package newscatcherapi

import (
    "testing"
    // "github.com/stretchr/testify/assert"
    // "github.com/stretchr/testify/require"
    // newscatcherapi "github.com/konfig-dev/newscatcher-go-sdk"
)

func Test_newscatcherapi_SearchApiService(t *testing.T) {

    // configuration := newscatcherapi.NewConfiguration()
    // configuration.SetHost("http://127.0.0.1:4010")
    /* 
    configuration.SetApiKey("X_API_TOKEN")
    client := newscatcherapi.NewAPIClient(configuration)
    */

    t.Run("Test SearchApiService Get", func(t *testing.T) {
        /* TODO: ENG-1367 Fix parameter values for Go SDK generated tests
        request := client.SearchApi.Get(
            "q_example",
        )
        request.SearchIn(""title_content"")
        request.PredefinedSources()
        request.Sources()
        request.NotSources()
        request.Lang()
        request.NotLang()
        request.Countries()
        request.NotCountries()
        request.NotAuthorName()
        request.From(from)
        request.To(to)
        request.PublishedDatePrecision("publishedDatePrecision_example")
        request.ByParseDate(false)
        request.SortBy(""relevancy"")
        request.RankedOnly("rankedOnly_example")
        request.FromRank(56)
        request.ToRank(56)
        request.IsHeadline(true)
        request.IsPaidContent(true)
        request.ParentUrl()
        request.AllLinks()
        request.AllDomainLinks()
        request.WordCountMin(56)
        request.WordCountMax(56)
        request.Page(1)
        request.PageSize(100)
        request.ClusteringVariable("clusteringVariable_example")
        request.ClusteringEnabled(true)
        request.ClusteringThreshold(8.14)
        request.IncludeNlpData(true)
        request.HasNlp(true)
        request.Theme("theme_example")
        request.NotTheme("notTheme_example")
        request.ORGEntityName("oRGEntityName_example")
        request.PEREntityName("pEREntityName_example")
        request.LOCEntityName("lOCEntityName_example")
        request.MISCEntityName("mISCEntityName_example")
        request.TitleSentimentMin(8.14)
        request.TitleSentimentMax(8.14)
        request.ContentSentimentMin(8.14)
        request.ContentSentimentMax(8.14)
        request.IptcTags()
        request.NotIptcTags()
        request.SourceName()
        request.IabTags()
        request.NotIabTags()
        request.ExcludeDuplicates(true)
        
        resp, httpRes, err := request.Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)
        */
    })

    t.Run("Test SearchApiService Post", func(t *testing.T) {
        /* TODO: ENG-1367 Fix parameter values for Go SDK generated tests
        from := *newscatcherapi.NewFrom()
        to := *newscatcherapi.NewTo()
        
        searchRequest := *newscatcherapi.NewSearchRequest(
            "null",
        )
        searchRequest.SetSearchIn("title_content")
        searchRequest.SetPredefinedSources(null)
        searchRequest.SetSources(null)
        searchRequest.SetNotSources(null)
        searchRequest.SetLang(null)
        searchRequest.SetNotLang(null)
        searchRequest.SetCountries(null)
        searchRequest.SetNotCountries(null)
        searchRequest.SetNotAuthorName(null)
        searchRequest.SetFrom(from)
        searchRequest.SetTo(to)
        searchRequest.SetPublishedDatePrecision("null")
        searchRequest.SetByParseDate(false)
        searchRequest.SetSortBy("relevancy")
        searchRequest.SetRankedOnly("null")
        searchRequest.SetFromRank(null)
        searchRequest.SetToRank(null)
        searchRequest.SetIsHeadline(null)
        searchRequest.SetIsPaidContent(null)
        searchRequest.SetParentUrl(null)
        searchRequest.SetAllLinks(null)
        searchRequest.SetAllDomainLinks(null)
        searchRequest.SetWordCountMin(null)
        searchRequest.SetWordCountMax(null)
        searchRequest.SetPage(1)
        searchRequest.SetPageSize(100)
        searchRequest.SetClusteringVariable("null")
        searchRequest.SetClusteringEnabled(null)
        searchRequest.SetClusteringThreshold(null)
        searchRequest.SetIncludeNlpData(null)
        searchRequest.SetHasNlp(null)
        searchRequest.SetTheme("null")
        searchRequest.SetNotTheme("null")
        searchRequest.SetORGEntityName("null")
        searchRequest.SetPEREntityName("null")
        searchRequest.SetLOCEntityName("null")
        searchRequest.SetMISCEntityName("null")
        searchRequest.SetTitleSentimentMin(null)
        searchRequest.SetTitleSentimentMax(null)
        searchRequest.SetContentSentimentMin(null)
        searchRequest.SetContentSentimentMax(null)
        searchRequest.SetIptcTags(null)
        searchRequest.SetNotIptcTags(null)
        searchRequest.SetSourceName(null)
        searchRequest.SetIabTags(null)
        searchRequest.SetNotIabTags(null)
        searchRequest.SetExcludeDuplicates(null)
        
        request := client.SearchApi.Post(
            searchRequest,
        )
        
        resp, httpRes, err := request.Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)
        */
    })

}