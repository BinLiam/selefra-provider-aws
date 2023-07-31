# Table: aws_cloudformation_stacks

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| stack_name | string | X | √ |  | 
| capabilities | string_array | X | √ |  | 
| parameters | json | X | √ |  | 
| role_arn | string | X | √ |  | 
| stack_status | string | X | √ |  | 
| notification_ar_ns | string_array | X | √ |  | 
| disable_rollback | bool | X | √ |  | 
| parent_id | string | X | √ |  | 
| root_id | string | X | √ |  | 
| arn | string | √ | √ |  | 
| change_set_id | string | X | √ |  | 
| deletion_time | timestamp | X | √ |  | 
| description | string | X | √ |  | 
| enable_termination_protection | bool | X | √ |  | 
| region | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| drift_information | json | X | √ |  | 
| rollback_configuration | json | X | √ |  | 
| stack_status_reason | string | X | √ |  | 
| last_updated_time | timestamp | X | √ |  | 
| outputs | json | X | √ |  | 
| account_id | string | X | √ |  | 
| id | string | X | √ |  | 
| creation_time | timestamp | X | √ |  | 
| stack_id | string | X | √ |  | 
| tags | json | X | √ |  | 
| timeout_in_minutes | big_int | X | √ |  | 


