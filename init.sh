#!/usr/bin/env bash
set -e
scaffold_path="gitlab.inn4science.com\/gophers\/service-scaffold"
VCS_DOMAIN=""
VCS_USER=""
PROJECT_NAME=""
FULL_PATH=""

init_domain() {
  def_VCS_DOMAIN="gitlab.inn4science.com"
  read -p "Enter VCS domain (default: ${def_VCS_DOMAIN}): " VCS_DOMAIN

  if [[ ${VCS_DOMAIN} = "" ]]; then
      VCS_DOMAIN=${def_VCS_DOMAIN}
  fi
}


init_user() {
  read -p "Enter VCS username or group: " VCS_USER

  if [[ ${VCS_USER} = "" ]]; then
      init_user
  fi
}

init_project_name() {
  read -p "Enter project name: " PROJECT_NAME

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
  read -p  "Is it correct? [Y/n]: " is_ok

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
  cd ${FULL_PATH}/
  find . -path ./vendor -prune -o  -type f -exec  sed -i '' -e "s/${scaffold_path}/${VCS_DOMAIN}\/${VCS_USER}\/${PROJECT_NAME}/g" {} +
}

copy_project() {
  orig_dir=`pwd`
  mkdir -p ${FULL_PATH}/
  shopt -s dotglob
  cp -rf ./* ${FULL_PATH}/

  cd ${FULL_PATH}/

  rm -rf ./.git
  rm -rf ./.idea
  rm -rf ./init.sh
  mv ./README.md.tmpl ./README.md
  mv ./.gitlab-ci.yml.tmpl ./.gitlab-ci.yml
}


git_init() {
  cd ${FULL_PATH}/

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