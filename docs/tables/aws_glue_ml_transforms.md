# Table: aws_glue_ml_transforms

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| input_record_tables | json | X | √ |  | 
| last_modified_on | timestamp | X | √ |  | 
| schema | json | X | √ |  | 
| status | string | X | √ |  | 
| arn | string | √ | √ |  | 
| evaluation_metrics | json | X | √ |  | 
| glue_version | string | X | √ |  | 
| max_capacity | float | X | √ |  | 
| max_retries | big_int | X | √ |  | 
| name | string | X | √ |  | 
| number_of_workers | big_int | X | √ |  | 
| transform_encryption | json | X | √ |  | 
| account_id | string | X | √ |  | 
| created_on | timestamp | X | √ |  | 
| label_count | big_int | X | √ |  | 
| transform_id | string | X | √ |  | 
| worker_type | string | X | √ |  | 
| role | string | X | √ |  | 
| timeout | big_int | X | √ |  | 
| region | string | X | √ |  | 
| tags | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| description | string | X | √ |  | 
| parameters | json | X | √ |  | 


