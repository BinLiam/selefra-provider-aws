# Table: aws_rds_cluster_parameters

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| apply_method | string | X | √ |  | 
| apply_type | string | X | √ |  | 
| is_modifiable | bool | X | √ |  | 
| parameter_name | string | X | √ |  | 
| data_type | string | X | √ |  | 
| minimum_engine_version | string | X | √ |  | 
| parameter_value | string | X | √ |  | 
| supported_engine_modes | string_array | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| description | string | X | √ |  | 
| source | string | X | √ |  | 
| aws_rds_engine_versions_selefra_id | string | X | X | fk to aws_rds_engine_versions.selefra_id | 
| allowed_values | string | X | √ |  | 


