# Table: aws_servicequotas_quotas

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| global_quota | bool | X | √ |  | 
| quota_arn | string | X | √ |  | 
| service_code | string | X | √ |  | 
| value | float | X | √ |  | 
| region | string | X | √ |  | 
| arn | string | √ | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| usage_metric | json | X | √ |  | 
| error_reason | json | X | √ |  | 
| quota_code | string | X | √ |  | 
| quota_name | string | X | √ |  | 
| service_name | string | X | √ |  | 
| unit | string | X | √ |  | 
| adjustable | bool | X | √ |  | 
| period | json | X | √ |  | 
| account_id | string | X | √ |  | 
| aws_servicequotas_services_selefra_id | string | X | X | fk to aws_servicequotas_services.selefra_id | 


