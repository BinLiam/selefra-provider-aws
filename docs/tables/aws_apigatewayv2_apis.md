# Table: aws_apigatewayv2_apis

## Primary Keys 

```
account_id, arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| description | string | X | √ |  | 
| disable_execute_api_endpoint | bool | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| api_endpoint | string | X | √ |  | 
| cors_configuration | json | X | √ |  | 
| warnings | string_array | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| api_gateway_managed | bool | X | √ |  | 
| disable_schema_validation | bool | X | √ |  | 
| import_info | string_array | X | √ |  | 
| tags | json | X | √ |  | 
| arn | string | X | √ |  | 
| id | string | X | √ |  | 
| protocol_type | string | X | √ |  | 
| created_date | timestamp | X | √ |  | 
| api_id | string | X | √ |  | 
| api_key_selection_expression | string | X | √ |  | 
| version | string | X | √ |  | 
| name | string | X | √ |  | 
| route_selection_expression | string | X | √ |  | 


