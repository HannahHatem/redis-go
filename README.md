# 🚀 Redis from Scratch

Welcome to my Redis implementation project! 🎉 This repository showcases my journey of building a Redis clone from scratch. Dive in to explore how I've implemented the RESP protocol, utilized mutex locks for thread-safe cache access, and set up replication for high availability.

## 🌟 Features

- **⚡ In-memory Data Store**: Store and retrieve your data at lightning speed.
- **📚 RESP Protocol**: Robust implementation of the Redis Serialization Protocol (RESP).
- **🔒 Thread-Safe Cache**: Utilizes mutex locks to ensure safe concurrent access to the cache.
- **🔁 Replication**: Set up replicas to distribute your data across multiple instances.
- **💾 Persistence**: Save your data to disk and restore it on restart.
- **📢 Pub/Sub Messaging**: Simple publish/subscribe messaging pattern.

## 🛠️ Getting Started

### Prerequisites

- Go 1.16 or higher

### Installation

1. Clone the repository:

    ```sh
    git clone https://github.com/yourusername/redis-from-scratch.git
    cd redis-from-scratch
    ```

2. Build the project:

    ```sh
    go build -o redis-server
    ```

3. Run the server:

    ```sh
    ./redis-server
    ```

## 🚀 Usage

Connect to your Redis server using a Redis client of your choice. Here are a few commands to get you started:

### 📜 Commands

#### 🔑 Set and Get

```sh
SET mykey "Hello, Redis!"
GET mykey
```

- `SET`: Stores a key-value pair.
- `GET`: Retrieves the value associated with a key.

#### 📋 Lists

```sh
RPUSH mylist "Hello"
RPUSH mylist "Redis"
LRANGE mylist 0 -1
```

- `RPUSH`: Adds an element to the end of a list.
- `LRANGE`: Retrieves a range of elements from a list.

#### 🔗 Sets

```sh
SADD myset "Hello"
SADD myset "Redis"
SMEMBERS myset
```

- `SADD`: Adds an element to a set.
- `SMEMBERS`: Returns all members of a set.

### 📢 Pub/Sub Messaging

```sh
SUBSCRIBE mychannel
PUBLISH mychannel "Hello, Redis!"
```

- `SUBSCRIBE`: Subscribes to a channel to receive messages.
- `PUBLISH`: Publishes a message to a channel.

### 💾 Persistence

Enable data persistence to save your data to disk and restore it on server restart. This ensures your data is not lost in case of server failure.

### 🔒 Thread-Safe Cache

To ensure safe concurrent access to the cache, I've used mutex locks. This guarantees that operations on the cache are thread-safe, preventing race conditions and ensuring data integrity.

### 🔁 Replication

Set up replication to distribute your data across multiple instances for high availability and reliability.

#### Setting Up a Replica

1. Start the master server:

    ```sh
    ./redis-server --port 6379
    ```

2. Start the replica server with the `--replicaof` flag:

    ```sh
    ./redis-server --port 6380 --replicaof 127.0.0.1:6379
    ```

The replica will automatically synchronize data from the master server. Any updates to the master will be propagated to the replica, ensuring data consistency across your Redis instances.

## 📚 Documentation

- [Commands](https://redis.io/commands)
- [Data Types](https://redis.io/topics/data-types)
- [Redis Serialization Protocol RESP](https://redis.io/docs/latest/develop/reference/protocol-spec/)
- [Replication](https://redis.io/topics/replication)
- [Persistence](https://redis.io/topics/persistence)
- [Pub/Sub](https://redis.io/topics/pubsub)

## 🤝 Contributing

Contributions, issues, and feature requests are welcome! Feel free to check the [issues page](https://github.com/HannahHatem/redis-from-scratch/issues) if you want to contribute.

## ✨ Acknowledgements

- Redis for the inspiration and reference.

## 📬 Contact

Feel free to reach out if you have any questions or suggestions! You can contact me at [hannahhatem.taha@gmail.com](mailto:hannahhatem.taha@gmail.com).

---

Enjoy exploring the world of Redis! 🚀✨

---
