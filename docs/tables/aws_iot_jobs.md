# Table: aws_iot_jobs

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| comment | string | X | √ |  | 
| completed_at | timestamp | X | √ |  | 
| reason_code | string | X | √ |  | 
| timeout_config | json | X | √ |  | 
| job_template_arn | string | X | √ |  | 
| force_canceled | bool | X | √ |  | 
| job_arn | string | X | √ |  | 
| job_process_details | json | X | √ |  | 
| scheduling_config | json | X | √ |  | 
| target_selection | string | X | √ |  | 
| status | string | X | √ |  | 
| is_concurrent | bool | X | √ |  | 
| job_executions_rollout_config | json | X | √ |  | 
| job_id | string | X | √ |  | 
| last_updated_at | timestamp | X | √ |  | 
| region | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| job_executions_retry_config | json | X | √ |  | 
| tags | json | X | √ |  | 
| created_at | timestamp | X | √ |  | 
| document_parameters | json | X | √ |  | 
| presigned_url_config | json | X | √ |  | 
| targets | string_array | X | √ |  | 
| account_id | string | X | √ |  | 
| abort_config | json | X | √ |  | 
| description | string | X | √ |  | 
| namespace_id | string | X | √ |  | 
| arn | string | √ | √ |  | 


