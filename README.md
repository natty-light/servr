### ServeR

This project serves as a demonstration of a full stack `R Markdown` report generation service

The architecture is as follows:

The reports are generated automatically using the body of a HTTP request at the `/api/generate` endpoint. the report is then uploaded to an S3 bucket, and the S3 URL is returned in the response

There will be a client side application to handle requesting the reports

Authorization is done via JWT

The addition of a cache for checking for generated reports, as well as a database for storing user data is planned