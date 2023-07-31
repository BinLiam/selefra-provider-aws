# Table: aws_ssoadmin_account_assignments

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| aws_ssoadmin_permission_sets_selefra_id | string | X | X | fk to aws_ssoadmin_permission_sets.selefra_id | 
| account_id | string | X | √ |  | 
| permission_set_arn | string | X | √ |  | 
| principal_id | string | X | √ |  | 
| principal_type | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 


