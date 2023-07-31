# Table: aws_ssm_instance_compliance_items

## Primary Keys 

```
id, instance_arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| execution_summary | json | X | √ |  | 
| resource_type | string | X | √ |  | 
| severity | string | X | √ |  | 
| compliance_type | string | X | √ |  | 
| id | string | X | √ |  | 
| resource_id | string | X | √ |  | 
| status | string | X | √ |  | 
| region | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| title | string | X | √ |  | 
| account_id | string | X | √ |  | 
| instance_arn | string | X | √ |  | 
| details | json | X | √ |  | 
| aws_ssm_instances_selefra_id | string | X | X | fk to aws_ssm_instances.selefra_id | 


