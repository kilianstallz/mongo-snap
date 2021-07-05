# mongo-snap
Wrapper around the mongo-tools-cli.
The vision is to provide a simpler cli interface and provide a snapshot/backup tool to integrate into any CI/CD process.

## Usage

### Snapshot (dump)

````bash
mongo-snap snap --uri=mongouri --db=db --out=dump --exlude=users,profiles,events
````


### Restore

````bash
mongo-snap restore --uri=mongouri --drop --dir=dump
````
