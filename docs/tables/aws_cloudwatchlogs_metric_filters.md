# Table: aws_cloudwatchlogs_metric_filters

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| arn | string | √ | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| creation_time | big_int | X | √ |  | 
| filter_name | string | X | √ |  | 
| filter_pattern | string | X | √ |  | 
| log_group_name | string | X | √ |  | 
| metric_transformations | json | X | √ |  | 


