# Table: aws_fsx_file_systems

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| failure_details | json | X | √ |  | 
| lustre_configuration | json | X | √ |  | 
| resource_arn | string | X | √ |  | 
| windows_configuration | json | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| creation_time | timestamp | X | √ |  | 
| kms_key_id | string | X | √ |  | 
| tags | json | X | √ |  | 
| administrative_actions | json | X | √ |  | 
| lifecycle | string | X | √ |  | 
| network_interface_ids | string_array | X | √ |  | 
| open_zfs_configuration | json | X | √ |  | 
| owner_id | string | X | √ |  | 
| subnet_ids | string_array | X | √ |  | 
| vpc_id | string | X | √ |  | 
| arn | string | √ | √ |  | 
| dns_name | string | X | √ |  | 
| file_system_type | string | X | √ |  | 
| file_system_type_version | string | X | √ |  | 
| ontap_configuration | json | X | √ |  | 
| storage_capacity | big_int | X | √ |  | 
| storage_type | string | X | √ |  | 
| file_system_id | string | X | √ |  | 


