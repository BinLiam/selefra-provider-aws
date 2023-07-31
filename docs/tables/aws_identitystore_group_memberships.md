# Table: aws_identitystore_group_memberships

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| identity_store_id | string | X | √ |  | 
| group_id | string | X | √ |  | 
| membership_id | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| aws_identitystore_groups_selefra_id | string | X | X | fk to aws_identitystore_groups.selefra_id | 


