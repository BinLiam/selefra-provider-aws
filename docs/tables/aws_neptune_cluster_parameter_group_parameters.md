# Table: aws_neptune_cluster_parameter_group_parameters

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| is_modifiable | bool | X | √ |  | 
| parameter_value | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| aws_neptune_cluster_parameter_groups_selefra_id | string | X | X | fk to aws_neptune_cluster_parameter_groups.selefra_id | 
| allowed_values | string | X | √ |  | 
| apply_method | string | X | √ |  | 
| apply_type | string | X | √ |  | 
| parameter_name | string | X | √ |  | 
| region | string | X | √ |  | 
| source | string | X | √ |  | 
| account_id | string | X | √ |  | 
| cluster_parameter_group_arn | string | X | √ |  | 
| data_type | string | X | √ |  | 
| description | string | X | √ |  | 
| minimum_engine_version | string | X | √ |  | 


