#!/usr/bin/env bash

k8=$GOPATH/src/k8s.io
code_gen=${k8}/code-generator
apimachinery=${k8}/apimachinery
clientgo=${k8}/client-go
kubernetes=${k8}/kubernetes

apimachinery_branch="release-1.11"
clientgo_branch="release-8.0"
code_gen_branch="release-1.11"


if [ ! -d ${apimachinery} ]; then
  cd ${k8} && git clone git@github.com:kubernetes/apimachinery.git && cd ${apimachinery}
else
  cd ${apimachinery} && git fetch origin
fi

git checkout ${apimachinery_branch}

if [ ! -d ${clientgo} ]; then
  cd ${k8} && git clone git@github.com:kubernetes/client-go.git && cd ${clientgo}
else
  cd ${clientgo} && git fetch origin
fi

git checkout ${clientgo_branch}

if [ ! -d ${kubernetes} ]; then
   cd $k8 && git clone --depth=1 git@github.com:kubernetes/kubernetes
else
  cd ${kubernetes} && git checkout master && git pull origin master
fi

if [ ! -d ${code_gen} ]; then
   mkdir -p ${k8} && cd $k8 && git clone git@github.com:kubernetes/code-generator.git && cd ${code_gen}
else
  cd ${code_gen} && git fetch origin
fi

git checkout ${code_gen_branch}

./generate-internal-groups.sh all  github.com/aerogear/mobile-crd-client/pkg/client/mobile github.com/aerogear/mobile-crd-client/pkg/apis github.com/aerogear/mobile-crd-client/pkg/apis  "mobile:v1alpha1"
./generate-internal-groups.sh client github.com/aerogear/mobile-crd-client/pkg/client/servicecatalog github.com/aerogear/mobile-crd-client/pkg/apis github.com/aerogear/mobile-crd-client/pkg/apis  "servicecatalog:v1beta1"
