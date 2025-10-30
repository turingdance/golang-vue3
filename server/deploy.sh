#!/bin/bash
APPNAME="microsite"
APPPATH="./$APPNAME"
case $1 in
        start)
                        if [ -f "$APPPATH" ];
                        then
                        nohup $APPPATH &
                        else
                          echo "resart $APPPATH fail because of $APPPATH not exist"
                        fi
                        ;;
        stop)
                        pkill -f $APPNAME
                        echo "pkill -f $APPNAME ok"
                        ;;
        restart)
                        if [ -f "$APPPATH" ];
                        then
                        pkill -f $APPNAME
                        nohup $APPPATH &
                        echo "restart $APPNAME @ $APPPATH OK"
                        else
                        echo "resart $APPPATH fail because of $APPPATH not exist"
                        fi
                        ;;
        status)
                       ps -aux | grep "$APPNAME"
                       ;;
         *)
        echo "Usage:`basename $0` (start|stop|restart|status)";;
esac
