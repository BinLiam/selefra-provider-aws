# Table: aws_apigatewayv2_domain_name_rest_api_mappings

## Primary Keys 

```
account_id, arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| api_mapping_key | string | X | √ |  | 
| region | string | X | √ |  | 
| arn | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| aws_apigatewayv2_domain_names_selefra_id | string | X | X | fk to aws_apigatewayv2_domain_names.selefra_id | 
| api_mapping_id | string | X | √ |  | 
| stage | string | X | √ |  | 
| account_id | string | X | √ |  | 
| domain_name_arn | string | X | √ |  | 
| api_id | string | X | √ |  | 


