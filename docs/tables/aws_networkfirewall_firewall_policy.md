# Table: aws_networkfirewall_firewall_policy

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| name | string | X | √ |  | 
| arn | string | X | √ |  | 
| firewall_policy_id | string | X | √ |  | 
| consumed_stateful_rule_capacity | int | X | √ |  | 
| consumed_stateless_rule_capacity | int | X | √ |  | 
| description | string | X | √ |  | 
| firewall_policy_status | string | X | √ |  | 
| last_modified_time | timestamp | X | √ |  | 
| number_of_associations | int | X | √ |  | 
| encryption_configuration | json | X | √ |  | 
| firewall_policy | json | X | √ |  | 
| tags | json | X | √ |  | 


