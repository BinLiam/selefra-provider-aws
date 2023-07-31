# Table: aws_glue_ml_transform_task_runs

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| last_modified_on | timestamp | X | √ |  | 
| started_on | timestamp | X | √ |  | 
| status | string | X | √ |  | 
| ml_transform_arn | string | X | √ |  | 
| error_string | string | X | √ |  | 
| aws_glue_ml_transforms_selefra_id | string | X | X | fk to aws_glue_ml_transforms.selefra_id | 
| execution_time | big_int | X | √ |  | 
| properties | json | X | √ |  | 
| task_run_id | string | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| completed_on | timestamp | X | √ |  | 
| log_group_name | string | X | √ |  | 
| transform_id | string | X | √ |  | 


