# Functional Go collections
[![CircleCI](https://circleci.com/gh/sghaida/fp.svg?style=svg)](https://circleci.com/gh/sghaida/fp)
![](coverage_badge.png)


this is my experiment to bridge the gap between Functional Programing languages (FP) and Go as it seems a good candidate for FP at the first glance so lets see how far this goes

## Helpers commands
1. download the dependencies
    >make init
    
2. run the tests
   >make test
                    
3. generate test coverage badge
   >gopherbadger -md="README.md"
                                   
4. run test coverage and generate coverage report in html page
    >make coverage

## Usage
 some basic HOWTO regarding the following functionality
### Options
you can check [Options Test Data](./src/testdata/options.go) and [Options Test](src/options/options_test.go)
### Either
you can check [Either Test Data](./src/testdata/either.go) and [Either Test](src/Either/either_test.go)
