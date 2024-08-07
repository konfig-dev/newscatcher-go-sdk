/*
NewsCatcher-V3 Production API

Testing LatestHeadlinesApiService

*/

// Code generated by Konfig (https://konfigthis.com);

package newscatcherapi

import (
    "testing"
    // "github.com/stretchr/testify/assert"
    // "github.com/stretchr/testify/require"
    // newscatcherapi "github.com/konfig-dev/newscatcher-go-sdk"
)

func Test_newscatcherapi_LatestHeadlinesApiService(t *testing.T) {

    // configuration := newscatcherapi.NewConfiguration()
    // configuration.SetHost("http://127.0.0.1:4010")
    /* 
    configuration.SetApiKey("X_API_TOKEN")
    client := newscatcherapi.NewAPIClient(configuration)
    */

    t.Run("Test LatestHeadlinesApiService Get", func(t *testing.T) {
        /* TODO: ENG-1367 Fix parameter values for Go SDK generated tests
        request := client.LatestHeadlinesApi.Get(
        )
        request.When(""7d"")
        request.ByParseDate(false)
        request.SortBy(""relevancy"")
        request.Lang()
        request.NotLang()
        request.Countries()
        request.NotCountries()
        request.Sources()
        request.PredefinedSources()
        request.NotSources()
        request.NotAuthorName()
        request.RankedOnly(rankedOnly)
        request.IsHeadline(true)
        request.IsOpinion(true)
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
        request.IabTags()
        request.NotIabTags()
        
        resp, httpRes, err := request.Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)
        */
    })

    t.Run("Test LatestHeadlinesApiService Post", func(t *testing.T) {
        /* TODO: ENG-1367 Fix parameter values for Go SDK generated tests
        rankedOnly := *newscatcherapi.NewRankedOnly()
        
        latestHeadlinesRequest := *newscatcherapi.NewLatestHeadlinesRequest()
        latestHeadlinesRequest.SetWhen("7d")
        latestHeadlinesRequest.SetByParseDate(false)
        latestHeadlinesRequest.SetSortBy("relevancy")
        latestHeadlinesRequest.SetLang(null)
        latestHeadlinesRequest.SetNotLang(null)
        latestHeadlinesRequest.SetCountries(null)
        latestHeadlinesRequest.SetNotCountries(null)
        latestHeadlinesRequest.SetSources(null)
        latestHeadlinesRequest.SetPredefinedSources(null)
        latestHeadlinesRequest.SetNotSources(null)
        latestHeadlinesRequest.SetNotAuthorName(null)
        latestHeadlinesRequest.SetRankedOnly(rankedOnly)
        latestHeadlinesRequest.SetIsHeadline(null)
        latestHeadlinesRequest.SetIsOpinion(null)
        latestHeadlinesRequest.SetIsPaidContent(null)
        latestHeadlinesRequest.SetParentUrl(null)
        latestHeadlinesRequest.SetAllLinks(null)
        latestHeadlinesRequest.SetAllDomainLinks(null)
        latestHeadlinesRequest.SetWordCountMin(null)
        latestHeadlinesRequest.SetWordCountMax(null)
        latestHeadlinesRequest.SetPage(1)
        latestHeadlinesRequest.SetPageSize(100)
        latestHeadlinesRequest.SetClusteringVariable("null")
        latestHeadlinesRequest.SetClusteringEnabled(null)
        latestHeadlinesRequest.SetClusteringThreshold(null)
        latestHeadlinesRequest.SetIncludeNlpData(null)
        latestHeadlinesRequest.SetHasNlp(null)
        latestHeadlinesRequest.SetTheme("null")
        latestHeadlinesRequest.SetNotTheme("null")
        latestHeadlinesRequest.SetORGEntityName("null")
        latestHeadlinesRequest.SetPEREntityName("null")
        latestHeadlinesRequest.SetLOCEntityName("null")
        latestHeadlinesRequest.SetMISCEntityName("null")
        latestHeadlinesRequest.SetTitleSentimentMin(null)
        latestHeadlinesRequest.SetTitleSentimentMax(null)
        latestHeadlinesRequest.SetContentSentimentMin(null)
        latestHeadlinesRequest.SetContentSentimentMax(null)
        latestHeadlinesRequest.SetIptcTags(null)
        latestHeadlinesRequest.SetNotIptcTags(null)
        latestHeadlinesRequest.SetIabTags(null)
        latestHeadlinesRequest.SetNotIabTags(null)
        
        request := client.LatestHeadlinesApi.Post(
            latestHeadlinesRequest,
        )
        
        resp, httpRes, err := request.Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)
        */
    })

}