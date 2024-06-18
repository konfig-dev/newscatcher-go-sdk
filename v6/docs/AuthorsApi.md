# APIClient.AuthorsApi

All URIs are relative to *https://v3-api.newscatcherapi.com*

Method | Path | Description
------------- | ------------- | -------------
[**Get**](AuthorsApi.md#Get) | **Get** /api/authors | [Get] Search By Author Request
[**Post**](AuthorsApi.md#Post) | **Post** /api/authors | [Post] Search By Author Request



## Get

[Get] Search By Author Request



### Example

```go
package main

import (
    "fmt"
    "os"
    newscatcherapi "github.com/konfig-dev/newscatcher-go-sdk"
)

func main() {
    configuration := newscatcherapi.NewConfiguration()
    configuration.SetApiKey("X_API_TOKEN")
    client := newscatcherapi.NewAPIClient(configuration)

    request := client.AuthorsApi.Get(
        "authorName_example",
    )
    request.NotAuthorName("notAuthorName_example")
    request.Sources()
    request.PredefinedSources()
    request.NotSources()
    request.Lang()
    request.NotLang()
    request.Countries()
    request.NotCountries()
    request.From("from_example")
    request.To("to_example")
    request.PublishedDatePrecision("publishedDatePrecision_example")
    request.ByParseDate(false)
    request.SortBy(""relevancy"")
    request.RankedOnly("rankedOnly_example")
    request.FromRank(56)
    request.ToRank(56)
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
    request.IncludeNlpData(true)
    request.HasNlp(true)
    request.Theme("theme_example")
    request.NotTheme("notTheme_example")
    request.NerName("nerName_example")
    request.TitleSentimentMin(8.14)
    request.TitleSentimentMax(8.14)
    request.ContentSentimentMin(8.14)
    request.ContentSentimentMax(8.14)
    request.IptcTags()
    request.NotIptcTags()
    request.IabTags()
    request.NotIabTags()
    
    resp, httpRes, err := request.Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorsApi.Get``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", httpRes)
    }
    // response from `Get`: FSearchResponse
    fmt.Fprintf(os.Stdout, "Response from `AuthorsApi.Get`: %v\n", resp)
    fmt.Fprintf(os.Stdout, "Response from `FSearchResponse.Get.Status`: %v\n", *resp.Status)
    fmt.Fprintf(os.Stdout, "Response from `FSearchResponse.Get.TotalHits`: %v\n", *resp.TotalHits)
    fmt.Fprintf(os.Stdout, "Response from `FSearchResponse.Get.Page`: %v\n", *resp.Page)
    fmt.Fprintf(os.Stdout, "Response from `FSearchResponse.Get.TotalPages`: %v\n", *resp.TotalPages)
    fmt.Fprintf(os.Stdout, "Response from `FSearchResponse.Get.PageSize`: %v\n", *resp.PageSize)
    fmt.Fprintf(os.Stdout, "Response from `FSearchResponse.Get.Articles`: %v\n", *resp.Articles)
    fmt.Fprintf(os.Stdout, "Response from `FSearchResponse.Get.UserInput`: %v\n", resp.UserInput)
}
```

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Post

[Post] Search By Author Request



### Example

```go
package main

import (
    "fmt"
    "os"
    newscatcherapi "github.com/konfig-dev/newscatcher-go-sdk"
)

func main() {
    configuration := newscatcherapi.NewConfiguration()
    configuration.SetApiKey("X_API_TOKEN")
    client := newscatcherapi.NewAPIClient(configuration)

    
    authorSearchRequest := *newscatcherapi.NewAuthorSearchRequest(
        "null",
    )
    authorSearchRequest.SetNotAuthorName("null")
    authorSearchRequest.SetSources(null)
    authorSearchRequest.SetPredefinedSources(null)
    authorSearchRequest.SetNotSources(null)
    authorSearchRequest.SetLang(null)
    authorSearchRequest.SetNotLang(null)
    authorSearchRequest.SetCountries(null)
    authorSearchRequest.SetNotCountries(null)
    authorSearchRequest.SetFrom("null")
    authorSearchRequest.SetTo("null")
    authorSearchRequest.SetPublishedDatePrecision("null")
    authorSearchRequest.SetByParseDate(false)
    authorSearchRequest.SetSortBy("relevancy")
    authorSearchRequest.SetRankedOnly("null")
    authorSearchRequest.SetFromRank(null)
    authorSearchRequest.SetToRank(null)
    authorSearchRequest.SetIsHeadline(null)
    authorSearchRequest.SetIsOpinion(null)
    authorSearchRequest.SetIsPaidContent(null)
    authorSearchRequest.SetParentUrl(null)
    authorSearchRequest.SetAllLinks(null)
    authorSearchRequest.SetAllDomainLinks(null)
    authorSearchRequest.SetWordCountMin(null)
    authorSearchRequest.SetWordCountMax(null)
    authorSearchRequest.SetPage(1)
    authorSearchRequest.SetPageSize(100)
    authorSearchRequest.SetIncludeNlpData(null)
    authorSearchRequest.SetHasNlp(null)
    authorSearchRequest.SetTheme("null")
    authorSearchRequest.SetNotTheme("null")
    authorSearchRequest.SetNerName("null")
    authorSearchRequest.SetTitleSentimentMin(null)
    authorSearchRequest.SetTitleSentimentMax(null)
    authorSearchRequest.SetContentSentimentMin(null)
    authorSearchRequest.SetContentSentimentMax(null)
    authorSearchRequest.SetIptcTags(null)
    authorSearchRequest.SetNotIptcTags(null)
    authorSearchRequest.SetIabTags(null)
    authorSearchRequest.SetNotIabTags(null)
    
    request := client.AuthorsApi.Post(
        authorSearchRequest,
    )
    
    resp, httpRes, err := request.Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorsApi.Post``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", httpRes)
    }
    // response from `Post`: FSearchResponse1
    fmt.Fprintf(os.Stdout, "Response from `AuthorsApi.Post`: %v\n", resp)
    fmt.Fprintf(os.Stdout, "Response from `FSearchResponse1.Post.Status`: %v\n", *resp.Status)
    fmt.Fprintf(os.Stdout, "Response from `FSearchResponse1.Post.TotalHits`: %v\n", *resp.TotalHits)
    fmt.Fprintf(os.Stdout, "Response from `FSearchResponse1.Post.Page`: %v\n", *resp.Page)
    fmt.Fprintf(os.Stdout, "Response from `FSearchResponse1.Post.TotalPages`: %v\n", *resp.TotalPages)
    fmt.Fprintf(os.Stdout, "Response from `FSearchResponse1.Post.PageSize`: %v\n", *resp.PageSize)
    fmt.Fprintf(os.Stdout, "Response from `FSearchResponse1.Post.Articles`: %v\n", *resp.Articles)
    fmt.Fprintf(os.Stdout, "Response from `FSearchResponse1.Post.UserInput`: %v\n", resp.UserInput)
}
```

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

