# Table: aws_apigateway_rest_api_documentation_versions

## Primary Keys 

```
account_id, arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| rest_api_arn | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| version | string | X | √ |  | 
| description | string | X | √ |  | 
| region | string | X | √ |  | 
| arn | string | X | √ |  | 
| aws_apigateway_rest_apis_selefra_id | string | X | X | fk to aws_apigateway_rest_apis.selefra_id | 
| created_date | timestamp | X | √ |  | 


