# Go Concurrency Learning Roadmap (20 Hands-On Problems)

## Goal

By completing these 20 problems, you should become comfortable with:

* Goroutines
* Channels
* Select
* WaitGroups
* Mutexes
* RWMutex
* Atomic Operations
* Context Cancellation
* Worker Pools
* Fan-Out/Fan-In
* Producer-Consumer
* Pipelines
* Backpressure
* Graceful Shutdown
* Concurrency-safe Data Structures
* Parallel Processing
* Real-world Backend Concurrency Patterns

---

# Level 1 — Foundations

## 1. Concurrent Word Counter

### Problem

You have 100 text files.

Count total words across all files concurrently.

### Example

```text
file1.txt
file2.txt
...
file100.txt
```

### Concepts

* Goroutines
* Channels
* WaitGroup

### Hint

Create one goroutine per file.

Send word counts to a result channel.

Aggregate counts in main goroutine.

### Testing

Generate 100 dummy text files.

Compare concurrent result with sequential result.

---

## 2. Concurrent URL Status Checker

### Problem

Check HTTP status codes for 100 URLs.

### Example

```text
google.com
github.com
reddit.com
```

### Concepts

* Network I/O
* Goroutines
* Result Collection

### Hint

Spawn one goroutine per URL.

Collect responses through a channel.

### Testing

Mix valid and invalid URLs.

Measure execution time.

---

## 3. Directory Size Calculator

### Problem

Calculate total size of a directory tree.

### Example

```text
Downloads/
├── Movies/
├── Images/
└── Documents/
```

### Concepts

* Recursive Concurrency
* WaitGroup
* Channel Aggregation

### Hint

Every directory can be processed concurrently.

### Testing

Create nested folders with random files.

Verify size with OS tools.

---

## 4. Producer-Consumer Queue

### Problem

Create producers generating jobs.

Consumers process jobs.

### Example

```text
Producer → Channel → Consumer
```

### Concepts

* Buffered Channels
* Producer Consumer Pattern

### Hint

Start multiple producers and consumers.

### Testing

Generate 10,000 jobs.

Verify every job is processed exactly once.

---

## 5. Fixed Worker Pool

### Problem

Process 100,000 jobs using only 10 workers.

### Concepts

* Worker Pool
* Bounded Concurrency

### Hint

Workers continuously read from a jobs channel.

### Testing

Add random delays.

Track throughput.

---

# Level 2 — Intermediate Patterns

## 6. Distributed Job Processor Simulation

### Problem

Simulate a job queue system.

### Example

```text
Dispatcher
    |
    v
Workers
```

### Concepts

* Worker Pool
* Context Cancellation
* Job Retry

### Hint

Create Job struct with status.

### Testing

Randomly fail workers.

Retry failed jobs.

---

## 7. Log Processing Pipeline

### Problem

Process log files through stages.

### Flow

```text
Read
 ↓
Parse
 ↓
Filter
 ↓
Store
```

### Concepts

* Pipelines
* Channel Chaining

### Testing

Inject malformed logs.

Verify filtered output.

---

## 8. CSV ETL Pipeline

### Problem

Build a mini ETL system.

### Flow

```text
Read CSV
 ↓
Transform
 ↓
Validate
 ↓
Write
```

### Concepts

* Multi-stage Pipelines
* Backpressure

### Testing

Generate 100k records.

Track throughput.

---

## 9. Dashboard Aggregator

### Problem

Simulate API Gateway aggregation.

### Example

```text
User Service
Order Service
Payment Service
```

All queried simultaneously.

### Concepts

* Fan-Out
* Fan-In

### Hint

Launch all requests concurrently.

Merge responses.

### Testing

Add random delays.

Measure latency improvements.

---

## 10. Search Aggregator

### Problem

Query multiple search providers.

Return merged results.

### Concepts

* Fan-Out
* Result Merging

### Testing

Introduce provider failures.

Verify partial results still work.

---

# Level 3 — Synchronization

## 11. Thread-Safe In-Memory Cache

### Problem

Implement:

```go
Get()
Set()
Delete()
```

### Concepts

* Mutex
* RWMutex

### Testing

Run 100 goroutines performing reads/writes.

Use race detector.

```bash
go test -race
```

---

## 12. Thread-Safe LRU Cache

### Problem

Build LRU cache.

### Features

```text
Capacity
Eviction
Concurrent Access
```

### Concepts

* RWMutex
* Shared State

### Testing

Heavy concurrent reads/writes.

Verify eviction correctness.

---

## 13. Concurrent Counter Service

### Problem

Expose:

```go
Increment()
Get()
```

### Concepts

* Atomic Operations

### Hint

Use atomic.Int64.

### Testing

1000 goroutines increment simultaneously.

Verify count.

---

## 14. Concurrent Metrics Collector

### Problem

Collect:

```text
Request Count
Error Count
Latency
```

### Concepts

* Atomics
* Aggregation

### Testing

Simulate API traffic.

---

# Level 4 — Context and Cancellation

## 15. Timeout Protected Downloader

### Problem

Download files with timeout.

### Concepts

* Context
* Select
* Cancellation

### Testing

Create slow HTTP server.

Cancel after timeout.

---

## 16. Distributed Worker Shutdown

### Problem

Workers process jobs.

Shutdown gracefully on signal.

### Concepts

* Context
* Graceful Shutdown

### Testing

Send SIGINT during execution.

Verify cleanup.

---

## 17. Concurrent Web Crawler

### Problem

Crawl pages starting from seed URL.

### Concepts

* Worker Pool
* Context
* Deduplication

### Testing

Limit crawl depth.

Track visited pages.

---

# Level 5 — Advanced Concurrency

## 18. Rate Limiter

### Problem

Implement:

```text
Token Bucket
```

### Concepts

* Mutex
* Ticker
* Background Goroutines

### Testing

Simulate burst traffic.

Verify rate enforcement.

---

## 19. Mini API Gateway

### Problem

Forward requests to backend services.

### Features

```text
Timeout
Retry
Load Balancing
```

### Concepts

* Context Propagation
* Fan-Out
* Select

### Testing

Add slow/failing services.

Measure behavior.

---

## 20. Mini MapReduce Framework

### Problem

Implement:

```text
Map
Shuffle
Reduce
```

### Example

Word Count

```text
Input Files
    ↓
Map
    ↓
Shuffle
    ↓
Reduce
```

### Concepts

* Parallel Processing
* Worker Pools
* Data Partitioning

### Testing

Process thousands of files.

Compare against sequential version.

---

# Stretch Goals

After completing all 20:

1. Distributed Rate Limiter
2. Load Balancer
3. Circuit Breaker
4. Job Queue
5. Kafka Consumer Group Simulator
6. Search Engine Indexer
7. Distributed Web Crawler
8. Concurrent File Search Engine

These projects combine nearly every concurrency pattern used in modern backend systems.
