# mongo-snap
CLI Tool wrapper around the mongo-tools cli.

## Usage

### Snapshot (dump)

````bash
mongo-snap snap --uri=mongouri --db=db --out=dump --exlude=users,profiles,events
````


### Restore

````bash
mongo-snap restore --uri=mongouri --drop --dir=dump
````
