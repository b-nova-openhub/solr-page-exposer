![sopagex-header](https://github.com/b-nova-openhub/jams-vanilla-content/raw/main/gosopagex-header.png)

# gosopagex (go-solr-page-exposer)

<!--- These are examples. See https://shields.io for others or to customize this set of shields. You might want to include dependencies, project status and licence info here --->
![GoDoc](https://godoc.org/github.com/go-git/go-git/v5?status.svg)
![GitHub repo size](https://img.shields.io/github/repo-size/b-nova-openhub/solr-page-exposer)
![GitHub contributors](https://img.shields.io/github/contributors/b-nova-openhub/solr-page-exposer)
![GitHub stars](https://img.shields.io/github/stars/b-nova-openhub/solr-page-exposer?style=social)
![GitHub forks](https://img.shields.io/github/forks/b-nova-openhub/solr-page-exposer?style=social)
![Twitter Follow](https://img.shields.io/twitter/follow/b__nova?style=social)

gosopagex is a `tool` that allows `content maintainers` to do `expose html content to the Apache Solr search platform`.

More specifically, gosopagex is microservice written in Go that maintains html files in a desired solr-friendly format and exposes them RESTful interface. It is operates on the principles of KISS and is best integrated in a JAMstack.

## Prerequisites

Before you begin, ensure you have met the following requirements:
<!--- These are just example requirements. Add, duplicate or remove as required --->

* You have installed the latest version of `go1.16.5`
* You have a `Linux/Mac OS` machine with working knowledge of the underlying filesystem and Go build process.

## Installing gosopagex

If you want to try out on your local machine, you can simply install gosopagex system-wide with the following command:

```
$ ❯  go install -v ./...
```

Alternatively, you can build a binary and run that one with a configurable shell script. See how to **build stapagen
with the Makefile** further down below (#Makefile).

## Running gosopagex

To run gosopagex, follow these steps:

```
$ ❯  sopagex
```

If you have build stapagen with the Makefile, you can simply adjust the parameters in the shell script `stapagen.sh` and
then launch the script.

```
$ ❯  ./sopagex.sh
```

This method is preferable as you're more flexible and don't have to retype the binary arguments everytime.

## Reading gosopagex's API

Once gosopagex is running, it does clone the git repository, converts its markdown files to html pages and forwards
them over a RESTful api. That api can be accessed over locally if the executable was run on your local
machine (`http://localhost:8080`) or over a K8s service definition.

## Makefile

There is a Makefile which automates the building process of the stapagen application. The Makefile has 6 build targets:
test, vet, fmt, mod, build, run, install and all. It can simply be run as such:

```
$ ❯  make build
```

## Dockerfile

There is also a Dockerfile by which one can containerize the stapagen application. The port that is being exposed is the
default 8080.

## Deployment to K8s

{coming soon}

### Quick Deployment to DigitalOcean's Kubernetes

[![Deploy to DO](https://www.deploytodo.com/do-btn-blue.svg)](https://cloud.digitalocean.com/apps/new?repo=https://github.com/b-nova-openhub/solr-page-exposer/tree/main)

## Contributing to gosopagex

<!--- If your README is long or you have some specific process or steps you want contributors to follow, consider creating a separate CONTRIBUTING.md file--->
To contribute to gosopagex, follow these steps:

1. Fork this repository.
2. Create a branch: `git checkout -b <branch_name>`.
3. Make your changes and commit them: `git commit -m '<commit_message>'`
4. Push to the original branch: `git push origin <project_name>/<location>`
5. Create the pull request.

Alternatively see the GitHub documentation
on [creating a pull request](https://help.github.com/en/github/collaborating-with-issues-and-pull-requests/creating-a-pull-request)
.

## Contributors

Thanks to the following people who have contributed to this project:

* [@raffaelschneider](https://github.com/raffaelschneider) 💻📖🐛
* [@stefanwelsch](https://github.com/bnova-stefan) 💻🧑‍🏫
* [@tomtrapp](https://github.com/tomtrapp) 🤔👀
* [@rickyelfner](https://github.com/ricky-bnova) 💬🐛

You might want to consider using something like
the [All Contributors](https://github.com/all-contributors/all-contributors) specification and
its [emoji key](https://allcontributors.org/docs/en/emoji-key).

## Contact

If you want to contact me you can reach me at [info@b-nova.com](info@b-nova.com).

## License

<!--- If you're not sure which open license to use see https://choosealicense.com/--->

This project uses the following license: [MIT License](https://opensource.org/licenses/MIT)
. https://opensource.org/licenses/MIT
