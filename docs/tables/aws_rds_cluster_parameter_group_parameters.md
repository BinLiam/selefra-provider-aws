# Table: aws_rds_cluster_parameter_group_parameters

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| apply_method | string | X | √ |  | 
| data_type | string | X | √ |  | 
| description | string | X | √ |  | 
| parameter_value | string | X | √ |  | 
| cluster_parameter_group_arn | string | X | √ |  | 
| allowed_values | string | X | √ |  | 
| apply_type | string | X | √ |  | 
| minimum_engine_version | string | X | √ |  | 
| account_id | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| is_modifiable | bool | X | √ |  | 
| source | string | X | √ |  | 
| supported_engine_modes | string_array | X | √ |  | 
| aws_rds_cluster_parameter_groups_selefra_id | string | X | X | fk to aws_rds_cluster_parameter_groups.selefra_id | 
| parameter_name | string | X | √ |  | 
| region | string | X | √ |  | 


