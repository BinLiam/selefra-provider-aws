# Table: aws_accessanalyzer_analyzer_findings

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| is_public | bool | X | √ |  | 
| principal | json | X | √ |  | 
| analyzer_arn | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| aws_accessanalyzer_analyzers_selefra_id | string | X | X | fk to aws_accessanalyzer_analyzers.selefra_id | 
| condition | json | X | √ |  | 
| created_at | timestamp | X | √ |  | 
| resource_type | string | X | √ |  | 
| arn | string | √ | √ |  | 
| status | string | X | √ |  | 
| resource | string | X | √ |  | 
| account_id | string | X | √ |  | 
| error | string | X | √ |  | 
| analyzed_at | timestamp | X | √ |  | 
| id | string | X | √ |  | 
| resource_owner_account | string | X | √ |  | 
| region | string | X | √ |  | 
| updated_at | timestamp | X | √ |  | 
| action | string_array | X | √ |  | 
| sources | json | X | √ |  | 


