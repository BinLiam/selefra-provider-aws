# Table: aws_docdb_engine_versions

## Primary Keys 

```
account_id, region
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| supports_log_exports_to_cloudwatch_logs | bool | X | √ |  | 
| valid_upgrade_target | json | X | √ |  | 
| region | string | X | √ |  | 
| db_engine_version_description | string | X | √ |  | 
| exportable_log_types | string_array | X | √ |  | 
| engine | string | X | √ |  | 
| engine_version | string | X | √ |  | 
| account_id | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| db_engine_description | string | X | √ |  | 
| db_parameter_group_family | string | X | √ |  | 


