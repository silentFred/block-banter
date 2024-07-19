# Block Banter 🗣️🦜

Block Banter is a simple full stack app built in Go that fetches ERC20 transfer event data and visualises it in a directed graph in the browser.

## Prerequisites

Ensure you have the following installed:

- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)

## Running the Project

To run the project locally, use Docker Compose:

```sh
docker-compose up
```

Open your browser and navigate to `http://localhost:9000` to view the graph.

Note: it might take some time for your PostgreSQL table to fill up with data but once there are a few blocks worth of data ingested the graph starts to look interesting.

### Usage

Once the project is running, you can access the API to generate and retrieve proofs.

### Contributing

I'm learning how to build Go projects, so any feedback is appreciated. Contributions are welcome too so please open an
issue or submit a pull request for any improvements or bug fixes.

### License

This project is licensed under the MIT License. See the LICENSE-MIT file for details.
