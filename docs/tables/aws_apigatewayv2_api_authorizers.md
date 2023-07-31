# Table: aws_apigatewayv2_api_authorizers

## Primary Keys 

```
account_id, arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| authorizer_uri | string | X | √ |  | 
| identity_validation_expression | string | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| authorizer_credentials_arn | string | X | √ |  | 
| authorizer_payload_format_version | string | X | √ |  | 
| authorizer_type | string | X | √ |  | 
| aws_apigatewayv2_apis_selefra_id | string | X | X | fk to aws_apigatewayv2_apis.selefra_id | 
| jwt_configuration | json | X | √ |  | 
| api_arn | string | X | √ |  | 
| name | string | X | √ |  | 
| authorizer_id | string | X | √ |  | 
| enable_simple_responses | bool | X | √ |  | 
| arn | string | X | √ |  | 
| authorizer_result_ttl_in_seconds | big_int | X | √ |  | 
| identity_source | string_array | X | √ |  | 
| api_id | string | X | √ |  | 


