# Table: aws_config_conformance_pack_rule_compliances

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| config_rule_invoked_time | timestamp | X | √ |  | 
| evaluation_result_identifier | json | X | √ |  | 
| aws_config_conformance_packs_selefra_id | string | X | X | fk to aws_config_conformance_packs.selefra_id | 
| compliance_type | string | X | √ |  | 
| config_rule_name | string | X | √ |  | 
| controls | string_array | X | √ |  | 
| region | string | X | √ |  | 
| conformance_pack_arn | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| result_recorded_time | timestamp | X | √ |  | 
| annotation | string | X | √ |  | 
| account_id | string | X | √ |  | 


