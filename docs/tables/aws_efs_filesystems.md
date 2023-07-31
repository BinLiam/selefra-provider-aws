# Table: aws_efs_filesystems

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| availability_zone_name | string | X | √ |  | 
| file_system_arn | string | X | √ |  | 
| region | string | X | √ |  | 
| backup_policy_status | string | X | √ |  | 
| tags | json | X | √ |  | 
| availability_zone_id | string | X | √ |  | 
| size_in_bytes | json | X | √ |  | 
| kms_key_id | string | X | √ |  | 
| account_id | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| file_system_id | string | X | √ |  | 
| performance_mode | string | X | √ |  | 
| name | string | X | √ |  | 
| throughput_mode | string | X | √ |  | 
| arn | string | √ | √ |  | 
| number_of_mount_targets | big_int | X | √ |  | 
| encrypted | bool | X | √ |  | 
| life_cycle_state | string | X | √ |  | 
| owner_id | string | X | √ |  | 
| provisioned_throughput_in_mibps | float | X | √ |  | 
| creation_time | timestamp | X | √ |  | 
| creation_token | string | X | √ |  | 


