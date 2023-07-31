# Table: aws_ssoadmin_permission_sets

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| session_duration | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| name | string | X | √ |  | 
| relay_state | string | X | √ |  | 
| permission_set_arn | string | X | √ |  | 
| inline_policy | json | X | √ |  | 
| aws_ssoadmin_instances_selefra_id | string | X | X | fk to aws_ssoadmin_instances.selefra_id | 
| created_date | timestamp | X | √ |  | 
| description | string | X | √ |  | 


