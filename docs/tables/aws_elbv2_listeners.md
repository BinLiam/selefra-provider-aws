# Table: aws_elbv2_listeners

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| region | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| alpn_policy | string_array | X | √ |  | 
| listener_arn | string | X | √ |  | 
| protocol | string | X | √ |  | 
| certificates | json | X | √ |  | 
| default_actions | json | X | √ |  | 
| arn | string | √ | √ |  | 
| tags | json | X | √ |  | 
| port | big_int | X | √ |  | 
| ssl_policy | string | X | √ |  | 
| account_id | string | X | √ |  | 
| load_balancer_arn | string | X | √ |  | 
| aws_elbv2_load_balancers_selefra_id | string | X | X | fk to aws_elbv2_load_balancers.selefra_id | 


