# Table: aws_fsx_data_repository_tasks

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| arn | string | √ | √ |  | 
| creation_time | timestamp | X | √ |  | 
| task_id | string | X | √ |  | 
| type | string | X | √ |  | 
| capacity_to_release | big_int | X | √ |  | 
| status | json | X | √ |  | 
| region | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| start_time | timestamp | X | √ |  | 
| lifecycle | string | X | √ |  | 
| end_time | timestamp | X | √ |  | 
| paths | string_array | X | √ |  | 
| report | json | X | √ |  | 
| tags | json | X | √ |  | 
| account_id | string | X | √ |  | 
| failure_details | json | X | √ |  | 
| file_cache_id | string | X | √ |  | 
| file_system_id | string | X | √ |  | 
| resource_arn | string | X | √ |  | 


