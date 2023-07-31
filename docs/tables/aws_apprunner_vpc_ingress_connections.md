# Table: aws_apprunner_vpc_ingress_connections

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| ingress_vpc_configuration | json | X | √ |  | 
| vpc_ingress_connection_name | string | X | √ |  | 
| deleted_at | timestamp | X | √ |  | 
| created_at | timestamp | X | √ |  | 
| status | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | X | √ |  | 
| service_arn | string | X | √ |  | 
| region | string | X | √ |  | 
| arn | string | √ | √ |  | 
| source_account_id | string | X | √ |  | 
| tags | json | X | √ |  | 
| domain_name | string | X | √ |  | 
| vpc_ingress_connection_arn | string | X | √ |  | 


