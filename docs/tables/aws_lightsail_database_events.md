# Table: aws_lightsail_database_events

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| aws_lightsail_databases_selefra_id | string | X | X | fk to aws_lightsail_databases.selefra_id | 
| message | string | X | √ |  | 
| resource | string | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| created_at | timestamp | X | √ |  | 
| event_categories | string_array | X | √ |  | 
| database_arn | string | X | √ |  | 


