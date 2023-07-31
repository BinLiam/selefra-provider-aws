# Table: aws_docdb_cluster_parameters

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| apply_method | string | X | √ |  | 
| description | string | X | √ |  | 
| parameter_name | string | X | √ |  | 
| parameter_value | string | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| allowed_values | string | X | √ |  | 
| is_modifiable | bool | X | √ |  | 
| minimum_engine_version | string | X | √ |  | 
| apply_type | string | X | √ |  | 
| data_type | string | X | √ |  | 
| source | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| aws_docdb_engine_versions_selefra_id | string | X | X | fk to aws_docdb_engine_versions.selefra_id | 


