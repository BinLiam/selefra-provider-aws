# Table: aws_securityhub_findings

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| network | json | X | √ |  | 
| vulnerabilities | json | X | √ |  | 
| request_region | string | X | √ |  | 
| finding_provider_fields | json | X | √ |  | 
| malware | json | X | √ |  | 
| confidence | big_int | X | √ |  | 
| note | json | X | √ |  | 
| patch_summary | json | X | √ |  | 
| record_state | string | X | √ |  | 
| user_defined_fields | json | X | √ |  | 
| request_account_id | string | X | √ |  | 
| description | string | X | √ |  | 
| updated_at | string | X | √ |  | 
| threat_intel_indicators | json | X | √ |  | 
| company_name | string | X | √ |  | 
| compliance | json | X | √ |  | 
| source_url | string | X | √ |  | 
| verification_state | string | X | √ |  | 
| network_path | json | X | √ |  | 
| severity | json | X | √ |  | 
| resources | json | X | √ |  | 
| action | json | X | √ |  | 
| first_observed_at | string | X | √ |  | 
| process | json | X | √ |  | 
| aws_account_id | string | X | √ |  | 
| generator_id | string | X | √ |  | 
| region | string | X | √ |  | 
| remediation | json | X | √ |  | 
| types | string_array | X | √ |  | 
| workflow | json | X | √ |  | 
| product_arn | string | X | √ |  | 
| schema_version | string | X | √ |  | 
| product_fields | json | X | √ |  | 
| sample | bool | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| created_at | string | X | √ |  | 
| title | string | X | √ |  | 
| last_observed_at | string | X | √ |  | 
| product_name | string | X | √ |  | 
| related_findings | json | X | √ |  | 
| threats | json | X | √ |  | 
| workflow_state | string | X | √ |  | 
| id | string | X | √ |  | 
| criticality | big_int | X | √ |  | 


