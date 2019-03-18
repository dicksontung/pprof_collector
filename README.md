# Information
```
Usage of ./pprof_collector:

Description:
  Continuously collects the pprof profile from `net/http/pprof` endpoints and optionally sends it to a AWS Bucket

Options:
      --application_name string   Application name to identify the current profile (default "undefined")
  -u, --aws_access_key string     AWS Access Key
  -p, --aws_secret_key string     AWS Secret Key
  -b, --bucket string             S3 Bucket name
      --cooldown int              Cool down period in seconds (default 5)
  -o, --output string             Output directory (default "./tmp/pprof_collector")
  -r, --region string             Region for the S3 Bucket
  -s, --seconds int               Polling interval in seconds (default 30)
      --url string                Url for pprof profile (default "http://localhost/debug/pprof/profile")

```