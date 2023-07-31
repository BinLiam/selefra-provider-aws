# Table: aws_accessanalyzer_analyzers

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| tags | json | X | √ |  | 
| account_id | string | X | √ |  | 
| arn | string | √ | √ |  | 
| name | string | X | √ |  | 
| status | string | X | √ |  | 
| type | string | X | √ |  | 
| last_resource_analyzed_at | timestamp | X | √ |  | 
| status_reason | json | X | √ |  | 
| region | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| created_at | timestamp | X | √ |  | 
| last_resource_analyzed | string | X | √ |  | 


