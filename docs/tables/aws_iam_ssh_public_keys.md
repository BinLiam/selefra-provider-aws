# Table: aws_iam_ssh_public_keys

## Primary Keys 

```
ssh_public_key_id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| upload_date | timestamp | X | √ |  | 
| user_name | string | X | √ |  | 
| account_id | string | X | √ |  | 
| aws_iam_users_selefra_id | string | X | X | fk to aws_iam_users.selefra_id | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| ssh_public_key_id | string | √ | √ |  | 
| status | string | X | √ |  | 
| user_arn | string | X | √ |  | 
| user_id | string | X | √ |  | 


