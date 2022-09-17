#!/bin/bash
export SUSPECT_ID=$(head -n 179 streets/Buckingham_Place | tail -n 1 | cut -d "#" -f2)
echo $SUSPECT_ID
find -name "*$SUSPECT_ID" -exec cat {} \;
echo $MAIN_SUSPECT
