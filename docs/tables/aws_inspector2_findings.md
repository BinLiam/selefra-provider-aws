# Table: aws_inspector2_findings

## Primary Keys 

```
arn
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| aws_account_id | string | X | √ |  | 
| last_observed_at | timestamp | X | √ |  | 
| inspector_score | float | X | √ |  | 
| package_vulnerability_details | json | X | √ |  | 
| title | string | X | √ |  | 
| account_id | string | X | √ |  | 
| arn | string | √ | √ |  | 
| resources | json | X | √ |  | 
| severity | string | X | √ |  | 
| status | string | X | √ |  | 
| exploit_available | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| description | string | X | √ |  | 
| type | string | X | √ |  | 
| inspector_score_details | json | X | √ |  | 
| network_reachability_details | json | X | √ |  | 
| region | string | X | √ |  | 
| finding_arn | string | X | √ |  | 
| first_observed_at | timestamp | X | √ |  | 
| remediation | json | X | √ |  | 
| exploitability_details | json | X | √ |  | 
| fix_available | string | X | √ |  | 
| updated_at | timestamp | X | √ |  | 


