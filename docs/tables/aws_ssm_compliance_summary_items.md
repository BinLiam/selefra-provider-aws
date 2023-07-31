# Table: aws_ssm_compliance_summary_items

## Primary Keys 

```
account_id, region, compliance_type
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| compliant_summary | json | X | √ |  | 
| non_compliant_summary | json | X | √ |  | 
| account_id | string | X | √ |  | 
| region | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| compliance_type | string | X | √ |  | 


