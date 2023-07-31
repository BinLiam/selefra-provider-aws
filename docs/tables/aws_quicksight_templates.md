# Table: aws_quicksight_templates

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| latest_version_number | big_int | X | √ |  | 
| tags | json | X | √ |  | 
| name | string | X | √ |  | 
| template_id | string | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| arn | string | √ | √ |  | 
| created_time | timestamp | X | √ |  | 
| last_updated_time | timestamp | X | √ |  | 


