# Table: aws_docdb_events

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| date | timestamp | X | √ |  | 
| message | string | X | √ |  | 
| source_type | string | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| event_categories | string_array | X | √ |  | 
| source_arn | string | X | √ |  | 
| source_identifier | string | X | √ |  | 


