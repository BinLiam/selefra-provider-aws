# Table: aws_wafv2_rule_groups

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| description | string | X | √ |  | 
| tags | json | X | √ |  | 
| policy | json | X | √ |  | 
| capacity | big_int | X | √ |  | 
| name | string | X | √ |  | 
| rules | json | X | √ |  | 
| account_id | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| consumed_labels | json | X | √ |  | 
| custom_response_bodies | json | X | √ |  | 
| label_namespace | string | X | √ |  | 
| arn | string | √ | √ |  | 
| id | string | X | √ |  | 
| visibility_config | json | X | √ |  | 
| available_labels | json | X | √ |  | 
| region | string | X | √ |  | 


