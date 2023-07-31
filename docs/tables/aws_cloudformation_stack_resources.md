# Table: aws_cloudformation_stack_resources

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| logical_resource_id | string | X | √ |  | 
| resource_type | string | X | √ |  | 
| physical_resource_id | string | X | √ |  | 
| account_id | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| last_updated_timestamp | timestamp | X | √ |  | 
| resource_status | string | X | √ |  | 
| drift_information | json | X | √ |  | 
| module_info | json | X | √ |  | 
| resource_status_reason | string | X | √ |  | 
| region | string | X | √ |  | 
| aws_cloudformation_stacks_selefra_id | string | X | X | fk to aws_cloudformation_stacks.selefra_id | 


