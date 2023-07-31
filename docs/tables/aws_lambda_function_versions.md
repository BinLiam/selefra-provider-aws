# Table: aws_lambda_function_versions

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| master_arn | string | X | √ |  | 
| runtime_version_config | json | X | √ |  | 
| timeout | big_int | X | √ |  | 
| architectures | string_array | X | √ |  | 
| function_arn | string | X | √ |  | 
| function_name | string | X | √ |  | 
| environment | json | X | √ |  | 
| version | string | X | √ |  | 
| image_config_response | json | X | √ |  | 
| last_update_status_reason_code | string | X | √ |  | 
| memory_size | big_int | X | √ |  | 
| snap_start | json | X | √ |  | 
| aws_lambda_functions_selefra_id | string | X | X | fk to aws_lambda_functions.selefra_id | 
| kms_key_arn | string | X | √ |  | 
| package_type | string | X | √ |  | 
| signing_job_arn | string | X | √ |  | 
| role | string | X | √ |  | 
| state | string | X | √ |  | 
| state_reason | string | X | √ |  | 
| state_reason_code | string | X | √ |  | 
| tracing_config | json | X | √ |  | 
| dead_letter_config | json | X | √ |  | 
| ephemeral_storage | json | X | √ |  | 
| handler | string | X | √ |  | 
| account_id | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| description | string | X | √ |  | 
| last_update_status | string | X | √ |  | 
| layers | json | X | √ |  | 
| last_modified | string | X | √ |  | 
| revision_id | string | X | √ |  | 
| signing_profile_version_arn | string | X | √ |  | 
| vpc_config | json | X | √ |  | 
| code_sha256 | string | X | √ |  | 
| code_size | big_int | X | √ |  | 
| file_system_configs | json | X | √ |  | 
| last_update_status_reason | string | X | √ |  | 
| runtime | string | X | √ |  | 
| region | string | X | √ |  | 


