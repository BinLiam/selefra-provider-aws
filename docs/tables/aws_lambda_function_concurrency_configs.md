# Table: aws_lambda_function_concurrency_configs

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| allocated_provisioned_concurrent_executions | big_int | X | √ |  | 
| available_provisioned_concurrent_executions | big_int | X | √ |  | 
| last_modified | string | X | √ |  | 
| requested_provisioned_concurrent_executions | big_int | X | √ |  | 
| status_reason | string | X | √ |  | 
| aws_lambda_functions_selefra_id | string | X | X | fk to aws_lambda_functions.selefra_id | 
| function_arn | string | X | √ |  | 
| status | string | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 


