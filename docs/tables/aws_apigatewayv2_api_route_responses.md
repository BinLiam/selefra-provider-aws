# Table: aws_apigatewayv2_api_route_responses

## Primary Keys 

```
account_id, arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| aws_apigatewayv2_api_routes_selefra_id | string | X | X | fk to aws_apigatewayv2_api_routes.selefra_id | 
| route_response_key | string | X | √ |  | 
| model_selection_expression | string | X | √ |  | 
| response_models | json | X | √ |  | 
| response_parameters | json | X | √ |  | 
| route_response_id | string | X | √ |  | 
| route_id | string | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| api_route_arn | string | X | √ |  | 
| arn | string | X | √ |  | 


