#!/bin/bash

# Parameter:
# actualbranch ($1), targetbranch ($2), mainbranch ($3), 
# devbranch ($4), FEATURE_TAG_IN_BRANCH ($5), SUBFEATURE_TAG_IN_BRANCH ($6), 
# BUG_TAG_IN_BRANCH ($7), VERSION_MAIN_MAJOR ($8), VERSION_MAIN_MINOR ($9), 
# VERSION_MAIN_Patch ($10), VERSION_MAIN_BUILD ($11), VERSION_DEV_MAJOR ($12), 
# VERSION_DEV_MINOR ($13), VERSION_DEV_Patch ($14), VERSION_DEV_BUILD ($15), 
# VERSION_DEV_NR ($16), build_now ($17)

actualbranch=${1}
targetbranch=${2}
mainbranch=${3}
devbranch=${4}
FEATURE_TAG_IN_BRANCH=${5}
SUBFEATURE_TAG_IN_BRANCH=${6}
BUG_TAG_IN_BRANCH=${7}
VERSION_MAIN_MAJOR=${8}
VERSION_MAIN_MINOR=${9}
VERSION_MAIN_Patch=${10}
VERSION_MAIN_BUILD=${11}
VERSION_DEV_MAJOR=${12}
VERSION_DEV_MINOR=${13}
VERSION_DEV_Patch=${14}
VERSION_DEV_BUILD=${15}
VERSION_DEV_NR=${16}
build_now=${17}

if [[ "$targetbranch" == "$mainbranch" ]]; then
    if [[ "$actualbranch" == "$devbranch" ]]; then
        VERSION_MAIN_MAJOR=$VERSION_DEV_MAJOR
        VERSION_MAIN_MINOR=$VERSION_DEV_MINOR
        VERSION_MAIN_Patch=$VERSION_DEV_Patch
        VERSION_MAIN_BUILD=$build_now

        VERSION_DEV_NR=0
    fi
elif [[ "$targetbranch" == "$devbranch" ]]; then
    VERSION_DEV_NR=$((VERSION_DEV_NR+1))
    VERSION_DEV_BUILD=$build_now

    if [[ "$actualbranch" == "$FEATURE_TAG_IN_BRANCH"* || "$actualbranch" == "$SUBFEATURE_TAG_IN_BRANCH"* ]]; then
        if [[ $VERSION_MAIN_MAJOR -eq $VERSION_DEV_MAJOR && $VERSION_MAIN_MINOR -eq $VERSION_DEV_MINOR && $VERSION_MAIN_Patch -eq $VERSION_DEV_Patch ]]; then
            VERSION_DEV_MINOR=$((VERSION_DEV_MINOR+1))
        elif [[ $VERSION_MAIN_MAJOR -eq $VERSION_DEV_MAJOR && $VERSION_MAIN_MINOR -eq $VERSION_DEV_MINOR && $VERSION_MAIN_Patch -ne $VERSION_DEV_Patch ]]; then
            VERSION_DEV_MINOR=$((VERSION_DEV_MINOR+1))
            VERSION_DEV_Patch=0
        fi
    elif [["$actualbranch" == "$BUG_TAG_IN_BRANCH"* ]]; then
        if [[ $VERSION_MAIN_MAJOR -eq $VERSION_DEV_MAJOR && $VERSION_MAIN_MINOR -eq $VERSION_DEV_MINOR && $VERSION_MAIN_Patch -eq $VERSION_DEV_Patch ]]; then
            VERSION_DEV_Patch=$((VERSION_DEV_Patch+1))
        fi
    fi
fi

echo "VERSION_MAIN_MAJOR="$VERSION_MAIN_MAJOR""
echo "VERSION_MAIN_MINOR="$VERSION_MAIN_MINOR""
echo "VERSION_MAIN_Patch="$VERSION_MAIN_Patch""
echo "VERSION_MAIN_BUILD="$VERSION_MAIN_BUILD""
echo "VERSION_DEV_MAJOR="$VERSION_DEV_MAJOR""
echo "VERSION_DEV_MINOR="$VERSION_DEV_MINOR""
echo "VERSION_DEV_Patch="$VERSION_DEV_Patch""
echo "VERSION_DEV_BUILD="$VERSION_DEV_BUILD""
echo "VERSION_DEV_NR="$VERSION_DEV_NR""