# Table: aws_wafregional_rate_based_rules

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| rate_limit | big_int | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| tags | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| match_predicates | json | X | √ |  | 
| rule_id | string | X | √ |  | 
| metric_name | string | X | √ |  | 
| name | string | X | √ |  | 
| arn | string | √ | √ |  | 
| rate_key | string | X | √ |  | 


