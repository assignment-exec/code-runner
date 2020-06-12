# Contributing to Code Runner

First of all, THANK YOU for contributing to the project. When contributing, please follow the below guidelines.

## Submit Changes or New Features

### Clone the source code
- To make changes to [code-runner](https://github.com/assignment-exec/code-runner) you need to fork the repository and clone it.
- Use the below commands to clone the forked repository using HTTPS and change the working directory.
```commandline
$ git clone <https_clone_url>
$ cd code-runner  # Use chdir if working on windows.
```

### Branching
Each change must be made in a separate branch, Branch out of the `develop` branch using the command below.
```commandline
git checkout -b <branch_name> develop
```

### Commit changes
- We like descriptive commit messages. When writing commit messages, make sure to follow the [git commit guidelines](https://git-scm.com/book/en/v2/Distributed-Git-Contributing-to-a-Project).
- It would be beneficial to have the commits atomic to ensure that the project has a cleaner history.

### Add Unit tests
- For every change unit test cases should be added if possible. All the tests should pass before submitting the change for review.
- See [details](https://golang.org/pkg/testing/) for writing unit tests in Go.
- Once the unit tests have been added, run the below command to make sure that all the tests pass.
```commandline
make test
```

### Send changes for review
- Once the changes are ready send a pull request to the `develop` branch of [code-runner](https://github.com/assignment-exec/code-runner/pull/new/develop) with a list of things you've done.
- See [pull requests](http://help.github.com/pull-requests/) to know more about it. 


Once again, we THANK YOU for taking the time to contribute !