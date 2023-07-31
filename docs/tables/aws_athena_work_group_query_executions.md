# Table: aws_athena_work_group_query_executions

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| query | string | X | √ |  | 
| query_execution_context | json | X | √ |  | 
| query_execution_id | string | X | √ |  | 
| statistics | json | X | √ |  | 
| region | string | X | √ |  | 
| work_group_arn | string | X | √ |  | 
| engine_version | json | X | √ |  | 
| statement_type | string | X | √ |  | 
| work_group | string | X | √ |  | 
| account_id | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| execution_parameters | string_array | X | √ |  | 
| status | json | X | √ |  | 
| aws_athena_work_groups_selefra_id | string | X | X | fk to aws_athena_work_groups.selefra_id | 
| result_configuration | json | X | √ |  | 
| result_reuse_configuration | json | X | √ |  | 


