# echttpserver
http server supporting RESTful api of eccrawler

# Deployment

### Download [release](https://github.com/SpeculationFund/echttpapi/releases)

### Unzip, Config, Install
```
unzip echttpserver_amd64-YYYYMMDD.zip
cd .Build 
./echttpserver
```


# Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. 

### Prerequisites

Install golang  `sudo apt install golang-go`


### Installing

```
git clone https://github.com/SpeculationFund/echttpserver.git
cd echttpserver
go build
./echttpserver
```

### Running the tests

```
go test
```

### Documentation
* [GoDoc of dbTool](https://godoc.org/github.com/SpeculationFund/echttpserver)
* [Development docs of dbTool](https://github.com/SpeculationFund/echttpserver/wiki)


### Build

```
cd github.com/SpeculationFund/echttpserver
./build_release.sh
```
The output is `echttpserver-yyyymmdd.zip`, details are shown as follow,

```
Archive:  echttpserver-20171228.zip
  Length      Date    Time    Name
---------  ---------- -----   ----
        0  2017-12-28 20:04   .Build/
  8490552  2017-12-28 20:04   .Build/echttpserver_amd64-20171228
  6791432  2017-12-28 20:04   .Build/echttpserver_386-20171228
  6798752  2017-12-28 20:04   .Build/echttpserver_arm-20171228
---------                     -------
 22080736                     4 files

``` 

# Logistics

### Contributing

Please read [CONTRIBUTING.md](https://github.com/SpeculationFund/echttpserver/blob/master/CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull requests to us.

### Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the tags on this repository

### Authors

* **PhoenixAndMachine** - *Initial work* - [PhoenixAndMachine](https://github.com/PhoenixAndMachine)

See also the list of [contributors](https://github.com/your/project/contributors) who participated in this project.

### Acknowledgments

* Hat tip to anyone who's code was used
* Inspiration
* etc


### License

This project is licensed under the MIT License - see the [LICENSE.md](https://gist.github.com/Brownyuan/0b754b6009b7a4257bde9d1a23586678) file for details


