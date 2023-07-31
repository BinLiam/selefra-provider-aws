# Table: aws_iam_users

## Primary Keys 

```
id, account_id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| arn | string | X | √ |  | 
| path | string | X | √ |  | 
| user_id | string | X | √ |  | 
| user_name | string | X | √ |  | 
| tags | json | X | √ |  | 
| id | string | X | √ |  | 
| account_id | string | X | √ |  | 
| create_date | timestamp | X | √ |  | 
| password_last_used | timestamp | X | √ |  | 
| permissions_boundary | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


