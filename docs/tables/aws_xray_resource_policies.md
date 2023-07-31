# Table: aws_xray_resource_policies

## Primary Keys 

```
account_id, region, policy_name, policy_revision_id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| last_updated_time | timestamp | X | √ |  | 
| policy_document | string | X | √ |  | 
| policy_name | string | X | √ |  | 
| policy_revision_id | string | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


