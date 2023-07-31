# Table: aws_appstream_app_blocks

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| setup_script_details | json | X | √ |  | 
| description | string | X | √ |  | 
| display_name | string | X | √ |  | 
| source_s3_location | json | X | √ |  | 
| account_id | string | X | √ |  | 
| name | string | X | √ |  | 
| created_time | timestamp | X | √ |  | 
| region | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| arn | string | √ | √ |  | 


