# Table: aws_iam_user_access_keys

## Primary Keys 

```
account_id, user_arn, access_key_id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| access_key_metadata | json | X | √ |  | 
| last_rotated | timestamp | X | √ |  | 
| last_used | timestamp | X | √ |  | 
| aws_iam_users_selefra_id | string | X | X | fk to aws_iam_users.selefra_id | 
| account_id | string | X | √ |  | 
| user_arn | string | X | √ |  | 
| access_key_id | string | X | √ |  | 
| user_id | string | X | √ |  | 
| last_used_service_name | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


