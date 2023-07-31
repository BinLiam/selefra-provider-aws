# Table: aws_glue_database_tables

## Primary Keys 

```
database_arn, name
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| created_by | string | X | √ |  | 
| account_id | string | X | √ |  | 
| aws_glue_databases_selefra_id | string | X | X | fk to aws_glue_databases.selefra_id | 
| create_time | timestamp | X | √ |  | 
| database_name | string | X | √ |  | 
| is_registered_with_lake_formation | bool | X | √ |  | 
| last_access_time | timestamp | X | √ |  | 
| partition_keys | json | X | √ |  | 
| storage_descriptor | json | X | √ |  | 
| view_original_text | string | X | √ |  | 
| region | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| name | string | X | √ |  | 
| description | string | X | √ |  | 
| last_analyzed_time | timestamp | X | √ |  | 
| owner | string | X | √ |  | 
| parameters | json | X | √ |  | 
| table_type | string | X | √ |  | 
| version_id | string | X | √ |  | 
| catalog_id | string | X | √ |  | 
| retention | big_int | X | √ |  | 
| target_table | json | X | √ |  | 
| update_time | timestamp | X | √ |  | 
| view_expanded_text | string | X | √ |  | 
| database_arn | string | X | √ |  | 


