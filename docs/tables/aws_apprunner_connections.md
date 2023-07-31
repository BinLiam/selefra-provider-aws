# Table: aws_apprunner_connections

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| status | string | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| connection_name | string | X | √ |  | 
| created_at | timestamp | X | √ |  | 
| arn | string | √ | √ |  | 
| tags | json | X | √ |  | 
| connection_arn | string | X | √ |  | 
| provider_type | string | X | √ |  | 


