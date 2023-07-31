# Table: aws_athena_work_group_prepared_statements

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| statement_name | string | X | √ |  | 
| work_group_name | string | X | √ |  | 
| account_id | string | X | √ |  | 
| description | string | X | √ |  | 
| last_modified_time | timestamp | X | √ |  | 
| query_statement | string | X | √ |  | 
| region | string | X | √ |  | 
| work_group_arn | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| aws_athena_work_groups_selefra_id | string | X | X | fk to aws_athena_work_groups.selefra_id | 


