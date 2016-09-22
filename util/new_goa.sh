#!/bin/bash

source $GOPATH/src/github.com/cboecking/goa-design-base/util/moeboe.properties

GOA_DESIGN_SOURCE=$GOPATH/src/github.com/cboecking/goa-design-base/design
GOA_DESIGN_DEST_BASE_REL=$MOEBOE_PROP_YOUR_REPO_NAME/$MOEBOE_PROP_YOUR_GIT_USER_NAME/$MOEBOE_PROP_GOA_PACKAGE_NAME
GOA_DESIGN_DEST_REL=$GOA_DESIGN_DEST_BASE_REL/design
GOA_DESIGN_DEST_BASE=$GOPATH/src/$GOA_DESIGN_DEST_BASE_REL
GOA_DESIGN_DEST=$GOA_DESIGN_DEST_BASE/design


#Action - check if GOA_DESIGN_DEST already exists, if so - exit
RESULT=$([ -d $GOA_DESIGN_DEST ] && echo "Y" || echo "N")
if [ $RESULT == "Y" ]; then
    echo "ISSUE: destination directory already exists. Stopping to prevent data loss."
    exit 1
fi

mkdir -p $GOA_DESIGN_DEST
cp $GOA_DESIGN_SOURCE/* $GOA_DESIGN_DEST/.

#set proper case
XAPI_NAME=${MOEBOE_PROP_GOA_API_NAME,,}
XRESOURCE_L=${MOEBOE_PROP_GOA_RESOURCE_NAME,,}
XRESOURCE_U=${XRESOURCE_L^}
XMEDIATYPE_L=${MOEBOE_PROP_GOA_MEDIA_TYPE_NAME,,}
XMEDIATYPE_U=${XMEDIATYPE_L^}
XSTORGROUP_L=${MOEBOE_PROP_GORMA_STORAGE_GROUP,,}
XSTORGROUP_U=${XSTORGROUP_L^}
XSTORMODEL_L=${MOEBOE_PROP_GORMA_STORAGE_MODEL,,}
XSTORMODEL_U=${XSTORMODEL_L^}


#set global replace text
sed -i "s|XAPI_NAME|$XAPI_NAME|g" $GOA_DESIGN_DEST/api_definition.go

sed -i "s|XRESOURCE_L|$XRESOURCE_L|g" $GOA_DESIGN_DEST/resources.go
sed -i "s|XRESOURCE_U|$XRESOURCE_U|g" $GOA_DESIGN_DEST/resources.go

sed -i "s|XMEDIATYPE_L|$XMEDIATYPE_L|g" $GOA_DESIGN_DEST/media_types.go
sed -i "s|XMEDIATYPE_U|$XMEDIATYPE_U|g" $GOA_DESIGN_DEST/media_types.go

sed -i "s|XSTORGROUP_L|$XSTORGROUP_L|g" $GOA_DESIGN_DEST/models.go
sed -i "s|XSTORGROUP_U|$XSTORGROUP_U|g" $GOA_DESIGN_DEST/models.go

sed -i "s|XSTORMODEL_L|$XSTORMODEL_L|g" $GOA_DESIGN_DEST/models.go
sed -i "s|XSTORMODEL_U|$XSTORMODEL_U|g" $GOA_DESIGN_DEST/models.go

echo "***************************"
echo "** Success"
echo "** Perform the following commands to create, build and start your application"
echo "** 1) cd $GOA_DESIGN_DEST_BASE"
echo "** 2) goagen bootstrap -d $GOA_DESIGN_DEST_REL"
echo "** 4) goagen --design=$GOA_DESIGN_DEST_REL gen --pkg-path=github.com/goadesign/gorma"
echo "** 5) go build ."
echo "** 6) ./$MOEBOE_PROP_GOA_PACKAGE_NAME"
echo "** 7) issue this command from another shell: http http://localhost:8081/$XAPI_NAME/${XRESOURCE_L}s"
echo "***************************"
