# Information
```
Usage of pprof_collector:

Description:
  Collects the pprof profile from `net/http/pprof` endpoints and send it to a AWS Bucket

Options:
      --application_name string   Application name to identify the current profile (default "undefined")
  -u, --aws_access_key string     [required] AWS Access Key
  -p, --aws_secret_key string     [required] AWS Secret Key
  -b, --bucket string             [required] S3 Bucket name
      --cooldown int              Cool down period in seconds (default 5)
  -o, --output string             Output directory (default ".")
  -r, --region string             [required] Region for the S3 Bucket
  -s, --seconds int               Polling interval in seconds (default 30)
      --url string                Url for pprof profile (default "http://localhost/debug/pprof/profile")

```