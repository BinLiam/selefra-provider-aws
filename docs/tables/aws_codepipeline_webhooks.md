# Table: aws_codepipeline_webhooks

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| definition | json | X | √ |  | 
| url | string | X | √ |  | 
| error_message | string | X | √ |  | 
| last_triggered | timestamp | X | √ |  | 
| tags | json | X | √ |  | 
| arn | string | √ | √ |  | 
| error_code | string | X | √ |  | 
| region | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


