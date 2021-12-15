# Deridex Backend

This is a backend implementation for Deridex market. The system contains several services written in Golang, and some supplementary products. All of the required component are dockerized, or use existing docker images.

## Install and launch

To launch the backend several steps are required:

 1. Clone the repository
 2. Move to the repo dir

    ```cd dex-backend```

 3. Build all required images

    ```docker-compose build```

 4. Start docker compose

    ```docker-compose up```

 5. Run the database migrations

    ```make seed```

https://kovan.infura.io/v3/10e6ea75b3174d3ca1c8e6cda6de0eac
