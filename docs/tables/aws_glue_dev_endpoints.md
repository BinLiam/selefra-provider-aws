# Table: aws_glue_dev_endpoints

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| security_configuration | string | X | √ |  | 
| status | string | X | √ |  | 
| created_timestamp | timestamp | X | √ |  | 
| extra_jars_s3_path | string | X | √ |  | 
| failure_reason | string | X | √ |  | 
| last_update_status | string | X | √ |  | 
| security_group_ids | string_array | X | √ |  | 
| tags | json | X | √ |  | 
| glue_version | string | X | √ |  | 
| last_modified_timestamp | timestamp | X | √ |  | 
| region | string | X | √ |  | 
| availability_zone | string | X | √ |  | 
| number_of_nodes | big_int | X | √ |  | 
| account_id | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| public_keys | string_array | X | √ |  | 
| role_arn | string | X | √ |  | 
| vpc_id | string | X | √ |  | 
| worker_type | string | X | √ |  | 
| zeppelin_remote_spark_interpreter_port | big_int | X | √ |  | 
| arguments | json | X | √ |  | 
| endpoint_name | string | X | √ |  | 
| private_address | string | X | √ |  | 
| subnet_id | string | X | √ |  | 
| yarn_endpoint_address | string | X | √ |  | 
| arn | string | √ | √ |  | 
| extra_python_libs_s3_path | string | X | √ |  | 
| number_of_workers | big_int | X | √ |  | 
| public_address | string | X | √ |  | 
| public_key | string | X | √ |  | 


