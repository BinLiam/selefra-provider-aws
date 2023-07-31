# Table: aws_mq_broker_configurations

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| engine_type | string | X | √ |  | 
| result_metadata | json | X | √ |  | 
| account_id | string | X | √ |  | 
| name | string | X | √ |  | 
| tags | json | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| authentication_strategy | string | X | √ |  | 
| created | timestamp | X | √ |  | 
| description | string | X | √ |  | 
| id | string | X | √ |  | 
| latest_revision | json | X | √ |  | 
| broker_arn | string | X | √ |  | 
| aws_mq_brokers_selefra_id | string | X | X | fk to aws_mq_brokers.selefra_id | 
| arn | string | X | √ |  | 
| engine_version | string | X | √ |  | 
| region | string | X | √ |  | 


