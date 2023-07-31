# Table: aws_fsx_volumes

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| open_zfs_configuration | json | X | √ |  | 
| resource_arn | string | X | √ |  | 
| volume_type | string | X | √ |  | 
| file_system_id | string | X | √ |  | 
| name | string | X | √ |  | 
| volume_id | string | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| administrative_actions | json | X | √ |  | 
| lifecycle_transition_reason | json | X | √ |  | 
| ontap_configuration | json | X | √ |  | 
| arn | string | √ | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| creation_time | timestamp | X | √ |  | 
| lifecycle | string | X | √ |  | 
| tags | json | X | √ |  | 


