# Table: aws_route53_health_checks

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| arn | string | √ | √ |  | 
| cloud_watch_alarm_configuration_dimensions | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| health_check | json | X | √ |  | 
| tags | json | X | √ | `The tags associated with the health check.` | 
| account_id | string | X | √ |  | 


