# Table: aws_backup_vault_recovery_points

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| last_restore_time | timestamp | X | √ |  | 
| recovery_point_arn | string | X | √ |  | 
| source_backup_vault_arn | string | X | √ |  | 
| status_message | string | X | √ |  | 
| is_encrypted | bool | X | √ |  | 
| resource_arn | string | X | √ |  | 
| resource_type | string | X | √ |  | 
| is_parent | bool | X | √ |  | 
| parent_recovery_point_arn | string | X | √ |  | 
| account_id | string | X | √ |  | 
| composite_member_identifier | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| iam_role_arn | string | X | √ |  | 
| tags | json | X | √ |  | 
| aws_backup_vaults_selefra_id | string | X | X | fk to aws_backup_vaults.selefra_id | 
| backup_vault_name | string | X | √ |  | 
| calculated_lifecycle | json | X | √ |  | 
| creation_date | timestamp | X | √ |  | 
| vault_arn | string | X | √ |  | 
| encryption_key_arn | string | X | √ |  | 
| lifecycle | json | X | √ |  | 
| status | string | X | √ |  | 
| backup_size_in_bytes | big_int | X | √ |  | 
| completion_date | timestamp | X | √ |  | 
| arn | string | √ | √ |  | 
| backup_vault_arn | string | X | √ |  | 
| created_by | json | X | √ |  | 
| region | string | X | √ |  | 


