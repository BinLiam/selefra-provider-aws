# Table: aws_iam_role_policies

## Primary Keys 

```
account_id, role_arn, policy_name
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| role_arn | string | X | √ |  | 
| aws_iam_roles_selefra_id | string | X | X | fk to aws_iam_roles.selefra_id | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| policy_document | json | X | √ |  | 
| policy_name | string | X | √ |  | 
| role_name | string | X | √ |  | 
| result_metadata | json | X | √ |  | 
| account_id | string | X | √ |  | 


