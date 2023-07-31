# Table: aws_dynamodb_table_replica_auto_scalings

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| table_arn | string | X | √ |  | 
| aws_dynamodb_tables_selefra_id | string | X | X | fk to aws_dynamodb_tables.selefra_id | 
| global_secondary_indexes | json | X | √ |  | 
| replica_status | string | X | √ |  | 
| replica_provisioned_write_capacity_auto_scaling_settings | json | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| region_name | string | X | √ |  | 
| replica_provisioned_read_capacity_auto_scaling_settings | json | X | √ |  | 


