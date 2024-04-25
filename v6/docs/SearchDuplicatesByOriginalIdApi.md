# APIClient.SearchDuplicatesByOriginalIdApi

All URIs are relative to *https://v3-api.newscatcherapi.com*

Method | Path | Description
------------- | ------------- | -------------
[**Get**](SearchDuplicatesByOriginalIdApi.md#Get) | **Get** /api/search_duplicates_by_original_id | [Get] Search Duplicate Articles For Articles Request
[**Post**](SearchDuplicatesByOriginalIdApi.md#Post) | **Post** /api/search_duplicates_by_original_id | [Post] Search Duplicate Articles For Articles Request



## Get

[Get] Search Duplicate Articles For Articles Request



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

    request := client.SearchDuplicatesByOriginalIdApi.Get(
        "originalArticleId_example",
    )
    request.Page(1)
    request.PageSize(100)
    
    resp, httpRes, err := request.Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchDuplicatesByOriginalIdApi.Get``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", httpRes)
    }
    // response from `Get`: SearchduplicatesbyoriginalidGetResponse
    fmt.Fprintf(os.Stdout, "Response from `SearchDuplicatesByOriginalIdApi.Get`: %v\n", resp)
    fmt.Fprintf(os.Stdout, "Response from `SearchduplicatesbyoriginalidGetResponse.Get.Status`: %v\n", *resp.Status)
    fmt.Fprintf(os.Stdout, "Response from `SearchduplicatesbyoriginalidGetResponse.Get.TotalHits`: %v\n", resp.TotalHits)
    fmt.Fprintf(os.Stdout, "Response from `SearchduplicatesbyoriginalidGetResponse.Get.Page`: %v\n", resp.Page)
    fmt.Fprintf(os.Stdout, "Response from `SearchduplicatesbyoriginalidGetResponse.Get.TotalPages`: %v\n", resp.TotalPages)
    fmt.Fprintf(os.Stdout, "Response from `SearchduplicatesbyoriginalidGetResponse.Get.PageSize`: %v\n", resp.PageSize)
    fmt.Fprintf(os.Stdout, "Response from `SearchduplicatesbyoriginalidGetResponse.Get.Articles`: %v\n", resp.Articles)
    fmt.Fprintf(os.Stdout, "Response from `SearchduplicatesbyoriginalidGetResponse.Get.UserInput`: %v\n", resp.UserInput)
    fmt.Fprintf(os.Stdout, "Response from `SearchduplicatesbyoriginalidGetResponse.Get.ClustersCount`: %v\n", resp.ClustersCount)
    fmt.Fprintf(os.Stdout, "Response from `SearchduplicatesbyoriginalidGetResponse.Get.Clusters`: %v\n", resp.Clusters)
}
```

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Post

[Post] Search Duplicate Articles For Articles Request



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

    
    duplicatesByOriginalIdRequest := *newscatcherapi.NewDuplicatesByOriginalIdRequest(
        "null",
    )
    duplicatesByOriginalIdRequest.SetPage(1)
    duplicatesByOriginalIdRequest.SetPageSize(100)
    
    request := client.SearchDuplicatesByOriginalIdApi.Post(
        duplicatesByOriginalIdRequest,
    )
    
    resp, httpRes, err := request.Execute()

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchDuplicatesByOriginalIdApi.Post``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", httpRes)
    }
    // response from `Post`: SearchduplicatesbyoriginalidPostResponse
    fmt.Fprintf(os.Stdout, "Response from `SearchDuplicatesByOriginalIdApi.Post`: %v\n", resp)
    fmt.Fprintf(os.Stdout, "Response from `SearchduplicatesbyoriginalidPostResponse.Post.Status`: %v\n", *resp.Status)
    fmt.Fprintf(os.Stdout, "Response from `SearchduplicatesbyoriginalidPostResponse.Post.TotalHits`: %v\n", resp.TotalHits)
    fmt.Fprintf(os.Stdout, "Response from `SearchduplicatesbyoriginalidPostResponse.Post.Page`: %v\n", resp.Page)
    fmt.Fprintf(os.Stdout, "Response from `SearchduplicatesbyoriginalidPostResponse.Post.TotalPages`: %v\n", resp.TotalPages)
    fmt.Fprintf(os.Stdout, "Response from `SearchduplicatesbyoriginalidPostResponse.Post.PageSize`: %v\n", resp.PageSize)
    fmt.Fprintf(os.Stdout, "Response from `SearchduplicatesbyoriginalidPostResponse.Post.Articles`: %v\n", resp.Articles)
    fmt.Fprintf(os.Stdout, "Response from `SearchduplicatesbyoriginalidPostResponse.Post.UserInput`: %v\n", resp.UserInput)
    fmt.Fprintf(os.Stdout, "Response from `SearchduplicatesbyoriginalidPostResponse.Post.ClustersCount`: %v\n", resp.ClustersCount)
    fmt.Fprintf(os.Stdout, "Response from `SearchduplicatesbyoriginalidPostResponse.Post.Clusters`: %v\n", resp.Clusters)
}
```

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

