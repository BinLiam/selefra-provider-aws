# Table: aws_waf_subscribed_rule_groups

## Primary Keys 

```
account_id, rule_group_id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| name | string | X | √ |  | 
| rule_group_id | string | X | √ | `A unique identifier for a RuleGroup.` | 
| account_id | string | X | √ | `The AWS Account ID of the resource.` | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| metric_name | string | X | √ |  | 


