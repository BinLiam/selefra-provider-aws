# Table: aws_apigateway_rest_api_gateway_responses

## Primary Keys 

```
account_id, arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| rest_api_arn | string | X | √ |  | 
| arn | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| response_templates | json | X | √ |  | 
| response_type | string | X | √ |  | 
| status_code | string | X | √ |  | 
| aws_apigateway_rest_apis_selefra_id | string | X | X | fk to aws_apigateway_rest_apis.selefra_id | 
| default_response | bool | X | √ |  | 
| response_parameters | json | X | √ |  | 


