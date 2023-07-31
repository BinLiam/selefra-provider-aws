# Table: aws_quicksight_dashboards

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| arn | string | √ | √ |  | 
| created_time | timestamp | X | √ |  | 
| dashboard_id | string | X | √ |  | 
| last_published_time | timestamp | X | √ |  | 
| published_version_number | big_int | X | √ |  | 
| last_updated_time | timestamp | X | √ |  | 
| name | string | X | √ |  | 
| tags | json | X | √ |  | 


