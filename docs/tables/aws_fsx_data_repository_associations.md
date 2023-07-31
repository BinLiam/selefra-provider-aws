# Table: aws_fsx_data_repository_associations

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| file_system_id | string | X | √ |  | 
| arn | string | √ | √ |  | 
| association_id | string | X | √ |  | 
| batch_import_meta_data_on_create | bool | X | √ |  | 
| file_cache_path | string | X | √ |  | 
| nfs | json | X | √ |  | 
| s3 | json | X | √ |  | 
| account_id | string | X | √ |  | 
| file_cache_id | string | X | √ |  | 
| imported_file_chunk_size | big_int | X | √ |  | 
| lifecycle | string | X | √ |  | 
| resource_arn | string | X | √ |  | 
| tags | json | X | √ |  | 
| creation_time | timestamp | X | √ |  | 
| data_repository_path | string | X | √ |  | 
| data_repository_subdirectories | string_array | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| failure_details | json | X | √ |  | 
| file_system_path | string | X | √ |  | 
| region | string | X | √ |  | 


