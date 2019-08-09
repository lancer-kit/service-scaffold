#!/usr/bin/env bash
set -e

# !!! DO NOT REMOVE !!!
# FIX of "sed: RE error: illegal byte sequence"
export LC_CTYPE=C
export LANG=C

scaffold_path="github.com\/lancer-kit\/service-scaffold"
VCS_DOMAIN=""
VCS_USER=""
PROJECT_NAME=""
FULL_PATH=""

init_domain() {
  def_VCS_DOMAIN="github.com"
  read -r -p "Enter VCS domain (default: ${def_VCS_DOMAIN}): " VCS_DOMAIN

  if [[ ${VCS_DOMAIN} = "" ]]; then
      VCS_DOMAIN=${def_VCS_DOMAIN}
  fi
}


init_user() {
  read -r -p "Enter VCS username or group: " VCS_USER

  if [[ ${VCS_USER} = "" ]]; then
      init_user
  fi
}

init_project_name() {
  read -r -p "Enter project name: " PROJECT_NAME

  if [[ ${PROJECT_NAME} = "" ]]; then
      init_project_name
  fi
}

init_values() {
  init_domain
  init_user
  init_project_name

  FULL_PATH="${GOPATH}/src/${VCS_DOMAIN}/${VCS_USER}/${PROJECT_NAME}"
  echo "Path for new project: ${FULL_PATH}"
  read -r -p  "Is it correct? [Y/n]: " is_ok

  case ${is_ok} in
    "y"|"Y")
      return
      ;;
    *)
      init_values
      ;;
  esac
}

replace_imports() {
  cd "${FULL_PATH}"/
  find . -path ./vendor -prune -o  -type f -exec  sed -i '' -e "s/${scaffold_path}/${VCS_DOMAIN}\/${VCS_USER}\/${PROJECT_NAME}/g" {} +
}

copy_project() {
#  orig_dir=`pwd`
  mkdir -p "${FULL_PATH}"/
  cp -rf ./ "${FULL_PATH}"/

  cd "${FULL_PATH}"/

  rm -rf ./.git
  rm -rf ./.idea
  rm -rf ./init.sh
  rm -rf ./LICENSE
  mv ./tmpl.README.md ./README.md
# fixme
#  mv ./.gitlab-ci.yml.tmpl ./.gitlab-ci.yml
}


git_init() {
  cd "${FULL_PATH}"/

  git init
  git remote add origin "https://${VCS_DOMAIN}/${VCS_USER}/${PROJECT_NAME}.git"
  git add .
}

main() {
  init_values
  copy_project
  replace_imports
  git_init
}

main
