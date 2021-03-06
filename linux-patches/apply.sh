#!/bin/bash
# Copyright 2018 u-root Authors
#
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file

dir=$(dirname $0)

for i in $(ls ${dir}/*.patch)
do
  echo "=> Applying ${i}"
  git apply $i
done
