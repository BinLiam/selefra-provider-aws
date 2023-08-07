# Table: aws_auditmanager_control

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
| action_plan_title | string | X | √ |  | 
| action_plan_instructions | string | X | √ |  | 
| control_sources | string | X | √ |  | 
| description | string | X | √ |  | 
| last_updated_at | timestamp | X | √ |  | 
| last_updated_by | string | X | √ |  | 
| testing_information | string | X | √ |  | 
| control_mapping_sources | json | X | √ |  | 


