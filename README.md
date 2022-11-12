## CSIbuilder

CSIbuilder is a SDK for building Kubernetes CSI Driver.

Similar to Kubebuilder, CSIbuilder does **not** exist as an example to *copy-paste*, but instead provides powerful
libraries and tools to simplify building and publishing Kubernetes CSI Driver from scratch.

### Installation

It is strongly recommended that you use a released version. Release binaries are available on
the [releases](https://github.com/zwwhdls/csibuilder/releases) page.

Install CSIbuilder:

```bash
# download csibuilder and install locally.
curl -L -o csibuilder.tar https://github.com/zwwhdls/csibuilder/releases/download/v0.1.0/csibuilder-darwin-amd64.tar
tar -zxvf csibuilder.tar  && chmod +x csibuilder && mv csibuilder /usr/local/bin/
```

### Getting Started

Only two steps to create a CSI Driver project:

1. initialize a project

 ```bash
 # init your project
 export GO111MODULE=on
 mkdir $GOPATH/src/csi-hdls
 cd $GOPATH/src/csi-hdls
 # init csi repo
 csibuilder init --repo hdls --owner "zwwhdls"
 ```

2. create a csi driver

 ```bash
 # create csi
 csibuilder create --csi hdls
 ```

