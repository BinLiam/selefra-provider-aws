# Table: aws_glue_jobs

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| last_modified_on | timestamp | X | √ |  | 
| number_of_workers | big_int | X | √ |  | 
| worker_type | string | X | √ |  | 
| allocated_capacity | big_int | X | √ |  | 
| glue_version | string | X | √ |  | 
| execution_class | string | X | √ |  | 
| timeout | big_int | X | √ |  | 
| tags | json | X | √ |  | 
| account_id | string | X | √ |  | 
| code_gen_configuration_nodes | json | X | √ |  | 
| connections | json | X | √ |  | 
| max_capacity | float | X | √ |  | 
| arn | string | √ | √ |  | 
| created_on | timestamp | X | √ |  | 
| execution_property | json | X | √ |  | 
| log_uri | string | X | √ |  | 
| security_configuration | string | X | √ |  | 
| command | json | X | √ |  | 
| description | string | X | √ |  | 
| name | string | X | √ |  | 
| non_overridable_arguments | json | X | √ |  | 
| role | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| source_control_details | json | X | √ |  | 
| default_arguments | json | X | √ |  | 
| max_retries | big_int | X | √ |  | 
| notification_property | json | X | √ |  | 
| region | string | X | √ |  | 


