# Table: aws_glue_registry_schema_versions

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| schema_definition | string | X | √ |  | 
| schema_version_id | string | X | √ |  | 
| status | string | X | √ |  | 
| metadata | json | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| aws_glue_registry_schemas_selefra_id | string | X | X | fk to aws_glue_registry_schemas.selefra_id | 
| created_time | string | X | √ |  | 
| data_format | string | X | √ |  | 
| region | string | X | √ |  | 
| registry_schema_arn | string | X | √ |  | 
| schema_arn | string | X | √ |  | 
| result_metadata | json | X | √ |  | 
| version_number | big_int | X | √ |  | 
| account_id | string | X | √ |  | 


