# terraform-null-echo

![Release Version](https://img.shields.io/github/v/release/mundoalem/terraform-null-echo)
![Pipeline Status](https://github.com/mundoalem/terraform-null-echo/actions/workflows/pipeline.yml/badge.svg)
![Contributors](https://img.shields.io/github/contributors/mundoalem/terraform-null-echo)

A terraform module that echoes its input message to an output.

## Introduction

The `echo` module received a message as an input variable and just output the
same value as an output. This module is only meant to be used to test Terraform
related projects such as the *DevOps* templates developed by us.

You can see this module like the *hello world* exercise for programming
languages.

## Usage

```
module "echo" {
  source  = "mundoalem/echo/null"
  version = "1.0.2"
}
```

## License

[GPLv3](https://choosealicense.com/licenses/gpl-3.0/)

## Feedback

If you have any feedback, please open an [issue](https://github.com/mundoalem/terraform-null-echo/issues).

## Contributing

In particular, this community seeks the following types of contributions:

- Participate in an issue thread or start your own to have your voice heard.
- Help us ensure that this repository documentation is comprehensive.
- Implement new features to the project.
- Fix open issues.

The first step is to fork the project and clone its source code locally, then
please run:

```
$ go mod vendor
```

This will install all the Go dependencies we use to run the pipeline steps
locally. Those pipeline steps are implemented using [mage](https://magefile.org/),
it accepts the following targets:

| Target   | Description                                                                     |
| -------- | ------------------------------------------------------------------------------- |
| build    | Runs `terraform init` and `terraform plan` saving the plan in `build` directory |
| clean    | Removes the files and directories generated by the build step                   |
| lint     | Runs `terraform fmt` for `modules` and `examples` directories                   |
| release  | Runs `terraform apply` from inside the `examples`dir                            |
| reset    | Does the same as the `clear` target plus removes `vendor` directory             |
| scan     | Runs `tfsec` to scan the code for unsafe patterns                               |
| test     | Runs all tests in `test` directory using `terratest`                            |

## Conduct

We are committed to providing a friendly, safe and welcoming environment for
all, regardless of gender, sexual orientation, disability, ethnicity,
religion, income or similar personal characteristic.

Please be kind and courteous. There's no need to be mean or rude. Respect that
people have differences of opinion and that every design or implementation
choice carries a trade-off and numerous costs. There is seldom a right answer,
merely an optimal answer given a set of values and circumstances.

Please keep unstructured critique to a minimum. If you have solid ideas you
want to experiment with, make a fork and see how it works.

We will exclude you from interaction if you insult, demean or harass anyone.
That is not welcome behavior. We interpret the term "harassment" as including
the definition in the [Citizen Code](http://citizencodeofconduct.org/) of
Conduct; if you have any lack of clarity about what might be included in that
concept, please read their definition. In particular, we don't tolerate
behavior that excludes people in socially marginalized groups.

Whether you're a regular contributor or a newcomer, we care about making this
community a safe place for you and we've got your back.

Likewise any spamming, trolling, flaming, baiting or other attention-stealing
behavior is not welcome.

## Communication

GitHub issues are the primary way for communicating about specific proposed
changes to this project.

In both contexts, please follow the conduct guidelines above. Language issues
are often contentious and we'd like to keep discussion brief, civil and focused
on what we're actually doing, not wandering off into too much imaginary stuff.
