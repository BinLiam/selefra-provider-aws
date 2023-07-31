# Table: aws_eventbridge_event_bus_rules

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| state | string | X | √ |  | 
| event_bus_arn | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| aws_eventbridge_event_buses_selefra_id | string | X | X | fk to aws_eventbridge_event_buses.selefra_id | 
| role_arn | string | X | √ |  | 
| schedule_expression | string | X | √ |  | 
| managed_by | string | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| tags | json | X | √ |  | 
| arn | string | X | √ |  | 
| description | string | X | √ |  | 
| event_bus_name | string | X | √ |  | 
| event_pattern | string | X | √ |  | 
| name | string | X | √ |  | 


