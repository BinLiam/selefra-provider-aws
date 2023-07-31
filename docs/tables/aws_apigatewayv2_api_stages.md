# Table: aws_apigatewayv2_api_stages

## Primary Keys 

```
account_id, arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| arn | string | X | √ |  | 
| api_gateway_managed | bool | X | √ |  | 
| auto_deploy | bool | X | √ |  | 
| client_certificate_id | string | X | √ |  | 
| deployment_id | string | X | √ |  | 
| stage_variables | json | X | √ |  | 
| api_arn | string | X | √ |  | 
| api_id | string | X | √ |  | 
| access_log_settings | json | X | √ |  | 
| created_date | timestamp | X | √ |  | 
| last_updated_date | timestamp | X | √ |  | 
| aws_apigatewayv2_apis_selefra_id | string | X | X | fk to aws_apigatewayv2_apis.selefra_id | 
| last_deployment_status_message | string | X | √ |  | 
| route_settings | json | X | √ |  | 
| account_id | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| stage_name | string | X | √ |  | 
| default_route_settings | json | X | √ |  | 
| description | string | X | √ |  | 
| tags | json | X | √ |  | 
| region | string | X | √ |  | 


