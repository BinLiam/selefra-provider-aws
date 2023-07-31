# Table: aws_config_config_rule_compliances

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| compliance | json | X | √ |  | 
| config_rule_name | string | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| aws_config_config_rules_selefra_id | string | X | X | fk to aws_config_config_rules.selefra_id | 
| selefra_id | string | √ | √ | random id | 


