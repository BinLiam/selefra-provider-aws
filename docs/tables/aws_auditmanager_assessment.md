# Table: aws_auditmanager_assessment

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
| id | string | √ | √ |  | 
| name | string | X | √ |  | 
| arn | string | X | √ |  | 
| status | string | X | √ |  | 
| compliance_type | string | X | √ |  | 
| assessment_report_destination | string | X | √ |  | 
| assessment_report_destination_type | string | X | √ |  | 
| creation_time | timestamp | X | √ |  | 
| description | string | X | √ |  | 
| last_updated | timestamp | X | √ |  | 
| delegations | json | X | √ |  | 
| framework | json | X | √ |  | 
| scope | json | X | √ |  | 
| roles | json | X | √ |  | 


