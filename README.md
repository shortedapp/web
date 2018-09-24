# Shorted

Shorted webapp

For an overview of the project structure please refer to the [Gatsby documentation - Building with Components](https://www.gatsbyjs.org/docs/building-with-components/).

## install

Make sure that you have the Gatsby CLI program installed:

```sh
npm install --global gatsby-cli
```

clone project and install packages:

```sh
git clone https://github.com/castlemilk/shorted.git
cd shorted
npm install
```

start development server:

```sh
ACTIVE_ENV=development gatsby develop
```

or

```sh
make web.develop
```

This runs the required commands via the Makefile found in the root of the project [here](Makefile)

## data

The shortedapp consumes from an available backend stack which is completely serverless, however for local development we have a full set of fixtures which enable simulated responses from the backend.

Additionall we have a range of CLI tools which help manage data for local development as well as for the "actual" production backend data storage such as elasticsearch, s3 and dynamoDB.

### fixtures

Sample has been generated by the instructions found [here](cli/sample_data/README.md).

The generated data can be found in [src/fixtures](src/services/sapi/fixtures)

### elasticsearch

Elasticsearch is used for the search components within the application. The makefile offers a range of commands for managaging the indexing of fixture data into elasticsearch as seen here:

```bash
Usage:
  make <target>

Targets:
  web.develop          start develop mode webapp
  web.build            build app for static serving
  web.serve            serve static build
  web.deploy           deploy webapp to s3
  elasticsearch.up     start elasticsearch cluster
  elasticsearch.down   stop elasticsearch cluster
  elasticsearch.restart restart elasticsearch cluster
  elasticsearch.logs   get elasticsearch cluster logs
  elasticsearch.test   test elaasticsearch endpoint
  nginx.install        intall nginx for CORS
  nginx.edit           edit nginx config
  nginx.configure      install project related nginx config
  nginx.restart        restart nginx
  nginx.test           test nginx reverse proxy
  index.init           initial company elasticsearch index
  index.create         create/ingest index into elasticsearch
  index.delete         delete elasticsearch index
  index.recreate       recreate index (fresh index)
  index.status         check index status
  help                 Show help
```
