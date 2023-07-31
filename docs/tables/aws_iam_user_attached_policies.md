# Table: aws_iam_user_attached_policies

## Primary Keys 

```
account_id, user_arn, policy_name
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| policy_arn | string | X | √ |  | 
| policy_name | string | X | √ |  | 
| account_id | string | X | √ |  | 
| user_arn | string | X | √ |  | 
| user_id | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| aws_iam_users_selefra_id | string | X | X | fk to aws_iam_users.selefra_id | 


