# Table: aws_glue_registry_schemas

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| latest_schema_version | big_int | X | √ |  | 
| schema_status | string | X | √ |  | 
| data_format | string | X | √ |  | 
| registry_name | string | X | √ |  | 
| schema_checkpoint | big_int | X | √ |  | 
| region | string | X | √ |  | 
| arn | string | √ | √ |  | 
| tags | json | X | √ |  | 
| compatibility | string | X | √ |  | 
| registry_arn | string | X | √ |  | 
| schema_arn | string | X | √ |  | 
| result_metadata | json | X | √ |  | 
| account_id | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| aws_glue_registries_selefra_id | string | X | X | fk to aws_glue_registries.selefra_id | 
| created_time | string | X | √ |  | 
| description | string | X | √ |  | 
| next_schema_version | big_int | X | √ |  | 
| schema_name | string | X | √ |  | 
| updated_time | string | X | √ |  | 


