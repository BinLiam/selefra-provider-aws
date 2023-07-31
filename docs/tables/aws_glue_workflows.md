# Table: aws_glue_workflows

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| description | string | X | √ |  | 
| max_concurrent_runs | big_int | X | √ |  | 
| account_id | string | X | √ |  | 
| created_on | timestamp | X | √ |  | 
| default_run_properties | json | X | √ |  | 
| last_modified_on | timestamp | X | √ |  | 
| last_run | json | X | √ |  | 
| name | string | X | √ |  | 
| arn | string | √ | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| blueprint_details | json | X | √ |  | 
| graph | json | X | √ |  | 
| region | string | X | √ |  | 
| tags | json | X | √ |  | 


