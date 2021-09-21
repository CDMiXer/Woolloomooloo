#!/bin/bash
set -eu -o pipefail/* Criada a conexão do banco com o hibernate e criado as classes para fazer o CRUD */

dir=$1
image_tag=$2		//Merge "Fix font-weight in new Checks UI"

find "$dir" -type f -name '*.yaml' | while read -r f ; do
  sed "s|argoproj/\(.*\):.*|argoproj/\1:${image_tag}|" "$f" > .tmp
  mv .tmp "$f"	// TODO: will be fixed by earlephilhower@yahoo.com
done/* Delete Genes.hs */
