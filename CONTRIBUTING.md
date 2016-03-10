## Table of Contents

- [Naming Conventions](#naming-conventions)
- [Build Tools](#build-tools)
  - [Compiling](#compiling)
  - [Dependency Management](#dependency-management)
  - [Linting, Formating, etc](#linting-formating-etc)
- [Documentation](#documentation)
    - [Changelog](#changelog)
    - [Readme](#readme)
    - [Copyright And Licensing](#copyright-and-licensing)
    - [External Documentation](#external-documentation)
- [Configuring And Usage](#configuring-and-usage)
    - [Exit Codes](#exit-codes) 

## Naming Conventions

- Binaries should not have an extension unless they are designed to work on Windows
- Dashes(`-`) and Underscores(`_`) should not be used in filenames or in the directory structure, Golang gets really unhappy. If you have to use one for the name of a binary then please use dashes.
- camelCase is the name of the game at Yieldbot when it comes to golang, please use this format. You are free to use other methods but tooling will only be tested against that scheme.

## Build tools

### Compiling

`go build` is the tool used for building the binary. In the case of using `godep`, `godep go build` is the command executed. Please see the godep documentation for more information. This will be the preferred and currently, only supported way to handle vendoring at Yieldbot but you are free to use your own methods as long as they work for you

### Dependency Management

The standard way to mange them will be via Godep. This tool allows developers to easily pin and manage third party dependencies without having to fork them or copy them them into the repo manually.

The simplest way to use this tool is as follows:

1. `cd <PACKAGE_ROOT>` Ex. `cd sensupluginfile`
1. `make updatedeps` Bring in the needed dependencies to the src tree
1. `make dep_tree` Pull the necessary dependencies into the package and automatically update the src import statements to reflect this.

For more details on managing dependencies using this tool check out the [project repo](https://github.com/tools/godep).

To update a dependency there are several methods. Below is the safest:

1. `cd <PACKAGE_ROOT>` Ex. `cd sensupluginfile`
1. `go get -u <pkg>`
1. `godep update`

If you use the makefile to handle dependency management then it will update all dependencies to their latest version. To ensure this it will remove the existing `godeps` tree and replace it wholesale.

### Linting, Formating, etc

`gofmt` is enforced in the build pipeline and any errors will return the file that failed. You can correct them with `make format_correct` or `gofmt -w <file>`.

Linting and Vetting is not currently enforced but those tools are available and their usage is encouraged. If you would like to avoid having to remember to use these tools then you can add a pre-commit hook to your `.git/hooks` and that way it will get run before every commit. The below example will run `gofmt` and abort the commit if it fails.

```shell
#!/bin/sh

gofiles=$(git diff --cached --name-only --diff-filter=ACM | grep '.go$')
[ -z "$gofiles" ] && exit 0

unformatted=$(gofmt -l $gofiles)
[ -z "$unformatted" ] && exit 0

echo >&2 "Go files must be formatted with gofmt. Please run:"
for fn in $unformatted; do
  echo >&2 "  gofmt -w $PWD/$fn"
done

exit 1
```

## Documentation

There are many ways to document your code, Godoc is the preferred method with more indepth information going into a man page.

### CHANGELOG

You need it plain and simple, as with a [README](https://github.com/sensu-plugins/sensu-plugins-kubernetes) this is pretty basic. When updating the [CHANGELOG](https://github.com/sensu-plugins/sensu-plugins-aws/blob/master/CHANGELOG.md), please use the format outlined in [Keep A Changelog](http://keepachangelog.com/) as this is the standard way. A complete example is detailed below. While [semver](http://semver.org/) is usually a good idea this project does not adhere to it but the CHANGELOG and README should be considered authorative.

```markdown
#Change Log
This project adheres to [Semantic Versioning](http://semver.org/).

This CHANGELOG follows the format listed at [Keep A Changelog](http://keepachangelog.com/)

## [Unreleased]
### Added
- check-ebs-snapshots.rb: added -i flag to ignore volumes with an IGNORE_BACKUP tag
- check-sensu-client.rb Ensures that ec2 instances are registered with Sensu. 
- check-trustedadvisor-service-limits.rb: New check for service limits based on Trusted Advisor API 

## [2.1.1] - 2016-02-05
### Added
- check-ec2-cpu_balance.rb: scans for any t2 instances that are below a certain threshold of cpu credits
- check-instance-health.rb: adding ec2 instance health and event data

### Changed
- Update to aws-sdk 2.2.11 and aws-sdk-v1 1.66.0

### Fixed
- check-vpc-vpn.rb: fix execution error by running with aws-sdk-v1
- handler-ec2_node.rb: default values for ec2_states were ignored
- added new certs

## [2.1.0] - 2016-01-15
### Added
- check-elb-health-sdk.rb: add option for warning instead of critical when unhealthy instances are found
- check-rds.rb: add M4 instances
- handler-sns.rb: add option to use a template to render body mail
- check-rds-events.rb: add RDS event message to output
- Added check-cloudwatch-metric that checks the values of cloudwatch metrics
- Added check-beanstalk-elb-metric that checks an ELB used in a Beanstalk environment
- Added check-certificate-expiry that checks the expiration date of certificates loaded into IAM
- Added test cases for check-certificate-expiry.rb

### Changed
- handler-ec2_node.rb: Update to new API event naming and simplifying ec2_node_should_be_deleted method and fixing match that will work with any user state defined, also improved docs
- metrics-elb-full.rb: flush hash in-between iterations
- check-ses-limit.rb: move to AWS-SDK v2, use common module, return unknown on empty responses

### Fixed
- metrics-memcached.rb: Fixed default scheme
- Fix typo in cloudwatch comparison check

## [2.0.1] - 2015-11-03
### Changed
- pinned all dependencies
- set gemspec to require > `2.0.0`

Nothing new added, this is functionally identical to `2.0.0`. Doing a github release which for some reason failed even though a gem was built and pushed.

## [2.0.0] - 2015-11-02

WARNING: This release drops support for Ruby 1.9.3, which is EOL as of 2015-02.

### Added
- Added check-cloudwatch-alarm to get alarm status
- Added connection metric for check-rds.rb
- Added check-s3-bucket that checks S3 bucket existence
- Added check-s3-object that checks S3 object existence
- Added check-emr-cluster that checks EMR cluster existence
- Added check-vpc-vpn that checks the health of VPC VPN connections

### Fixed
- handler-ec2_node checks for state_reason being nil prior to code access
- Cosmetic fixes to metrics-elb, check-rds, and check-rds-events
- Return correct metrics values in check-elb-sum-requests

### Removed
- Removed Ruby 1.9.3 support

## [1.2.0] - 2015-08-04
### Added
- Added check-ec2-filter to compare filter results to given thresholds
- Added check-vpc-nameservers, which given a VPC will validate the name servers in the DHCP option set.

### Fixed
- handler-ec2_node accounts for an empty instances array

## [1.1.0] - 2015-07-24
### Added
- Added new AWS SES handler - handler-ses.rb
- Add metrics-ec2-filter to store node ids and count matching a given filter
- Check to alert on unlisted EIPs

## [1.0.0] - 2015-07-22

WARNING:  This release contains major breaking changes that will impact all users.  The flags used for access key and secret key have been standardized accross all plugins resulting in changed flags for the majority of plugins. The new flags are -a AWS_ACCESS_KEY and -k AWS_SECRET_KEY.

### Added
- EC2 node handler will now remove nodes terminated by a user
- Transitioned EC2 node handler from fog to aws sdk v2
- Allowed ignoring nil values returned from Cloudwatch in the check-rds plugin. Previously if Cloudwatch fell behind you would be alerted
- Added support for checking multiple ELB instances at once by passing a comma separated list of ELB instance names in metrics-elb-full.rb
- Added check-autoscaling-cpucredits.rb for checking T2 instances in autoscaling groups that are running low on CPU credits
- Updated the fog and aws-sdk gems to the latest versions to improve performance, reduce 3rd party gem dependencies, and add support for newer AWS features.
- Add metrics-ec2-filter to store node ids and count matching a given filter

### Fixed
- Renamed autoscaling-instance-count-metrics.rb -> metrics-autoscaling-instance-count.rb to match our naming scheme
- Reworked check-rds-events.rb to avoid the ABCSize warning from rubocop
- Corrected the list of plugins / files in the readme
- Make ELB name a required flag for the metrics ELB plugins to prevent nil class errors when it isn't provided
- Properly document that all plugins default to us-east-1 unless the region flag is passed
- Fix the ELB metrics plugins to properly use the passed auth data
- Fixed the metrics-elb-full plugin to still add the ELB instance name when a graphite schema is appended
- Fixed all plugins to support passing the AWS access and secret keys from shell variables. Plugin help listed this as an option for all plugins, but the support wasn't actually there.

## [0.0.4] - 2015-07-05
### Added
- Added the ability to alert on un-snapshotted ebs volumes

## [0.0.3] - 2015-06-26
### Fixed
- Access key and secret key should be optional
- Added 3XX metric collection to the ELB metrics plugins
- Fixed the metric type for SurgeQueueLength ELB metrics
- Fixed logic for ec2 instance event inclusion

## [0.0.2] - 2015-06-02
### Fixed
- added binstubs

### Changed
- removed cruft from /lib

## 0.0.1 - 2015-05-21
### Added
- initial release

[Unreleased]: https://github.com/sensu-plugins/sensu-plugins-aws/compare/2.1.1...HEAD
[2.1.1]: https://github.com/sensu-plugins/sensu-plugins-aws/compare/2.1.0...2.1.1
[2.1.0]: https://github.com/sensu-plugins/sensu-plugins-aws/compare/2.0.1...2.1.0
[2.0.1]: https://github.com/sensu-plugins/sensu-plugins-aws/compare/2.0.0...2.0.1
[2.0.0]: https://github.com/sensu-plugins/sensu-plugins-aws/compare/1.2.0...2.0.0
[1.2.0]: https://github.com/sensu-plugins/sensu-plugins-aws/compare/1.1.0...1.2.0
[1.1.0]: https://github.com/sensu-plugins/sensu-plugins-aws/compare/1.0.0...1.1.0
[1.0.0]: https://github.com/sensu-plugins/sensu-plugins-aws/compare/0.0.4...1.0.0
[0.0.4]: https://github.com/sensu-plugins/sensu-plugins-aws/compare/0.0.3...0.0.4
[0.0.3]: https://github.com/sensu-plugins/sensu-plugins-aws/compare/0.0.2...0.0.3
[0.0.2]: https://github.com/sensu-plugins/sensu-plugins-aws/compare/0.0.1...0.0.2
```

### README

It goes without saying that packages should have a README and that is should contain any relavent data. If any of the following applies to a package it should be noted:
- modifications to the build pipeline
- specific dependencies
- install or usage instructions
- contribution guidelines (if offically Open Sourced)

### Copyright and Licensing
The preferred license for all code associated with code that is to be released is the [MIT License](http://opensource.org/licenses/MIT).

### External Documentation

Godocs is the easiest method to generate docs but creating and building man and info pages is also supported via the Makefile. Please see the documentation of each of these for details on how to use them.

## Configuring and Usage

All golang binaries should adhere to the [12 Factor App](http://12factor.net/) where applicable. This is not currently enforced but is **strongly** encouraged. To assist in this please use the [Cobra](https://github.com/spf13/cobra) and [Viper](https://github.com/spf13/viper) tools as they will generate binary scaffolding that does conform to this standard. Concerning monitoring applications, The order of preference for configuing is as follows:

1. Environament Variables
1. Commandline Flags
1. [Viper](https://github.com/spf13/viper) Configuration File (dropped via Chef) 

### Exit Codes

All binaries are strongly encouraged to use these [error codes](https://github.com/yieldbot/sensuplugin/blob/master/sensuutil/common.go). You can use other ones but if possible please add them here. As monitoring progrresses there is functionality that will be able to do more fined-grained exclusions based upon error codes. To take full advantage of this it is encouraged that you use these.

