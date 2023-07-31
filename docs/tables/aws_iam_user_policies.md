# Table: aws_iam_user_policies

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| user_name | string | X | √ |  | 
| user_arn | string | X | √ |  | 
| aws_iam_users_selefra_id | string | X | X | fk to aws_iam_users.selefra_id | 
| policy_document | json | X | √ |  | 
| policy_name | string | X | √ |  | 
| result_metadata | json | X | √ |  | 
| account_id | string | X | √ |  | 
| user_id | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 


