# Table: aws_dynamodb_table_continuous_backups

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| region | string | X | √ |  | 
| table_arn | string | X | √ |  | 
| aws_dynamodb_tables_selefra_id | string | X | X | fk to aws_dynamodb_tables.selefra_id | 
| selefra_id | string | √ | √ | random id | 
| continuous_backups_status | string | X | √ |  | 
| point_in_time_recovery_description | json | X | √ |  | 
| account_id | string | X | √ |  | 


