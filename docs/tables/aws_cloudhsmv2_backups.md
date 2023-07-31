# Table: aws_cloudhsmv2_backups

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| source_cluster | string | X | √ |  | 
| backup_id | string | X | √ |  | 
| copy_timestamp | timestamp | X | √ |  | 
| tag_list | json | X | √ |  | 
| delete_timestamp | timestamp | X | √ |  | 
| source_backup | string | X | √ |  | 
| source_region | string | X | √ |  | 
| account_id | string | X | √ |  | 
| arn | string | √ | √ |  | 
| tags | json | X | √ |  | 
| backup_state | string | X | √ |  | 
| cluster_id | string | X | √ |  | 
| region | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| create_timestamp | timestamp | X | √ |  | 
| never_expires | bool | X | √ |  | 


