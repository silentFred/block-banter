# Block Banter 🗣️🦜

Block Banter is a simple full stack app built in Go that fetches ERC20 transfer event data and visualises it in a
directed graph in the browser.

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

Note: it might take some time for your PostgreSQL table to fill up with data but once there are a few blocks worth of
data ingested the graph starts to look interesting.

### Demo video

[![Watch the video](https://raw.githubusercontent.com/silentFred/block-banter/main/thumbnail.png)](https://raw.githubusercontent.com/silentFred/block-banter/main/block-banter-ui-demo.mp4)

### Usage

Once the project is running, you can access the API to generate and retrieve proofs.

### Contributing

I'm learning how to build Go projects, so any feedback is appreciated. Contributions are welcome too so please open an
issue or submit a pull request for any improvements or bug fixes.

### License

This project is licensed under the MIT License. See the LICENSE-MIT file for details.

## TODO

### Backend

- Refactor the app into a layered architecture? (api, background, database, models, etc.)
- Pull balance for each address in the chart and display it in a heatmap style.
- Get dApp names for common addresses and display them in the chart.
- Introduce messaging with grpc and protobuf to communicate between the different components.
- Use websockets to communicate with the frontend instead of polling the api.
- Write tests for the different components.
- Kubernetes deployment.

### Frontend

- Deal with large set of data - introduce block range selection (slider perhaps) (rethink backend and api).
- Deal with multiple transactions between the same addresses.
- Fix strange re-rendering force that is pushing nodes away from the center.