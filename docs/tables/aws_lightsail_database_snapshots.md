# Table: aws_lightsail_database_snapshots

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| arn | string | √ | √ |  | 
| created_at | timestamp | X | √ |  | 
| engine | string | X | √ |  | 
| engine_version | string | X | √ |  | 
| from_relational_database_bundle_id | string | X | √ |  | 
| account_id | string | X | √ |  | 
| from_relational_database_arn | string | X | √ |  | 
| state | string | X | √ |  | 
| region | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| from_relational_database_name | string | X | √ |  | 
| location | json | X | √ |  | 
| name | string | X | √ |  | 
| from_relational_database_blueprint_id | string | X | √ |  | 
| resource_type | string | X | √ |  | 
| size_in_gb | big_int | X | √ |  | 
| support_code | string | X | √ |  | 
| tags | json | X | √ |  | 


