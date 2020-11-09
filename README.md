# hdock

## A few goals for this project:

1. Make it easy to build and run docker images created from a Python repository. This would be kind of like [`repo2docker`](https://github.com/jupyterhub/repo2docker) but simpler and in Go. I wrote up this [doc](https://github.com/hdoupe/Tax-Brain/blob/add-dockerfile/run_docker.md) so that a friend could debug some dependency issues. I'm hoping this tool can handle that flow in a more intuitive way.

   - [ ] Build from an `environment.yaml` or `requirements.txt` file.
   - [ ] Run commands with volume mounted by default.
   - [ ] Support interactive mode and/or jupyter notebooks.

2. Learn some Go. One of the cool things about Go that is missing in languages that I usually use like Node or Python is that you can build an executable and run it wherever you want. The install instructions become a one-line bash command instead of links out to miniconda installation instructions.

## Usage

_(Right now this will only build the image!)_

```
➜  hdock run --help
Run a docker command mounted on some dir.

Usage:
  hdock run [buildPath] [flags]

Flags:
  -h, --help   help for run

Global Flags:
      --config string   config file (default is $HOME/.hdock.yaml)
      --tag string      name of docker image
```
