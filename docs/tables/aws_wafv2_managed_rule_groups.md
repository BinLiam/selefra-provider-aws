# Table: aws_wafv2_managed_rule_groups

## Primary Keys 

```
account_id, region, scope
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| versioning_supported | bool | X | √ |  | 
| account_id | string | X | √ |  | 
| scope | string | X | √ |  | 
| description | string | X | √ |  | 
| vendor_name | string | X | √ |  | 
| region | string | X | √ |  | 
| properties | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| name | string | X | √ |  | 


