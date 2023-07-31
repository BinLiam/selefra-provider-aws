# Table: aws_lambda_layer_versions

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| version | big_int | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| arn | string | √ | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| compatible_runtimes | string_array | X | √ |  | 
| layer_version_arn | string | X | √ |  | 
| description | string | X | √ |  | 
| license_info | string | X | √ |  | 
| layer_arn | string | X | √ |  | 
| aws_lambda_layers_selefra_id | string | X | X | fk to aws_lambda_layers.selefra_id | 
| compatible_architectures | string_array | X | √ |  | 
| created_date | string | X | √ |  | 


