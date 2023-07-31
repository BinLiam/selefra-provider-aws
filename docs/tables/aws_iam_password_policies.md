# Table: aws_iam_password_policies

## Primary Keys 

```
account_id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | √ | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| password_policy | json | X | √ |  | 
| policy_exists | bool | X | √ |  | 


