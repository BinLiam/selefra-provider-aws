# Table: aws_iam_group_policies

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| policy_document | json | X | √ |  | 
| result_metadata | json | X | √ |  | 
| account_id | string | X | √ |  | 
| group_id | string | X | √ |  | 
| group_name | string | X | √ |  | 
| policy_name | string | X | √ |  | 
| group_arn | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| aws_iam_groups_selefra_id | string | X | X | fk to aws_iam_groups.selefra_id | 


