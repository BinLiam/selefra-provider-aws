# Table: aws_ecr_registries

## Primary Keys 

```
account_id, region, registry_id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| region | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| registry_id | string | X | √ |  | 
| replication_configuration | json | X | √ |  | 
| result_metadata | json | X | √ |  | 
| account_id | string | X | √ |  | 


