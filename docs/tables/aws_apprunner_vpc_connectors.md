# Table: aws_apprunner_vpc_connectors

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| subnets | string_array | X | √ |  | 
| vpc_connector_arn | string | X | √ |  | 
| vpc_connector_name | string | X | √ |  | 
| account_id | string | X | √ |  | 
| arn | string | √ | √ |  | 
| created_at | timestamp | X | √ |  | 
| deleted_at | timestamp | X | √ |  | 
| security_groups | string_array | X | √ |  | 
| tags | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| status | string | X | √ |  | 
| vpc_connector_revision | big_int | X | √ |  | 
| region | string | X | √ |  | 


