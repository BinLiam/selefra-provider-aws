# Table: aws_inspector_findings

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| service | string | X | √ |  | 
| service_attributes | json | X | √ |  | 
| account_id | string | X | √ |  | 
| arn | string | √ | √ |  | 
| created_at | timestamp | X | √ |  | 
| id | string | X | √ |  | 
| numeric_severity | float | X | √ |  | 
| schema_version | big_int | X | √ |  | 
| severity | string | X | √ |  | 
| title | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| region | string | X | √ |  | 
| asset_attributes | json | X | √ |  | 
| asset_type | string | X | √ |  | 
| description | string | X | √ |  | 
| recommendation | string | X | √ |  | 
| indicator_of_compromise | bool | X | √ |  | 
| attributes | json | X | √ |  | 
| updated_at | timestamp | X | √ |  | 
| user_attributes | json | X | √ |  | 
| confidence | big_int | X | √ |  | 


