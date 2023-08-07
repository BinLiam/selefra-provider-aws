# Table: aws_auditmanager_evidence

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
| aws_auditmanager_evidence_folder_selefra_id | string | X | X | fk to aws_s3_buckets.selefra_id | 
| id | string | √ | √ |  | 
| arn | string | X | √ |  | 
| assessment_id | string | X | √ |  | 
| control_set_id | string | X | √ |  | 
| evidence_folder_id | string | X | √ |  | 
| assessment_report_selection | string | X | √ |  | 
| aws_account_id | string | X | √ |  | 
| aws_organization | string | X | √ |  | 
| compliance_check | string | X | √ |  | 
| data_source | string | X | √ |  | 
| event_name | string | X | √ |  | 
| event_source | string | X | √ |  | 
| evidence_aws_account_id | string | X | √ |  | 
| evidence_by_type | string | X | √ |  | 
| iam_id | string | X | √ |  | 
| time | timestamp | X | √ |  | 
| attributes | json | X | √ |  | 
| resources_included | json | X | √ |  | 


