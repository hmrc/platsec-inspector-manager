
### Platsec-Inspector-Manager
PlatSec Inspector Manager is Golang command line tool that enables PlatSec Team members 
to create suppression rules in AWS Inspector. 

The requirements and in depth specification of the tool can be found
https://docs.google.com/document/d/1blVzVRaNws7LxWAf7hbWwyGgqma5_cnjAdY31VgqX6Y/edit#

Platform Teams can then apply the suppression rules created by PlatSec to Inspector 
findings and filter out what is not relevant to them depending on their use case. 

A suppression rule can be made up of multiple filters such as AWS Account, CVE, Type.

The command below is to add suppression rules 

./inspector -profile webops-users -account 638924580364 -filter-name test3 -filter-type account -username donald.duck 
-mfa-token 2342342

Where 
 - profile is the aws profile specifying that you have permissions to assume the role such as the webops account
 - account is the auth account
 - filter-name is the name of the filter you want to add
 - username is the username that is part of the serial number
 - filter-type can be cve, awsAccounts or typeCategory
 - mfa-token is the token for mfa 

### Setup
There is a config.yml file within the root of the project that has a section entitled aws
You must enter the account number for account setting where you want the suppression rule to be added (Cloudtrail account).
You must also supply the rolename that has permissions to add the suppression rule normally this is RoleSecurityAdministrator.

You must create the config.yml file in the project root with the following content

aws:
    account: "<AWS Account number where rule to be applied>"
    rolename: "<Name of role to assume>"

You must have a profile setup for your webops account.
### Version 1.0.0
### Release History
| Version Number | Name             | Description                                                                                    | Status   | Date       |
|----------------|------------------|------------------------------------------------------------------------------------------------|----------|------------|
| 1.0.0          | Initial Release  | Contains the functional requirements R-01 and R-04 as detailed in the specification  document. | Released | 22/06/2022 |
| 1.0.1          | Storage          | Implementation of storage for suppression rules R-06                                           | Proposed |            |
| 1.0.2          | Rules Management | Implementation of rules management R-05, R-07, R-08                                            | Proposed |            |
| 1.0.3          | Reporting        | Implementation of reporting R-09                                                               | Proposed |            |
### License

This code is open source software licensed under the [Apache 2.0 License]("http://www.apache.org/licenses/LICENSE-2.0.html").