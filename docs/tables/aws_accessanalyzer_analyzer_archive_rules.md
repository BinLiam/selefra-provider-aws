# Table: aws_accessanalyzer_analyzer_archive_rules

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| aws_accessanalyzer_analyzers_selefra_id | string | X | X | fk to aws_accessanalyzer_analyzers.selefra_id | 
| created_at | timestamp | X | √ |  | 
| updated_at | timestamp | X | √ |  | 
| region | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| filter | json | X | √ |  | 
| rule_name | string | X | √ |  | 
| account_id | string | X | √ |  | 
| analyzer_arn | string | X | √ |  | 


