# Table: aws_wafregional_rule_groups

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| metric_name | string | X | √ |  | 
| name | string | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| arn | string | √ | √ |  | 
| tags | json | X | √ | `Rule group tags.` | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| rule_group_id | string | X | √ |  | 


