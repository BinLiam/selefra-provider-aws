# Table: aws_elbv2_listener_certificates

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| is_default | bool | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| listener_arn | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| aws_elbv2_listeners_selefra_id | string | X | X | fk to aws_elbv2_listeners.selefra_id | 
| certificate_arn | string | X | √ |  | 


