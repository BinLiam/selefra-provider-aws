# Table: aws_apigatewayv2_domain_names

## Primary Keys 

```
account_id, arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| api_mapping_selection_expression | string | X | √ |  | 
| tags | json | X | √ |  | 
| region | string | X | √ |  | 
| arn | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| domain_name | string | X | √ |  | 
| domain_name_configurations | json | X | √ |  | 
| mutual_tls_authentication | json | X | √ |  | 
| account_id | string | X | √ |  | 


