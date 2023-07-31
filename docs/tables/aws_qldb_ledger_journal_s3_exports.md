# Table: aws_qldb_ledger_journal_s3_exports

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| s3_export_configuration | json | X | √ |  | 
| account_id | string | X | √ |  | 
| ledger_arn | string | X | √ |  | 
| exclusive_end_time | timestamp | X | √ |  | 
| export_creation_time | timestamp | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| inclusive_start_time | timestamp | X | √ |  | 
| status | string | X | √ |  | 
| region | string | X | √ |  | 
| export_id | string | X | √ |  | 
| ledger_name | string | X | √ |  | 
| role_arn | string | X | √ |  | 
| output_format | string | X | √ |  | 
| aws_qldb_ledgers_selefra_id | string | X | X | fk to aws_qldb_ledgers.selefra_id | 


