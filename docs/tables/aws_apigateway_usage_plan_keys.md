# Table: aws_apigateway_usage_plan_keys

## Primary Keys 

```
account_id, arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| value | string | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| arn | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| aws_apigateway_usage_plans_selefra_id | string | X | X | fk to aws_apigateway_usage_plans.selefra_id | 
| id | string | X | √ |  | 
| type | string | X | √ |  | 
| usage_plan_arn | string | X | √ |  | 
| name | string | X | √ |  | 


