# Table: aws_apigateway_domain_name_base_path_mappings

## Primary Keys 

```
account_id, arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| rest_api_id | string | X | √ |  | 
| account_id | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| aws_apigateway_domain_names_selefra_id | string | X | X | fk to aws_apigateway_domain_names.selefra_id | 
| base_path | string | X | √ |  | 
| stage | string | X | √ |  | 
| region | string | X | √ |  | 
| domain_name_arn | string | X | √ |  | 
| arn | string | X | √ |  | 


