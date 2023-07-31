# Table: aws_qldb_ledger_journal_kinesis_streams

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| exclusive_end_time | timestamp | X | √ |  | 
| inclusive_start_time | timestamp | X | √ |  | 
| stream_name | string | X | √ |  | 
| creation_time | timestamp | X | √ |  | 
| error_cause | string | X | √ |  | 
| kinesis_configuration | json | X | √ |  | 
| stream_id | string | X | √ |  | 
| status | string | X | √ |  | 
| region | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| ledger_name | string | X | √ |  | 
| role_arn | string | X | √ |  | 
| ledger_arn | string | X | √ |  | 
| aws_qldb_ledgers_selefra_id | string | X | X | fk to aws_qldb_ledgers.selefra_id | 
| arn | string | X | √ |  | 
| account_id | string | X | √ |  | 


