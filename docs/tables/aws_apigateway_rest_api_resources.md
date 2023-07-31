# Table: aws_apigateway_rest_api_resources

## Primary Keys 

```
account_id, arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| aws_apigateway_rest_apis_selefra_id | string | X | X | fk to aws_apigateway_rest_apis.selefra_id | 
| path | string | X | √ |  | 
| resource_methods | json | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| rest_api_arn | string | X | √ |  | 
| arn | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| id | string | X | √ |  | 
| parent_id | string | X | √ |  | 
| path_part | string | X | √ |  | 


