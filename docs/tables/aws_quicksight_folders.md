# Table: aws_quicksight_folders

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| arn | string | √ | √ |  | 
| created_time | timestamp | X | √ |  | 
| last_updated_time | timestamp | X | √ |  | 
| tags | json | X | √ |  | 
| folder_id | string | X | √ |  | 
| folder_path | string_array | X | √ |  | 
| folder_type | string | X | √ |  | 
| name | string | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


