# Table: aws_cloudtrail_trail_event_selectors

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| include_management_events | bool | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| trail_arn | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| aws_cloudtrail_trails_selefra_id | string | X | X | fk to aws_cloudtrail_trails.selefra_id | 
| exclude_management_event_sources | string_array | X | √ |  | 
| read_write_type | string | X | √ |  | 
| data_resources | json | X | √ |  | 


