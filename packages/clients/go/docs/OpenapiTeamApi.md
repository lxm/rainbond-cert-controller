# \OpenapiTeamApi

All URIs are relative to *https://cloud.goodrain.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OpenapiV1TeamsCertificatesCreate**](OpenapiTeamApi.md#OpenapiV1TeamsCertificatesCreate) | **Post** /openapi/v1/teams/{team_id}/certificates | 
[**OpenapiV1TeamsCertificatesDelete**](OpenapiTeamApi.md#OpenapiV1TeamsCertificatesDelete) | **Delete** /openapi/v1/teams/{team_id}/certificates/{certificate_id} | 
[**OpenapiV1TeamsCertificatesList**](OpenapiTeamApi.md#OpenapiV1TeamsCertificatesList) | **Get** /openapi/v1/teams/{team_id}/certificates | 
[**OpenapiV1TeamsCertificatesRead**](OpenapiTeamApi.md#OpenapiV1TeamsCertificatesRead) | **Get** /openapi/v1/teams/{team_id}/certificates/{certificate_id} | 
[**OpenapiV1TeamsCertificatesUpdate**](OpenapiTeamApi.md#OpenapiV1TeamsCertificatesUpdate) | **Put** /openapi/v1/teams/{team_id}/certificates/{certificate_id} | 
[**OpenapiV1TeamsCreate**](OpenapiTeamApi.md#OpenapiV1TeamsCreate) | **Post** /openapi/v1/teams | 
[**OpenapiV1TeamsList**](OpenapiTeamApi.md#OpenapiV1TeamsList) | **Get** /openapi/v1/teams | 


# **OpenapiV1TeamsCertificatesCreate**
> TeamCertificatesR OpenapiV1TeamsCertificatesCreate(ctx, teamId, data)


添加证书

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **teamId** | **string**|  | 
  **data** | [**TeamCertificatesC**](TeamCertificatesC.md)|  | 

### Return type

[**TeamCertificatesR**](TeamCertificatesR.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OpenapiV1TeamsCertificatesDelete**
> OpenapiV1TeamsCertificatesDelete(ctx, certificateId, teamId)


删除证书

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **certificateId** | **string**|  | 
  **teamId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OpenapiV1TeamsCertificatesList**
> TeamCertificatesL OpenapiV1TeamsCertificatesList(ctx, teamId, optional)


获取团队下证书列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **teamId** | **string**|  | 
 **optional** | ***OpenapiTeamApiOpenapiV1TeamsCertificatesListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OpenapiTeamApiOpenapiV1TeamsCertificatesListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Float32**| 页码 | 
 **pageSize** | **optional.Float32**| 每页数量 | 

### Return type

[**TeamCertificatesL**](TeamCertificatesL.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OpenapiV1TeamsCertificatesRead**
> TeamCertificatesR OpenapiV1TeamsCertificatesRead(ctx, certificateId, teamId)


获取团队下证书列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **certificateId** | **string**|  | 
  **teamId** | **string**|  | 

### Return type

[**TeamCertificatesR**](TeamCertificatesR.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OpenapiV1TeamsCertificatesUpdate**
> TeamCertificatesR OpenapiV1TeamsCertificatesUpdate(ctx, certificateId, teamId, data)


更新证书

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **certificateId** | **string**|  | 
  **teamId** | **string**|  | 
  **data** | [**TeamCertificatesC**](TeamCertificatesC.md)|  | 

### Return type

[**TeamCertificatesR**](TeamCertificatesR.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OpenapiV1TeamsCreate**
> TeamBaseInfo OpenapiV1TeamsCreate(ctx, data)


add team

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**CreateTeamReq**](CreateTeamReq.md)|  | 

### Return type

[**TeamBaseInfo**](TeamBaseInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OpenapiV1TeamsList**
> ListTeamResp OpenapiV1TeamsList(ctx, optional)


获取用户所在团队列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OpenapiTeamApiOpenapiV1TeamsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OpenapiTeamApiOpenapiV1TeamsListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **optional.String**| 团队名称搜索 | 
 **page** | **optional.String**| 页码 | 
 **pageSize** | **optional.String**| 每页数量 | 

### Return type

[**ListTeamResp**](ListTeamResp.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

