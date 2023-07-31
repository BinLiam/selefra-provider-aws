# Table: aws_cognito_user_pool_identity_providers

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| provider_details | json | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| idp_identifiers | string_array | X | √ |  | 
| creation_date | timestamp | X | √ |  | 
| last_modified_date | timestamp | X | √ |  | 
| provider_name | string | X | √ |  | 
| provider_type | string | X | √ |  | 
| user_pool_id | string | X | √ |  | 
| user_pool_arn | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| attribute_mapping | json | X | √ |  | 
| aws_cognito_user_pools_selefra_id | string | X | X | fk to aws_cognito_user_pools.selefra_id | 


