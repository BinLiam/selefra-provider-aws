# Table: aws_mq_broker_configuration_revisions

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| broker_configuration_arn | string | X | √ |  | 
| aws_mq_broker_configurations_selefra_id | string | X | X | fk to aws_mq_broker_configurations.selefra_id | 
| description | string | X | √ |  | 
| result_metadata | json | X | √ |  | 
| data | json | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| configuration_id | string | X | √ |  | 
| created | timestamp | X | √ |  | 


