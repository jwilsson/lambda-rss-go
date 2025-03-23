# lambda-rss
Go AWS Lambda function for scraping sites not offering RSS feeds and publish them as RSS feeds.

## Setup
1. Start by [bootstraping your AWS environment](https://docs.aws.amazon.com/cdk/v2/guide/bootstrapping-env.html).
2. [Setup a new policy](https://stackoverflow.com/questions/57118082/what-iam-permissions-are-needed-to-use-cdk-deploy/61102280#61102280) granting permission to assume CDK roles.
3. Add the following GitHub Actions secrets: `AWS_ACCESS_KEY_ID`, `AWS_SECRET_ACCESS_KEY`, and `AWS_DEFAULT_REGION`.
4. Profit.

## Usage
Each `.go`-file in `app/sites` contains the logic to extract content from a URL and expose it as an RSS feed. Each feed can then be accessed via https://my-function-url/?s=site-name.
