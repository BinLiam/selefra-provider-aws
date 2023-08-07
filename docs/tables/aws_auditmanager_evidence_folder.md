# Table: aws_auditmanager_evidence_folder

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
| aws_auditmanager_assessment_selefra_id | string | X | X | fk to aws_s3_buckets.selefra_id | 
| name | string | X | √ |  | 
| id | string | √ | √ |  | 
| arn | string | X | √ |  | 
| assessment_id | string | X | √ |  | 
| control_set_id | string | X | √ |  | 
| assessment_report_selection_count | big_int | X | √ |  | 
| author | string | X | √ |  | 
| control_id | string | X | √ |  | 
| control_name | string | X | √ |  | 
| data_source | string | X | √ |  | 
| date | timestamp | X | √ |  | 
| evidence_aws_service_source_count | big_int | X | √ |  | 
| evidence_by_type_compliance_check_count | big_int | X | √ |  | 
| evidence_by_type_compliance_check_issues_count | big_int | X | √ |  | 
| evidence_by_type_configuration_data_count | big_int | X | √ |  | 
| evidence_by_type_manual_count | big_int | X | √ |  | 
| evidence_by_type_user_activity_count | big_int | X | √ |  | 
| evidence_resources_included_count | big_int | X | √ |  | 
| total_evidence | big_int | X | √ |  | 


