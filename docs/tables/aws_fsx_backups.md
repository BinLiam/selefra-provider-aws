# Table: aws_fsx_backups

## Primary Keys 

```
account_id, region, id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| progress_percent | big_int | X | √ |  | 
| resource_arn | string | X | √ |  | 
| volume | json | X | √ |  | 
| tags | json | X | √ |  | 
| account_id | string | X | √ |  | 
| id | string | X | √ |  | 
| backup_id | string | X | √ |  | 
| file_system | json | X | √ |  | 
| kms_key_id | string | X | √ |  | 
| owner_id | string | X | √ |  | 
| source_backup_id | string | X | √ |  | 
| directory_information | json | X | √ |  | 
| failure_details | json | X | √ |  | 
| resource_type | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| creation_time | timestamp | X | √ |  | 
| lifecycle | string | X | √ |  | 
| type | string | X | √ |  | 
| source_backup_region | string | X | √ |  | 
| region | string | X | √ |  | 


