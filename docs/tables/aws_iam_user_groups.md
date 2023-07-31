# Table: aws_iam_user_groups

## Primary Keys 

```
user_id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| group_name | string | X | √ |  | 
| path | string | X | √ |  | 
| user_arn | string | X | √ |  | 
| aws_iam_users_selefra_id | string | X | X | fk to aws_iam_users.selefra_id | 
| arn | string | X | √ |  | 
| group_id | string | X | √ |  | 
| account_id | string | X | √ |  | 
| user_id | string | √ | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| create_date | timestamp | X | √ |  | 


