# Table: aws_backup_vaults

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| access_policy | json | X | √ |  | 
| encryption_key_arn | string | X | √ |  | 
| region | string | X | √ |  | 
| account_id | string | X | √ |  | 
| arn | string | √ | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| backup_vault_arn | string | X | √ |  | 
| min_retention_days | big_int | X | √ |  | 
| locked | bool | X | √ |  | 
| tags | json | X | √ |  | 
| creator_request_id | string | X | √ |  | 
| lock_date | timestamp | X | √ |  | 
| max_retention_days | big_int | X | √ |  | 
| number_of_recovery_points | big_int | X | √ |  | 
| notifications | json | X | √ |  | 
| backup_vault_name | string | X | √ |  | 
| creation_date | timestamp | X | √ |  | 


