
### platsec-inspector-manager
PlatSec Inspector Manager is Golang command line tool that enables PlatSec Team members 
to create suppression rules in AWS Inspector. 

Platform Teams can then apply the suppression rules created by PlatSec to Inspector 
findings and filter out what is not relevant to them depending on their use case. 

A suppression rule can be made up of multiple filters such as AWS Account, CVE, finding Type (Network reachability / Package vulnerability).

### Example Usage

An example command to create a suppression rule is as follows:
```
./inspector -profile aws-profile -account 123456789101 -filter-name my-amazing-filter -filter-type a-filter-type -username your.name -mfa-token 123456

```

### License

This code is open source software licensed under the [Apache 2.0 License]("http://www.apache.org/licenses/LICENSE-2.0.html").