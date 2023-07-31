# Table: aws_fsx_storage_virtual_machines

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| lifecycle_transition_reason | json | X | √ |  | 
| storage_virtual_machine_id | string | X | √ |  | 
| uuid | string | X | √ |  | 
| account_id | string | X | √ |  | 
| arn | string | √ | √ |  | 
| endpoints | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| creation_time | timestamp | X | √ |  | 
| name | string | X | √ |  | 
| root_volume_security_style | string | X | √ |  | 
| subtype | string | X | √ |  | 
| lifecycle | string | X | √ |  | 
| file_system_id | string | X | √ |  | 
| resource_arn | string | X | √ |  | 
| tags | json | X | √ |  | 
| region | string | X | √ |  | 
| active_directory_configuration | json | X | √ |  | 


