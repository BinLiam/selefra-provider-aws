# Table: aws_mq_broker_users

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| region | string | X | √ |  | 
| broker_arn | string | X | √ |  | 
| groups | string_array | X | √ |  | 
| pending | json | X | √ |  | 
| username | string | X | √ |  | 
| result_metadata | json | X | √ |  | 
| aws_mq_brokers_selefra_id | string | X | X | fk to aws_mq_brokers.selefra_id | 
| broker_id | string | X | √ |  | 
| console_access | bool | X | √ |  | 
| account_id | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 


