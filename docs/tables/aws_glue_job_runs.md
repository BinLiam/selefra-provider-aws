# Table: aws_glue_job_runs

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| arguments | json | X | √ |  | 
| execution_class | string | X | √ |  | 
| job_run_state | string | X | √ |  | 
| last_modified_on | timestamp | X | √ |  | 
| number_of_workers | big_int | X | √ |  | 
| predecessor_runs | json | X | √ |  | 
| started_on | timestamp | X | √ |  | 
| region | string | X | √ |  | 
| id | string | X | √ |  | 
| trigger_name | string | X | √ |  | 
| aws_glue_jobs_selefra_id | string | X | X | fk to aws_glue_jobs.selefra_id | 
| selefra_id | string | √ | √ | random id | 
| completed_on | timestamp | X | √ |  | 
| previous_run_id | string | X | √ |  | 
| execution_time | big_int | X | √ |  | 
| security_configuration | string | X | √ |  | 
| attempt | big_int | X | √ |  | 
| glue_version | string | X | √ |  | 
| max_capacity | float | X | √ |  | 
| account_id | string | X | √ |  | 
| dpu_seconds | float | X | √ |  | 
| notification_property | json | X | √ |  | 
| timeout | big_int | X | √ |  | 
| job_arn | string | X | √ |  | 
| allocated_capacity | big_int | X | √ |  | 
| error_message | string | X | √ |  | 
| job_name | string | X | √ |  | 
| log_group_name | string | X | √ |  | 
| worker_type | string | X | √ |  | 


