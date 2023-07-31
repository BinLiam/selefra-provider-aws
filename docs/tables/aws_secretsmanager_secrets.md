# Table: aws_secretsmanager_secrets

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| last_changed_date | timestamp | X | √ |  | 
| last_rotated_date | timestamp | X | √ |  | 
| rotation_rules | json | X | √ |  | 
| result_metadata | json | X | √ |  | 
| created_date | timestamp | X | √ |  | 
| deleted_date | timestamp | X | √ |  | 
| description | string | X | √ |  | 
| name | string | X | √ |  | 
| primary_region | string | X | √ |  | 
| replication_status | json | X | √ |  | 
| account_id | string | X | √ |  | 
| arn | string | √ | √ |  | 
| kms_key_id | string | X | √ |  | 
| next_rotation_date | timestamp | X | √ |  | 
| owning_service | string | X | √ |  | 
| rotation_enabled | bool | X | √ |  | 
| tags | json | X | √ |  | 
| region | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| last_accessed_date | timestamp | X | √ |  | 
| rotation_lambda_arn | string | X | √ |  | 
| version_ids_to_stages | json | X | √ |  | 
| policy | json | X | √ | `A JSON-formatted string that describes the permissions that are associated with the attached secret.` | 


