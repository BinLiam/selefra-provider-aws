# Table: aws_glue_triggers

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| workflow_name | string | X | √ |  | 
| region | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| actions | json | X | √ |  | 
| event_batching_condition | json | X | √ |  | 
| state | string | X | √ |  | 
| arn | string | √ | √ |  | 
| tags | json | X | √ |  | 
| predicate | json | X | √ |  | 
| schedule | string | X | √ |  | 
| description | string | X | √ |  | 
| id | string | X | √ |  | 
| account_id | string | X | √ |  | 
| name | string | X | √ |  | 
| type | string | X | √ |  | 


