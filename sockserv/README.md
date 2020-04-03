## Idea

Run a socket server. Workers will connect to it and say "I am here!".
The controller will know how many workers are around, so will
reconcile the cluster.

As long as worker disconnects from the controller, then the controller
reconciles the cluster map again and reshuffles the client nodes
across the workers that left around.

## Caveats

What if worker flacky? Controller will start reconcile cluster every
time. If worker is too much flacky, then things will get nuts. How to
prevent that?

Possible solution:

1. Timeout worker to feel reliable. Single one worker does not need
   reconciliation of th cluster, because it gets all anyway. But if a
   new worker just joined the club, it gets on pause for e.g. 10-15
   minutes to be proven reliably running. If it gets off and on, its
   last join is recorded and so if the last join was too early, his
   default pause increases by some step



## Knowhow

1. Detect socket disconnect:
https://stackoverflow.com/questions/5686490/detect-socket-hangup-without-sending-or-receiving
