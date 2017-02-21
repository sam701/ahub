# ahub
Github goodies

## Installation

### Install from source
```sh
git clone https://github.com/sam701/ahub.git $GOPATH/src/github.com/sam701/ahub
cd $GOPATH/src/github.com/sam701/ahub
glide install
go install

# test it
$GOPATH/bin/ahub
```

## Configure
Generate a Github access key, as described [here](https://help.github.com/articles/creating-an-access-token-for-command-line-use/),
make sure it has `Full control of private repositories` permission
and export it 
```sh
export GITHUB_TOKEN=<token>
```

