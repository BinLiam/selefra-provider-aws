# Table: aws_apigatewayv2_api_routes

## Primary Keys 

```
account_id, arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| arn | string | X | √ |  | 
| aws_apigatewayv2_apis_selefra_id | string | X | X | fk to aws_apigatewayv2_apis.selefra_id | 
| authorization_scopes | string_array | X | √ |  | 
| request_models | json | X | √ |  | 
| request_parameters | json | X | √ |  | 
| route_id | string | X | √ |  | 
| route_response_selection_expression | string | X | √ |  | 
| api_id | string | X | √ |  | 
| api_key_required | bool | X | √ |  | 
| route_key | string | X | √ |  | 
| authorization_type | string | X | √ |  | 
| authorizer_id | string | X | √ |  | 
| operation_name | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| api_gateway_managed | bool | X | √ |  | 
| model_selection_expression | string | X | √ |  | 
| target | string | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| api_arn | string | X | √ |  | 


