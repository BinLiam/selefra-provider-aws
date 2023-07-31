# Table: aws_mwaa_environments

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| logging_configuration | json | X | √ |  | 
| tags | json | X | √ |  | 
| created_at | timestamp | X | √ |  | 
| min_workers | big_int | X | √ |  | 
| name | string | X | √ |  | 
| webserver_access_mode | string | X | √ |  | 
| weekly_maintenance_window_start | string | X | √ |  | 
| arn | string | √ | √ |  | 
| last_update | json | X | √ |  | 
| requirements_s3_path | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| kms_key | string | X | √ |  | 
| requirements_s3_object_version | string | X | √ |  | 
| schedulers | big_int | X | √ |  | 
| dag_s3_path | string | X | √ |  | 
| airflow_configuration_options | json | X | √ |  | 
| airflow_version | string | X | √ |  | 
| environment_class | string | X | √ |  | 
| execution_role_arn | string | X | √ |  | 
| plugins_s3_object_version | string | X | √ |  | 
| service_role_arn | string | X | √ |  | 
| status | string | X | √ |  | 
| webserver_url | string | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| max_workers | big_int | X | √ |  | 
| plugins_s3_path | string | X | √ |  | 
| network_configuration | json | X | √ |  | 
| source_bucket_arn | string | X | √ |  | 


