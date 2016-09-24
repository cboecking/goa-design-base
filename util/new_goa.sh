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

#set variables with proper case
XAPI_NAME=${MOEBOE_PROP_GOA_API_NAME,,}
XAPI_PORT=$MOEBOE_PROP_GOA_API_PORT
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
echo "** NOTE: 	This program will execute; however, it will not persist any accounts."
echo "** 	You need to update the $XRESOURCE_L.go controller."
echo "**	Here is an example: https://github.com/goadesign/gorma-cellar/blob/master/account.go"
echo "**	If using a db, update main.go with the db details:" 
echo "**	Here is an example: https://github.com/goadesign/gorma-cellar/blob/master/main.go"
echo "** 7) use the following commands to interact with the service (commands begins with http):"
echo "****** LIST:    http http://localhost:8081/$XAPI_NAME/${XRESOURCE_L}s"
echo "****** CREATE:  http POST http://localhost:8081/$XAPI_NAME/${XRESOURCE_L}s name=chuck1"
echo "****** LIST:    http http://localhost:8081/$XAPI_NAME/${XRESOURCE_L}s"
echo "****** CREATE:  http POST http://localhost:8081/$XAPI_NAME/${XRESOURCE_L}s name=chuck2"
echo "****** LIST:    http http://localhost:8081/$XAPI_NAME/${XRESOURCE_L}s"
echo "****** SHOW:    http http://localhost:8081/$XAPI_NAME/${XRESOURCE_L}s/2"
echo "****** DELETE:  http DELETE http://localhost:8081/$XAPI_NAME/${XRESOURCE_L}s/2"
echo "****** UPDATE:  http PUT http://localhost:8081/$XAPI_NAME/${XRESOURCE_L}s/1 name=chuckles"
echo "****** LIST:    http http://localhost:8081/$XAPI_NAME/${XRESOURCE_L}s"
echo "***************************"
