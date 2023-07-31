# Table: aws_iam_roles

## Primary Keys 

```
account_id, id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| id | string | X | √ |  | 
| path | string | X | √ |  | 
| role_id | string | X | √ |  | 
| role_name | string | X | √ |  | 
| description | string | X | √ |  | 
| account_id | string | X | √ |  | 
| assume_role_policy_document | json | X | √ |  | 
| role_last_used | json | X | √ |  | 
| tags | json | X | √ |  | 
| policies | json | X | √ |  | 
| arn | string | X | √ |  | 
| create_date | timestamp | X | √ |  | 
| max_session_duration | big_int | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| permissions_boundary | json | X | √ |  | 


