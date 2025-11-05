## MongoDB Atlas Metrics for Grafana

Monitor your MongoDB Atlas projects in Grafana using native Atlas metrics (process, database, and disk). This data source discovers projects and clusters automatically and lets you build panels and alerts from Atlas performance metrics.

![Panel example](https://raw.githubusercontent.com/fernandoescolar/grafana-mongodb-atlas-datasource/main/src/img/screenshots/query_example.png)

## Features

- Process, Database, and Disk metrics from MongoDB Atlas
- Project and Cluster discovery from your API keys
- Dimension selection (e.g., process, database, disk, metric dimension)
- Query alias templating with variables (project, cluster, mongo, database, disk, dimension)
- Works with Grafana alerting (backend plugin)
- Supports MongoDB Atlas Cloud API and Public Cloud API

## Requirements

- Grafana >= 7.0.0 (tested with Grafana 12.x)
- MongoDB Atlas Programmatic API Key with read access to the projects you want to monitor
- Network access from your Grafana server to the MongoDB Atlas API

## Installation

Pick the option that fits your setup.

### Option A: grafana-cli (from a release)

Install the latest packaged version directly from GitHub releases:

```bash
grafana-cli \
  --pluginUrl https://github.com/fernandoescolar/grafana-mongodb-atlas-datasource/releases/latest/download/fernandoescolar-mongodbatlas-datasource.zip \
  plugins install fernandoescolar-mongodbatlas-datasource
```

Restart Grafana after installation.

If the plugin is unsigned in your environment, allow it explicitly:

```bash
export GF_PLUGINS_ALLOW_LOADING_UNSIGNED_PLUGINS=fernandoescolar-mongodbatlas-datasource
```

### Option B: Docker

Run Grafana with the plugin installed automatically:

```bash
docker run -p 3000:3000 \
  -e GF_INSTALL_PLUGINS="https://github.com/fernandoescolar/grafana-mongodb-atlas-datasource/releases/latest/download/fernandoescolar-mongodbatlas-datasource.zip;fernandoescolar-mongodbatlas-datasource" \
  -e GF_PLUGINS_ALLOW_LOADING_UNSIGNED_PLUGINS=fernandoescolar-mongodbatlas-datasource \
  grafana/grafana:latest
```

### Option C: Manual (offline)

1) Download the ZIP from the latest release and extract it.
2) Copy the folder to Grafana's plugin directory, e.g. `/var/lib/grafana/plugins/fernandoescolar-mongodbatlas-datasource`.
3) Restart Grafana.
4) If needed, allow unsigned plugins via `GF_PLUGINS_ALLOW_LOADING_UNSIGNED_PLUGINS`.

For general plugin installation guidance, see the Grafana docs: https://grafana.com/docs/plugins/installation/

## Configuration

1) In Grafana, go to Connections → Data sources → Add data source.
2) Search for "MongoDB Atlas Metrics" and select it.
3) Enter your Programmatic API credentials:
   - Public Key
   - Private Key (stored securely by Grafana)
4) Choose API type:
   - MongoDB Atlas Cloud API (default)
   - MongoDB Public Cloud API
5) Click "Save & test" to verify the connection.

For creating API keys, see the MongoDB Atlas docs: https://docs.atlas.mongodb.com/configure-api-access/#programmatic-api-keys

![Select data source](https://raw.githubusercontent.com/fernandoescolar/grafana-mongodb-atlas-datasource/main/src/img/screenshots/datasource_list.png)
![Configure credentials](https://raw.githubusercontent.com/fernandoescolar/grafana-mongodb-atlas-datasource/main/src/img/screenshots/datasource_setup.png)

## Usage

Create a panel and add a query:

1) Select Project and Cluster (fetched from your account)
2) Choose a Metric type:
   - Process metrics: https://docs.atlas.mongodb.com/reference/api/process-measurements/
   - Database metrics: https://docs.atlas.mongodb.com/reference/api/process-databases-measurements/
   - Disk metrics: https://docs.atlas.mongodb.com/reference/api/process-disks-measurements/
3) Depending on the metric type, select additional fields (process, database, or disk)
4) Select a Dimension and run the query
5) Optionally set an Alias. You can use variables: `{{ projectName }}`, `{{ clusterName }}`, `{{ mongo }}`, `{{ database }}`, `{{ disk }}`, `{{ dimensionName }}`

![Build a query](https://raw.githubusercontent.com/fernandoescolar/grafana-mongodb-atlas-datasource/main/src/img/screenshots/query_setup.png)

## Links

- Changelog: https://github.com/fernandoescolar/grafana-mongodb-atlas-datasource/blob/main/CHANGELOG.md
- License (MIT): https://github.com/fernandoescolar/grafana-mongodb-atlas-datasource/blob/main/LICENSE
- Plugin ID: `fernandoescolar-mongodbatlas-datasource`
