A simple Pub/Sub example.

## Installation
1. Clone this repo
2. Run `docker-compose up -d` (Assuming docker is intalled and running locally).
3. Tail the logs of the producer and consumer containers to watch the magic.
   (`docker logs -f consultations`)

There are several containers built, a producer, consumer, and a bus with an
associated manager.

It takes a few moments before the containers have completed being created and
started, but all containers are configured to restart if there is a failure.

The producer sends events via the bus at random intervals, forever, and the
consumer will process them in the order that they were created (with kafka
determining that order).

There is a possibility for consumers to receive the same event multiple times.
This would be alleviated by putting a unique ID in the event allowing the
consumers to check a local store if the event had been dealt with previously,
and ensure that the saving of the event (and any associated processing) is
idempotent.


