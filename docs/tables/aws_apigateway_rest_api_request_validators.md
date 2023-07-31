# Table: aws_apigateway_rest_api_request_validators

## Primary Keys 

```
account_id, arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| id | string | X | √ |  | 
| name | string | X | √ |  | 
| validate_request_body | bool | X | √ |  | 
| arn | string | X | √ |  | 
| aws_apigateway_rest_apis_selefra_id | string | X | X | fk to aws_apigateway_rest_apis.selefra_id | 
| validate_request_parameters | bool | X | √ |  | 
| region | string | X | √ |  | 
| rest_api_arn | string | X | √ |  | 


