# Table: aws_auditmanager_framework

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| name | string | X | √ |  | 
| arn | string | X | √ |  | 
| id | string | √ | √ |  | 
| type | string | X | √ |  | 
| created_at | timestamp | X | √ |  | 
| created_by | string | X | √ |  | 
| compliance_type | string | X | √ |  | 
| controls_count | int | X | √ |  | 
| control_sets_count | int | X | √ |  | 
| control_sources | string | X | √ |  | 
| description | string | X | √ |  | 
| last_updated_at | timestamp | X | √ |  | 
| last_updated_by | string | X | √ |  | 
| logo | string | X | √ |  | 
| control_sets | json | X | √ |  | 


