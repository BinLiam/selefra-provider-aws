# Table: aws_config_config_rules

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| source | json | X | √ |  | 
| config_rule_arn | string | X | √ |  | 
| config_rule_id | string | X | √ |  | 
| description | string | X | √ |  | 
| scope | json | X | √ |  | 
| created_by | string | X | √ |  | 
| arn | string | √ | √ |  | 
| config_rule_name | string | X | √ |  | 
| config_rule_state | string | X | √ |  | 
| evaluation_modes | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| input_parameters | string | X | √ |  | 
| maximum_execution_frequency | string | X | √ |  | 
| region | string | X | √ |  | 


