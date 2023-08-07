# Table: aws_networkfirewall_firewall

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| arn | string | X | √ |  | 
| name | string | X | √ |  | 
| vpc_id | string | X | √ |  | 
| delete_protection | bool | X | √ |  | 
| description | string | X | √ |  | 
| id | string | X | √ |  | 
| policy_arn | string | X | √ |  | 
| policy_change_protection | bool | X | √ |  | 
| subnet_change_protection | bool | X | √ |  | 
| encryption_configuration | json | X | √ |  | 
| firewall_status | json | X | √ |  | 
| subnet_mappings | json | X | √ |  | 
| tags_src | json | X | √ |  | 


