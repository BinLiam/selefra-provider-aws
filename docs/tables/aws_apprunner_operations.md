# Table: aws_apprunner_operations

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| status | string | X | √ |  | 
| type | string | X | √ |  | 
| updated_at | timestamp | X | √ |  | 
| aws_apprunner_services_selefra_id | string | X | X | fk to aws_apprunner_services.selefra_id | 
| selefra_id | string | √ | √ | random id | 
| ended_at | timestamp | X | √ |  | 
| id | string | X | √ |  | 
| started_at | timestamp | X | √ |  | 
| target_arn | string | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 


