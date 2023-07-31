# Table: aws_glue_crawlers

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| classifiers | string_array | X | √ |  | 
| last_updated | timestamp | X | √ |  | 
| role | string | X | √ |  | 
| table_prefix | string | X | √ |  | 
| version | big_int | X | √ |  | 
| account_id | string | X | √ |  | 
| creation_time | timestamp | X | √ |  | 
| description | string | X | √ |  | 
| last_crawl | json | X | √ |  | 
| schedule | json | X | √ |  | 
| arn | string | √ | √ |  | 
| crawl_elapsed_time | big_int | X | √ |  | 
| lake_formation_configuration | json | X | √ |  | 
| lineage_configuration | json | X | √ |  | 
| tags | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| targets | json | X | √ |  | 
| configuration | string | X | √ |  | 
| crawler_security_configuration | string | X | √ |  | 
| database_name | string | X | √ |  | 
| name | string | X | √ |  | 
| recrawl_policy | json | X | √ |  | 
| schema_change_policy | json | X | √ |  | 
| state | string | X | √ |  | 
| region | string | X | √ |  | 


