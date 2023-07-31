# Table: aws_apigatewayv2_api_models

## Primary Keys 

```
account_id, arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| model_id | string | X | √ |  | 
| schema | string | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| arn | string | X | √ |  | 
| model_template | string | X | √ |  | 
| name | string | X | √ |  | 
| description | string | X | √ |  | 
| api_arn | string | X | √ |  | 
| api_id | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| aws_apigatewayv2_apis_selefra_id | string | X | X | fk to aws_apigatewayv2_apis.selefra_id | 
| content_type | string | X | √ |  | 


