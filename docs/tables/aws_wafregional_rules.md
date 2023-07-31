# Table: aws_wafregional_rules

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| arn | string | √ | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| rule_id | string | X | √ |  | 
| metric_name | string | X | √ |  | 
| name | string | X | √ |  | 
| region | string | X | √ |  | 
| predicates | json | X | √ |  | 
| account_id | string | X | √ |  | 
| tags | json | X | √ | `Rule tags.` | 


