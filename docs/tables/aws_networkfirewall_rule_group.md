# Table: aws_networkfirewall_rule_group

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| rule_group_name | string | X | √ |  | 
| arn | string | X | √ |  | 
| capacity | int | X | √ |  | 
| consumed_capacity | int | X | √ |  | 
| description | string | X | √ |  | 
| number_of_associations | int | X | √ |  | 
| rule_group_id | string | X | √ |  | 
| rule_group_status | string | X | √ |  | 
| rule_variables | json | X | √ |  | 
| rules_source | json | X | √ |  | 
| stateful_rule_options | json | X | √ |  | 
| type | string | X | √ |  | 
| tags | json | X | √ |  | 


