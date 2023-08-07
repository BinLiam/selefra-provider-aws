# Table: aws_cloudformation_stack_set

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| stack_set_id | string | X | √ |  | 
| stack_set_name | string | X | √ |  | 
| arn | string | √ | √ |  | 
| status | string | X | √ |  | 
| description | string | X | √ |  | 
| drift_status | string | X | √ |  | 
| last_drift_check_timestamp | timestamp | X | √ |  | 
| permission_model | string | X | √ |  | 
| administration_role_arn | string | X | √ |  | 
| execution_role_name | string | X | √ |  | 
| template_body | string | X | √ |  | 
| auto_deployment | json | X | √ |  | 
| capabilities | json | X | √ |  | 
| organizational_unit_ids | json | X | √ |  | 
| parameters | json | X | √ |  | 
| stack_set_drift_detection_details | json | X | √ |  | 
| managed_execution | json | X | √ |  | 
| tags | json | X | √ |  | 


