# Table: aws_quicksight_analyses

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| arn | string | √ | √ |  | 
| created_time | timestamp | X | √ |  | 
| account_id | string | X | √ |  | 
| data_set_arns | string_array | X | √ |  | 
| last_updated_time | timestamp | X | √ |  | 
| region | string | X | √ |  | 
| tags | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| errors | json | X | √ |  | 
| status | string | X | √ |  | 
| theme_arn | string | X | √ |  | 
| analysis_id | string | X | √ |  | 
| name | string | X | √ |  | 
| sheets | json | X | √ |  | 


