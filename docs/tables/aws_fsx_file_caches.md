# Table: aws_fsx_file_caches

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| file_cache_type_version | string | X | √ |  | 
| subnet_ids | string_array | X | √ |  | 
| account_id | string | X | √ |  | 
| dns_name | string | X | √ |  | 
| file_cache_id | string | X | √ |  | 
| file_cache_type | string | X | √ |  | 
| lifecycle | string | X | √ |  | 
| lustre_configuration | json | X | √ |  | 
| owner_id | string | X | √ |  | 
| storage_capacity | big_int | X | √ |  | 
| region | string | X | √ |  | 
| data_repository_association_ids | string_array | X | √ |  | 
| tags | json | X | √ |  | 
| arn | string | √ | √ |  | 
| failure_details | json | X | √ |  | 
| kms_key_id | string | X | √ |  | 
| resource_arn | string | X | √ |  | 
| vpc_id | string | X | √ |  | 
| creation_time | timestamp | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| network_interface_ids | string_array | X | √ |  | 


